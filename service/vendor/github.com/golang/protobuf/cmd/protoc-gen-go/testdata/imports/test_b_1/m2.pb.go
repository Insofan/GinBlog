// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_b_1/m2.proto

package beta

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type M2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2) ProtoReflect() protoreflect.Message {
	return xxx_File_imports_test_b_1_m2_proto_messageTypes[0].MessageOf(m)
}
func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return xxx_File_imports_test_b_1_m2_proto_rawdesc_gzipped, []int{0}
}

func (m *M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2.Unmarshal(m, b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2.Marshal(b, m, deterministic)
}
func (m *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(m, src)
}
func (m *M2) XXX_Size() int {
	return xxx_messageInfo_M2.Size(m)
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func init() {
	proto.RegisterFile("imports/test_b_1/m2.proto", xxx_File_imports_test_b_1_m2_proto_rawdesc_gzipped)
	proto.RegisterType((*M2)(nil), "test.b.part2.M2")
}

var xxx_File_imports_test_b_1_m2_proto_rawdesc = []byte{
	// 137 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62,
	0x5f, 0x31, 0x2f, 0x6d, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x62, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x32, 0x22, 0x04, 0x0a, 0x02, 0x4d, 0x32, 0x42,
	0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x5f, 0x31, 0x3b, 0x62, 0x65, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_imports_test_b_1_m2_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_imports_test_b_1_m2_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_imports_test_b_1_m2_proto protoreflect.FileDescriptor

var xxx_File_imports_test_b_1_m2_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_imports_test_b_1_m2_proto_goTypes = []interface{}{
	(*M2)(nil), // 0: test.b.part2.M2
}
var xxx_File_imports_test_b_1_m2_proto_depIdxs = []int32{}

func init() { xxx_File_imports_test_b_1_m2_proto_init() }
func xxx_File_imports_test_b_1_m2_proto_init() {
	if File_imports_test_b_1_m2_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 1)
	File_imports_test_b_1_m2_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_imports_test_b_1_m2_proto_rawdesc,
		GoTypes:            xxx_File_imports_test_b_1_m2_proto_goTypes,
		DependencyIndexes:  xxx_File_imports_test_b_1_m2_proto_depIdxs,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_imports_test_b_1_m2_proto_goTypes[0:][:1]
	for i, mt := range messageTypes {
		xxx_File_imports_test_b_1_m2_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_imports_test_b_1_m2_proto_messageTypes[i].PBType = mt
	}
	xxx_File_imports_test_b_1_m2_proto_goTypes = nil
	xxx_File_imports_test_b_1_m2_proto_depIdxs = nil
}
