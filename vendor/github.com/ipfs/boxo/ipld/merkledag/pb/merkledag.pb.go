// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: github.com/ipfs/boxo/ipld/merkledag/pb/merkledag.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

// DoNotUpgradeFileEverItWillChangeYourHashes warns users about not breaking
// their file hashes.
const DoNotUpgradeFileEverItWillChangeYourHashes = `
This file does not produce canonical protobufs. Unfortunately, if we change it,
we'll change the hashes of the files we produce.

Do *not regenerate this file.
`

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An IPFS MerkleDAG Link
type PBLink struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// multihash of the target object
	Hash []byte `protobuf:"bytes,1,opt,name=Hash" json:"Hash,omitempty"`
	// utf string name. should be unique per object
	Name *string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	// cumulative size of target object
	Tsize         *uint64 `protobuf:"varint,3,opt,name=Tsize" json:"Tsize,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PBLink) Reset() {
	*x = PBLink{}
	mi := &file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PBLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBLink) ProtoMessage() {}

func (x *PBLink) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBLink.ProtoReflect.Descriptor instead.
func (*PBLink) Descriptor() ([]byte, []int) {
	return file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDescGZIP(), []int{0}
}

func (x *PBLink) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *PBLink) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *PBLink) GetTsize() uint64 {
	if x != nil && x.Tsize != nil {
		return *x.Tsize
	}
	return 0
}

// An IPFS MerkleDAG Node
type PBNode struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// opaque user data
	Data []byte `protobuf:"bytes,1,opt,name=Data" json:"Data,omitempty"`
	// refs to other objects
	Links         []*PBLink `protobuf:"bytes,2,rep,name=Links" json:"Links,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PBNode) Reset() {
	*x = PBNode{}
	mi := &file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PBNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PBNode) ProtoMessage() {}

func (x *PBNode) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PBNode.ProtoReflect.Descriptor instead.
func (*PBNode) Descriptor() ([]byte, []int) {
	return file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDescGZIP(), []int{1}
}

func (x *PBNode) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PBNode) GetLinks() []*PBLink {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto protoreflect.FileDescriptor

var file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70, 0x66,
	0x73, 0x2f, 0x62, 0x6f, 0x78, 0x6f, 0x2f, 0x69, 0x70, 0x6c, 0x64, 0x2f, 0x6d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x64, 0x61, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x64,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x69, 0x70, 0x66, 0x73, 0x2e, 0x62,
	0x6f, 0x78, 0x6f, 0x2e, 0x69, 0x70, 0x6c, 0x64, 0x2e, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x64,
	0x61, 0x67, 0x2e, 0x70, 0x62, 0x22, 0x46, 0x0a, 0x06, 0x50, 0x42, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x54, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x57, 0x0a,
	0x06, 0x50, 0x42, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x05, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x70, 0x66,
	0x73, 0x2e, 0x62, 0x6f, 0x78, 0x6f, 0x2e, 0x69, 0x70, 0x6c, 0x64, 0x2e, 0x6d, 0x65, 0x72, 0x6b,
	0x6c, 0x65, 0x64, 0x61, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x42, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70, 0x66, 0x73, 0x2f, 0x62, 0x6f, 0x78, 0x6f, 0x2f, 0x69,
	0x70, 0x6c, 0x64, 0x2f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x64, 0x61, 0x67, 0x2f, 0x70, 0x62,
}

var (
	file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDescOnce sync.Once
	file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDescData = file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDesc
)

func file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDescGZIP() []byte {
	file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDescOnce.Do(func() {
		file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDescData)
	})
	return file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDescData
}

var file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_goTypes = []any{
	(*PBLink)(nil), // 0: ipfs.boxo.ipld.merkledag.pb.PBLink
	(*PBNode)(nil), // 1: ipfs.boxo.ipld.merkledag.pb.PBNode
}
var file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_depIdxs = []int32{
	0, // 0: ipfs.boxo.ipld.merkledag.pb.PBNode.Links:type_name -> ipfs.boxo.ipld.merkledag.pb.PBLink
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_init() }
func file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_init() {
	if File_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_goTypes,
		DependencyIndexes: file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_depIdxs,
		MessageInfos:      file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_msgTypes,
	}.Build()
	File_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto = out.File
	file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_rawDesc = nil
	file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_goTypes = nil
	file_github_com_ipfs_boxo_ipld_merkledag_pb_merkledag_proto_depIdxs = nil
}
