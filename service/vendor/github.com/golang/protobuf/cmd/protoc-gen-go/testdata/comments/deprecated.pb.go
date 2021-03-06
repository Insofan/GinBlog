// Code generated by protoc-gen-go. DO NOT EDIT.
// comments/deprecated.proto is a deprecated file.

package comments

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

type DeprecatedEnum int32 // Deprecated: Do not use.
const (
	DeprecatedEnum_DEPRECATED DeprecatedEnum = 0 // Deprecated: Do not use.
)

func (e DeprecatedEnum) Type() protoreflect.EnumType {
	return xxx_File_comments_deprecated_proto_enumTypes[0]
}
func (e DeprecatedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var DeprecatedEnum_name = map[int32]string{
	0: "DEPRECATED",
}

var DeprecatedEnum_value = map[string]int32{
	"DEPRECATED": 0,
}

func (x DeprecatedEnum) String() string {
	return proto.EnumName(DeprecatedEnum_name, int32(x))
}

func (DeprecatedEnum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_comments_deprecated_proto_rawdesc_gzipped, []int{0}
}

// Deprecated: Do not use.
type DeprecatedMessage struct {
	DeprecatedField      string   `protobuf:"bytes,1,opt,name=deprecated_field,json=deprecatedField,proto3" json:"deprecated_field,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeprecatedMessage) ProtoReflect() protoreflect.Message {
	return xxx_File_comments_deprecated_proto_messageTypes[0].MessageOf(m)
}
func (m *DeprecatedMessage) Reset()         { *m = DeprecatedMessage{} }
func (m *DeprecatedMessage) String() string { return proto.CompactTextString(m) }
func (*DeprecatedMessage) ProtoMessage()    {}
func (*DeprecatedMessage) Descriptor() ([]byte, []int) {
	return xxx_File_comments_deprecated_proto_rawdesc_gzipped, []int{0}
}

func (m *DeprecatedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeprecatedMessage.Unmarshal(m, b)
}
func (m *DeprecatedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeprecatedMessage.Marshal(b, m, deterministic)
}
func (m *DeprecatedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeprecatedMessage.Merge(m, src)
}
func (m *DeprecatedMessage) XXX_Size() int {
	return xxx_messageInfo_DeprecatedMessage.Size(m)
}
func (m *DeprecatedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DeprecatedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DeprecatedMessage proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *DeprecatedMessage) GetDeprecatedField() string {
	if m != nil {
		return m.DeprecatedField
	}
	return ""
}

func init() {
	proto.RegisterFile("comments/deprecated.proto", xxx_File_comments_deprecated_proto_rawdesc_gzipped)
	proto.RegisterEnum("goproto.protoc.comments.DeprecatedEnum", DeprecatedEnum_name, DeprecatedEnum_value)
	proto.RegisterType((*DeprecatedMessage)(nil), "goproto.protoc.comments.DeprecatedMessage")
}

var xxx_File_comments_deprecated_proto_rawdesc = []byte{
	// 246 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x46, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x10, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x2a, 0x28, 0x0a, 0x0e,
	0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12,
	0x0a, 0x0a, 0x44, 0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x02,
	0x08, 0x01, 0x1a, 0x02, 0x18, 0x01, 0x42, 0x46, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0xb8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_comments_deprecated_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_comments_deprecated_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_comments_deprecated_proto protoreflect.FileDescriptor

var xxx_File_comments_deprecated_proto_enumTypes = make([]protoreflect.EnumType, 1)
var xxx_File_comments_deprecated_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_comments_deprecated_proto_goTypes = []interface{}{
	(DeprecatedEnum)(0),       // 0: goproto.protoc.comments.DeprecatedEnum
	(*DeprecatedMessage)(nil), // 1: goproto.protoc.comments.DeprecatedMessage
}
var xxx_File_comments_deprecated_proto_depIdxs = []int32{}

func init() { xxx_File_comments_deprecated_proto_init() }
func xxx_File_comments_deprecated_proto_init() {
	if File_comments_deprecated_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 1)
	File_comments_deprecated_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_comments_deprecated_proto_rawdesc,
		GoTypes:            xxx_File_comments_deprecated_proto_goTypes,
		DependencyIndexes:  xxx_File_comments_deprecated_proto_depIdxs,
		EnumOutputTypes:    xxx_File_comments_deprecated_proto_enumTypes,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_comments_deprecated_proto_goTypes[1:][:1]
	for i, mt := range messageTypes {
		xxx_File_comments_deprecated_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_comments_deprecated_proto_messageTypes[i].PBType = mt
	}
	xxx_File_comments_deprecated_proto_goTypes = nil
	xxx_File_comments_deprecated_proto_depIdxs = nil
}
