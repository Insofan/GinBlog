// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/b.proto

package import_public

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	sub "github.com/golang/protobuf/v2/cmd/protoc-gen-go/testdata/import_public/sub"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Local struct {
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	E                    *sub.E   `protobuf:"varint,2,opt,name=e,enum=goproto.protoc.import_public.sub.E" json:"e,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Local) ProtoReflect() protoreflect.Message {
	return xxx_File_import_public_b_proto_messageTypes[0].MessageOf(m)
}
func (m *Local) Reset()         { *m = Local{} }
func (m *Local) String() string { return proto.CompactTextString(m) }
func (*Local) ProtoMessage()    {}
func (*Local) Descriptor() ([]byte, []int) {
	return xxx_File_import_public_b_proto_rawdesc_gzipped, []int{0}
}

func (m *Local) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Local.Unmarshal(m, b)
}
func (m *Local) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Local.Marshal(b, m, deterministic)
}
func (m *Local) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Local.Merge(m, src)
}
func (m *Local) XXX_Size() int {
	return xxx_messageInfo_Local.Size(m)
}
func (m *Local) XXX_DiscardUnknown() {
	xxx_messageInfo_Local.DiscardUnknown(m)
}

var xxx_messageInfo_Local proto.InternalMessageInfo

func (m *Local) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Local) GetE() sub.E {
	if m != nil && m.E != nil {
		return *m.E
	}
	return sub.E_ZERO
}

func init() {
	proto.RegisterFile("import_public/b.proto", xxx_File_import_public_b_proto_rawdesc_gzipped)
	proto.RegisterType((*Local)(nil), "goproto.protoc.import_public.Local")
}

var xxx_File_import_public_b_proto_rawdesc = []byte{
	// 265 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x15, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x1a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x75, 0x62, 0x2f, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6d, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x01, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x4d, 0x52, 0x01, 0x6d, 0x12, 0x31, 0x0a, 0x01,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x45, 0x52, 0x01, 0x65, 0x42,
	0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
}

var xxx_File_import_public_b_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_import_public_b_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_import_public_b_proto protoreflect.FileDescriptor

var xxx_File_import_public_b_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_import_public_b_proto_goTypes = []interface{}{
	(*Local)(nil), // 0: goproto.protoc.import_public.Local
	(*sub.M)(nil), // 1: goproto.protoc.import_public.sub.M
	(sub.E)(0),    // 2: goproto.protoc.import_public.sub.E
}
var xxx_File_import_public_b_proto_depIdxs = []int32{
	1, // goproto.protoc.import_public.Local.m:type_name -> goproto.protoc.import_public.sub.M
	2, // goproto.protoc.import_public.Local.e:type_name -> goproto.protoc.import_public.sub.E
}

func init() { xxx_File_import_public_b_proto_init() }
func xxx_File_import_public_b_proto_init() {
	if File_import_public_b_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 1)
	File_import_public_b_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_import_public_b_proto_rawdesc,
		GoTypes:            xxx_File_import_public_b_proto_goTypes,
		DependencyIndexes:  xxx_File_import_public_b_proto_depIdxs,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_import_public_b_proto_goTypes[0:][:1]
	for i, mt := range messageTypes {
		xxx_File_import_public_b_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_import_public_b_proto_messageTypes[i].PBType = mt
	}
	xxx_File_import_public_b_proto_goTypes = nil
	xxx_File_import_public_b_proto_depIdxs = nil
}
