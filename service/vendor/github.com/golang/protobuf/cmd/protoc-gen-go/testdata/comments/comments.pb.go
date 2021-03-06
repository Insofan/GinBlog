// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comments/comments.proto

// COMMENT: package goproto.protoc.comments;

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

// COMMENT: Message1
type Message1 struct {
	// COMMENT: Field1A
	Field1A *string `protobuf:"bytes,1,opt,name=Field1A" json:"Field1A,omitempty"`
	// COMMENT: Oneof1A
	//
	// Types that are valid to be assigned to Oneof1A:
	// COMMENT: Oneof1AField1
	//	*Message1_Oneof1AField1
	Oneof1A              isMessage1_Oneof1A `protobuf_oneof:"Oneof1a"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Message1) ProtoReflect() protoreflect.Message {
	return xxx_File_comments_comments_proto_messageTypes[0].MessageOf(m)
}
func (m *Message1) Reset()         { *m = Message1{} }
func (m *Message1) String() string { return proto.CompactTextString(m) }
func (*Message1) ProtoMessage()    {}
func (*Message1) Descriptor() ([]byte, []int) {
	return xxx_File_comments_comments_proto_rawdesc_gzipped, []int{0}
}

func (m *Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1.Unmarshal(m, b)
}
func (m *Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1.Marshal(b, m, deterministic)
}
func (m *Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1.Merge(m, src)
}
func (m *Message1) XXX_Size() int {
	return xxx_messageInfo_Message1.Size(m)
}
func (m *Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1.DiscardUnknown(m)
}

var xxx_messageInfo_Message1 proto.InternalMessageInfo

func (m *Message1) GetField1A() string {
	if m != nil && m.Field1A != nil {
		return *m.Field1A
	}
	return ""
}

type isMessage1_Oneof1A interface {
	isMessage1_Oneof1A()
}

type Message1_Oneof1AField1 struct {
	Oneof1AField1 string `protobuf:"bytes,2,opt,name=Oneof1AField1,oneof"`
}

func (*Message1_Oneof1AField1) isMessage1_Oneof1A() {}

func (m *Message1) GetOneof1A() isMessage1_Oneof1A {
	if m != nil {
		return m.Oneof1A
	}
	return nil
}

func (m *Message1) GetOneof1AField1() string {
	if x, ok := m.GetOneof1A().(*Message1_Oneof1AField1); ok {
		return x.Oneof1AField1
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message1) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message1_Oneof1AField1)(nil),
	}
}

// COMMENT: Message2
type Message2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message2) ProtoReflect() protoreflect.Message {
	return xxx_File_comments_comments_proto_messageTypes[1].MessageOf(m)
}
func (m *Message2) Reset()         { *m = Message2{} }
func (m *Message2) String() string { return proto.CompactTextString(m) }
func (*Message2) ProtoMessage()    {}
func (*Message2) Descriptor() ([]byte, []int) {
	return xxx_File_comments_comments_proto_rawdesc_gzipped, []int{1}
}

func (m *Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2.Unmarshal(m, b)
}
func (m *Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2.Marshal(b, m, deterministic)
}
func (m *Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2.Merge(m, src)
}
func (m *Message2) XXX_Size() int {
	return xxx_messageInfo_Message2.Size(m)
}
func (m *Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2.DiscardUnknown(m)
}

var xxx_messageInfo_Message2 proto.InternalMessageInfo

// COMMENT: Message1A
type Message1_Message1A struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message1_Message1A) ProtoReflect() protoreflect.Message {
	return xxx_File_comments_comments_proto_messageTypes[2].MessageOf(m)
}
func (m *Message1_Message1A) Reset()         { *m = Message1_Message1A{} }
func (m *Message1_Message1A) String() string { return proto.CompactTextString(m) }
func (*Message1_Message1A) ProtoMessage()    {}
func (*Message1_Message1A) Descriptor() ([]byte, []int) {
	return xxx_File_comments_comments_proto_rawdesc_gzipped, []int{0, 0}
}

func (m *Message1_Message1A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1_Message1A.Unmarshal(m, b)
}
func (m *Message1_Message1A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1_Message1A.Marshal(b, m, deterministic)
}
func (m *Message1_Message1A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1_Message1A.Merge(m, src)
}
func (m *Message1_Message1A) XXX_Size() int {
	return xxx_messageInfo_Message1_Message1A.Size(m)
}
func (m *Message1_Message1A) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1_Message1A.DiscardUnknown(m)
}

var xxx_messageInfo_Message1_Message1A proto.InternalMessageInfo

// COMMENT: Message1B
type Message1_Message1B struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message1_Message1B) ProtoReflect() protoreflect.Message {
	return xxx_File_comments_comments_proto_messageTypes[3].MessageOf(m)
}
func (m *Message1_Message1B) Reset()         { *m = Message1_Message1B{} }
func (m *Message1_Message1B) String() string { return proto.CompactTextString(m) }
func (*Message1_Message1B) ProtoMessage()    {}
func (*Message1_Message1B) Descriptor() ([]byte, []int) {
	return xxx_File_comments_comments_proto_rawdesc_gzipped, []int{0, 1}
}

func (m *Message1_Message1B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1_Message1B.Unmarshal(m, b)
}
func (m *Message1_Message1B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1_Message1B.Marshal(b, m, deterministic)
}
func (m *Message1_Message1B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1_Message1B.Merge(m, src)
}
func (m *Message1_Message1B) XXX_Size() int {
	return xxx_messageInfo_Message1_Message1B.Size(m)
}
func (m *Message1_Message1B) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1_Message1B.DiscardUnknown(m)
}

var xxx_messageInfo_Message1_Message1B proto.InternalMessageInfo

// COMMENT: Message2A
type Message2_Message2A struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message2_Message2A) ProtoReflect() protoreflect.Message {
	return xxx_File_comments_comments_proto_messageTypes[4].MessageOf(m)
}
func (m *Message2_Message2A) Reset()         { *m = Message2_Message2A{} }
func (m *Message2_Message2A) String() string { return proto.CompactTextString(m) }
func (*Message2_Message2A) ProtoMessage()    {}
func (*Message2_Message2A) Descriptor() ([]byte, []int) {
	return xxx_File_comments_comments_proto_rawdesc_gzipped, []int{1, 0}
}

func (m *Message2_Message2A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2_Message2A.Unmarshal(m, b)
}
func (m *Message2_Message2A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2_Message2A.Marshal(b, m, deterministic)
}
func (m *Message2_Message2A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2_Message2A.Merge(m, src)
}
func (m *Message2_Message2A) XXX_Size() int {
	return xxx_messageInfo_Message2_Message2A.Size(m)
}
func (m *Message2_Message2A) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2_Message2A.DiscardUnknown(m)
}

var xxx_messageInfo_Message2_Message2A proto.InternalMessageInfo

// COMMENT: Message2B
type Message2_Message2B struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message2_Message2B) ProtoReflect() protoreflect.Message {
	return xxx_File_comments_comments_proto_messageTypes[5].MessageOf(m)
}
func (m *Message2_Message2B) Reset()         { *m = Message2_Message2B{} }
func (m *Message2_Message2B) String() string { return proto.CompactTextString(m) }
func (*Message2_Message2B) ProtoMessage()    {}
func (*Message2_Message2B) Descriptor() ([]byte, []int) {
	return xxx_File_comments_comments_proto_rawdesc_gzipped, []int{1, 1}
}

func (m *Message2_Message2B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2_Message2B.Unmarshal(m, b)
}
func (m *Message2_Message2B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2_Message2B.Marshal(b, m, deterministic)
}
func (m *Message2_Message2B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2_Message2B.Merge(m, src)
}
func (m *Message2_Message2B) XXX_Size() int {
	return xxx_messageInfo_Message2_Message2B.Size(m)
}
func (m *Message2_Message2B) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2_Message2B.DiscardUnknown(m)
}

var xxx_messageInfo_Message2_Message2B proto.InternalMessageInfo

func init() {
	proto.RegisterFile("comments/comments.proto", xxx_File_comments_comments_proto_rawdesc_gzipped)
	proto.RegisterType((*Message1)(nil), "goproto.protoc.comments.Message1")
	proto.RegisterType((*Message2)(nil), "goproto.protoc.comments.Message2")
	proto.RegisterType((*Message1_Message1A)(nil), "goproto.protoc.comments.Message1.Message1A")
	proto.RegisterType((*Message1_Message1B)(nil), "goproto.protoc.comments.Message1.Message1B")
	proto.RegisterType((*Message2_Message2A)(nil), "goproto.protoc.comments.Message2.Message2A")
	proto.RegisterType((*Message2_Message2B)(nil), "goproto.protoc.comments.Message2.Message2B")
}

var xxx_File_comments_comments_proto_rawdesc = []byte{
	// 272 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x71, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x18,
	0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x41, 0x12, 0x26, 0x0a, 0x0d, 0x4f, 0x6e, 0x65, 0x6f,
	0x66, 0x31, 0x41, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0d, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x31, 0x41, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x1a, 0x0b, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x41, 0x1a, 0x0b, 0x0a,
	0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x42, 0x42, 0x09, 0x0a, 0x07, 0x4f, 0x6e,
	0x65, 0x6f, 0x66, 0x31, 0x61, 0x22, 0x24, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x1a, 0x0b, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x41, 0x1a, 0x0b,
	0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x42, 0x42, 0x43, 0x5a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
}

var xxx_File_comments_comments_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_comments_comments_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_comments_comments_proto protoreflect.FileDescriptor

var xxx_File_comments_comments_proto_messageTypes = make([]protoimpl.MessageType, 6)
var xxx_File_comments_comments_proto_goTypes = []interface{}{
	(*Message1)(nil),           // 0: goproto.protoc.comments.Message1
	(*Message2)(nil),           // 1: goproto.protoc.comments.Message2
	(*Message1_Message1A)(nil), // 2: goproto.protoc.comments.Message1.Message1A
	(*Message1_Message1B)(nil), // 3: goproto.protoc.comments.Message1.Message1B
	(*Message2_Message2A)(nil), // 4: goproto.protoc.comments.Message2.Message2A
	(*Message2_Message2B)(nil), // 5: goproto.protoc.comments.Message2.Message2B
}
var xxx_File_comments_comments_proto_depIdxs = []int32{}

func init() { xxx_File_comments_comments_proto_init() }
func xxx_File_comments_comments_proto_init() {
	if File_comments_comments_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 6)
	File_comments_comments_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_comments_comments_proto_rawdesc,
		GoTypes:            xxx_File_comments_comments_proto_goTypes,
		DependencyIndexes:  xxx_File_comments_comments_proto_depIdxs,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_comments_comments_proto_goTypes[0:][:6]
	for i, mt := range messageTypes {
		xxx_File_comments_comments_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_comments_comments_proto_messageTypes[i].PBType = mt
	}
	xxx_File_comments_comments_proto_goTypes = nil
	xxx_File_comments_comments_proto_depIdxs = nil
}
