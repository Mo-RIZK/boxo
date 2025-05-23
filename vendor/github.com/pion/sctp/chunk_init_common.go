// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package sctp

import (
	"encoding/binary"
	"errors"
	"fmt"
)

/*
chunkInitCommon represents an SCTP Chunk body of type INIT and INIT ACK

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|   Type = 1    |  Chunk Flags  |      Chunk Length             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         Initiate Tag                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|           Advertised Receiver Window Credit (a_rwnd)          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|  Number of Outbound Streams   |  Number of Inbound Streams    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                          Initial TSN                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               |
|              Optional/Variable-Length Parameters              |
|                                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

The INIT chunk contains the following parameters.  Unless otherwise
noted, each parameter MUST only be included once in the INIT chunk.

Fixed Parameters                     Status
----------------------------------------------
Initiate Tag                        Mandatory
Advertised Receiver Window Credit   Mandatory
Number of Outbound Streams          Mandatory
Number of Inbound Streams           Mandatory
Initial TSN                         Mandatory
*/

type chunkInitCommon struct {
	initiateTag                    uint32
	advertisedReceiverWindowCredit uint32
	numOutboundStreams             uint16
	numInboundStreams              uint16
	initialTSN                     uint32
	params                         []param
	unrecognizedParams             []paramHeader
}

const (
	initChunkMinLength          = 16
	initOptionalVarHeaderLength = 4
)

// Init chunk errors.
var (
	ErrInitChunkParseParamTypeFailed = errors.New("failed to parse param type")
	ErrInitAckMarshalParam           = errors.New("unable to marshal parameter for INIT/INITACK")
)

func (i *chunkInitCommon) unmarshal(raw []byte) error {
	i.initiateTag = binary.BigEndian.Uint32(raw[0:])
	i.advertisedReceiverWindowCredit = binary.BigEndian.Uint32(raw[4:])
	i.numOutboundStreams = binary.BigEndian.Uint16(raw[8:])
	i.numInboundStreams = binary.BigEndian.Uint16(raw[10:])
	i.initialTSN = binary.BigEndian.Uint32(raw[12:])

	// https://tools.ietf.org/html/rfc4960#section-3.2.1
	//
	// Chunk values of SCTP control chunks consist of a chunk-type-specific
	// header of required fields, followed by zero or more parameters.  The
	// optional and variable-length parameters contained in a chunk are
	// defined in a Type-Length-Value format as shown below.
	//
	// 0                   1                   2                   3
	// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |          Parameter Type       |       Parameter Length        |
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |                                                               |
	// |                       Parameter Value                         |
	// |                                                               |
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

	offset := initChunkMinLength
	remaining := len(raw) - offset
	for remaining > 0 {
		if remaining > initOptionalVarHeaderLength {
			var pHeader paramHeader
			if err := pHeader.unmarshal(raw[offset:]); err != nil {
				return fmt.Errorf("%w: %v", ErrInitChunkParseParamTypeFailed, err) //nolint:errorlint
			}

			p, err := buildParam(pHeader.typ, raw[offset:])
			if err != nil {
				i.unrecognizedParams = append(i.unrecognizedParams, pHeader)
			} else {
				i.params = append(i.params, p)
			}

			padding := getPadding(pHeader.length())
			offset += pHeader.length() + padding
			remaining -= pHeader.length() + padding
		} else {
			break
		}
	}

	return nil
}

func (i *chunkInitCommon) marshal() ([]byte, error) {
	out := make([]byte, initChunkMinLength)
	binary.BigEndian.PutUint32(out[0:], i.initiateTag)
	binary.BigEndian.PutUint32(out[4:], i.advertisedReceiverWindowCredit)
	binary.BigEndian.PutUint16(out[8:], i.numOutboundStreams)
	binary.BigEndian.PutUint16(out[10:], i.numInboundStreams)
	binary.BigEndian.PutUint32(out[12:], i.initialTSN)
	for idx, p := range i.params {
		pp, err := p.marshal()
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInitAckMarshalParam, err) //nolint:errorlint
		}

		out = append(out, pp...) //nolint:makezero // TODO: fix

		// Chunks (including Type, Length, and Value fields) are padded out
		// by the sender with all zero bytes to be a multiple of 4 bytes
		// long.  This padding MUST NOT be more than 3 bytes in total.  The
		// Chunk Length value does not include terminating padding of the
		// chunk.  *However, it does include padding of any variable-length
		// parameter except the last parameter in the chunk.*  The receiver
		// MUST ignore the padding.
		if idx != len(i.params)-1 {
			out = padByte(out, getPadding(len(pp)))
		}
	}

	return out, nil
}

// String makes chunkInitCommon printable.
func (i chunkInitCommon) String() string {
	format := `initiateTag: %d
	advertisedReceiverWindowCredit: %d
	numOutboundStreams: %d
	numInboundStreams: %d
	initialTSN: %d`

	res := fmt.Sprintf(format,
		i.initiateTag,
		i.advertisedReceiverWindowCredit,
		i.numOutboundStreams,
		i.numInboundStreams,
		i.initialTSN,
	)

	for i, param := range i.params {
		res += fmt.Sprintf("Param %d:\n %s", i, param)
	}

	return res
}
