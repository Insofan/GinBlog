// Code generated by protoc-gen-go. DO NOT EDIT.
// source: legacy/legacy.proto

package legacy

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	proto2_v0_0 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v0.0.0-20160225-2fc053c5"
	proto2_v0_01 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v0.0.0-20160519-a4ab9ec5"
	proto2_v1_0 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.0.0-20180125-92554152"
	proto2_v1_1 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.1.0-20180430-b4deda09"
	proto2_v1_2 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.2.0-20180814-aa810b61"
	proto2_v1_21 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto2.v1.2.1-20181126-8d0c54c1"
	proto3_v0_0 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v0.0.0-20160225-2fc053c5"
	proto3_v0_01 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v0.0.0-20160519-a4ab9ec5"
	proto3_v1_0 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.0.0-20180125-92554152"
	proto3_v1_1 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.1.0-20180430-b4deda09"
	proto3_v1_2 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.2.0-20180814-aa810b61"
	proto3_v1_21 "github.com/golang/protobuf/v2/internal/testprotos/legacy/proto3.v1.2.1-20181126-8d0c54c1"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Legacy struct {
	F1                   *proto2_v0_0.Message  `protobuf:"bytes,1,opt,name=f1,proto3" json:"f1,omitempty"`
	F2                   *proto3_v0_0.Message  `protobuf:"bytes,2,opt,name=f2,proto3" json:"f2,omitempty"`
	F3                   *proto2_v0_01.Message `protobuf:"bytes,3,opt,name=f3,proto3" json:"f3,omitempty"`
	F4                   *proto3_v0_01.Message `protobuf:"bytes,4,opt,name=f4,proto3" json:"f4,omitempty"`
	F5                   *proto2_v1_0.Message  `protobuf:"bytes,5,opt,name=f5,proto3" json:"f5,omitempty"`
	F6                   *proto3_v1_0.Message  `protobuf:"bytes,6,opt,name=f6,proto3" json:"f6,omitempty"`
	F7                   *proto2_v1_1.Message  `protobuf:"bytes,7,opt,name=f7,proto3" json:"f7,omitempty"`
	F8                   *proto3_v1_1.Message  `protobuf:"bytes,8,opt,name=f8,proto3" json:"f8,omitempty"`
	F9                   *proto2_v1_2.Message  `protobuf:"bytes,9,opt,name=f9,proto3" json:"f9,omitempty"`
	F10                  *proto3_v1_2.Message  `protobuf:"bytes,10,opt,name=f10,proto3" json:"f10,omitempty"`
	F11                  *proto2_v1_21.Message `protobuf:"bytes,11,opt,name=f11,proto3" json:"f11,omitempty"`
	F12                  *proto3_v1_21.Message `protobuf:"bytes,12,opt,name=f12,proto3" json:"f12,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Legacy) ProtoReflect() protoreflect.Message {
	return xxx_File_legacy_legacy_proto_messageTypes[0].MessageOf(m)
}
func (m *Legacy) Reset()         { *m = Legacy{} }
func (m *Legacy) String() string { return proto.CompactTextString(m) }
func (*Legacy) ProtoMessage()    {}
func (*Legacy) Descriptor() ([]byte, []int) {
	return xxx_File_legacy_legacy_proto_rawdesc_gzipped, []int{0}
}

func (m *Legacy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Legacy.Unmarshal(m, b)
}
func (m *Legacy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Legacy.Marshal(b, m, deterministic)
}
func (m *Legacy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Legacy.Merge(m, src)
}
func (m *Legacy) XXX_Size() int {
	return xxx_messageInfo_Legacy.Size(m)
}
func (m *Legacy) XXX_DiscardUnknown() {
	xxx_messageInfo_Legacy.DiscardUnknown(m)
}

var xxx_messageInfo_Legacy proto.InternalMessageInfo

func (m *Legacy) GetF1() *proto2_v0_0.Message {
	if m != nil {
		return m.F1
	}
	return nil
}

func (m *Legacy) GetF2() *proto3_v0_0.Message {
	if m != nil {
		return m.F2
	}
	return nil
}

func (m *Legacy) GetF3() *proto2_v0_01.Message {
	if m != nil {
		return m.F3
	}
	return nil
}

func (m *Legacy) GetF4() *proto3_v0_01.Message {
	if m != nil {
		return m.F4
	}
	return nil
}

func (m *Legacy) GetF5() *proto2_v1_0.Message {
	if m != nil {
		return m.F5
	}
	return nil
}

func (m *Legacy) GetF6() *proto3_v1_0.Message {
	if m != nil {
		return m.F6
	}
	return nil
}

func (m *Legacy) GetF7() *proto2_v1_1.Message {
	if m != nil {
		return m.F7
	}
	return nil
}

func (m *Legacy) GetF8() *proto3_v1_1.Message {
	if m != nil {
		return m.F8
	}
	return nil
}

func (m *Legacy) GetF9() *proto2_v1_2.Message {
	if m != nil {
		return m.F9
	}
	return nil
}

func (m *Legacy) GetF10() *proto3_v1_2.Message {
	if m != nil {
		return m.F10
	}
	return nil
}

func (m *Legacy) GetF11() *proto2_v1_21.Message {
	if m != nil {
		return m.F11
	}
	return nil
}

func (m *Legacy) GetF12() *proto3_v1_21.Message {
	if m != nil {
		return m.F12
	}
	return nil
}

func init() {
	proto.RegisterFile("legacy/legacy.proto", xxx_File_legacy_legacy_proto_rawdesc_gzipped)
	proto.RegisterType((*Legacy)(nil), "google.golang.org.Legacy")
}

var xxx_File_legacy_legacy_proto_rawdesc = []byte{
	// 1457 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x13, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x1a, 0x31, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x76, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2d, 0x32,
	0x30, 0x31, 0x36, 0x30, 0x32, 0x32, 0x35, 0x2d, 0x32, 0x66, 0x63, 0x30, 0x35, 0x33, 0x63, 0x35,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6c, 0x65, 0x67,
	0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x76, 0x30, 0x2e, 0x30, 0x2e,
	0x30, 0x2d, 0x32, 0x30, 0x31, 0x36, 0x30, 0x32, 0x32, 0x35, 0x2d, 0x32, 0x66, 0x63, 0x30, 0x35,
	0x33, 0x63, 0x35, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31,
	0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x76, 0x30,
	0x2e, 0x30, 0x2e, 0x30, 0x2d, 0x32, 0x30, 0x31, 0x36, 0x30, 0x35, 0x31, 0x39, 0x2d, 0x61, 0x34,
	0x61, 0x62, 0x39, 0x65, 0x63, 0x35, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x31, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x2e, 0x76, 0x30, 0x2e, 0x30, 0x2e, 0x30, 0x2d, 0x32, 0x30, 0x31, 0x36, 0x30, 0x35, 0x31, 0x39,
	0x2d, 0x61, 0x34, 0x61, 0x62, 0x39, 0x65, 0x63, 0x35, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x2d, 0x32, 0x30, 0x31, 0x38, 0x30,
	0x31, 0x32, 0x35, 0x2d, 0x39, 0x32, 0x35, 0x35, 0x34, 0x31, 0x35, 0x32, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x2d, 0x32, 0x30,
	0x31, 0x38, 0x30, 0x31, 0x32, 0x35, 0x2d, 0x39, 0x32, 0x35, 0x35, 0x34, 0x31, 0x35, 0x32, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6c, 0x65, 0x67, 0x61,
	0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x31, 0x2e, 0x30,
	0x2d, 0x32, 0x30, 0x31, 0x38, 0x30, 0x34, 0x33, 0x30, 0x2d, 0x62, 0x34, 0x64, 0x65, 0x64, 0x61,
	0x30, 0x39, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x76, 0x31, 0x2e,
	0x31, 0x2e, 0x30, 0x2d, 0x32, 0x30, 0x31, 0x38, 0x30, 0x34, 0x33, 0x30, 0x2d, 0x62, 0x34, 0x64,
	0x65, 0x64, 0x61, 0x30, 0x39, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x31, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e,
	0x76, 0x31, 0x2e, 0x32, 0x2e, 0x30, 0x2d, 0x32, 0x30, 0x31, 0x38, 0x30, 0x38, 0x31, 0x34, 0x2d,
	0x61, 0x61, 0x38, 0x31, 0x30, 0x62, 0x36, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x32, 0x2e, 0x30, 0x2d, 0x32, 0x30, 0x31, 0x38, 0x30, 0x38,
	0x31, 0x34, 0x2d, 0x61, 0x61, 0x38, 0x31, 0x30, 0x62, 0x36, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x32, 0x2e, 0x31, 0x2d, 0x32, 0x30, 0x31,
	0x38, 0x31, 0x31, 0x32, 0x36, 0x2d, 0x38, 0x64, 0x30, 0x63, 0x35, 0x34, 0x63, 0x31, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x32, 0x2e, 0x31, 0x2d,
	0x32, 0x30, 0x31, 0x38, 0x31, 0x31, 0x32, 0x36, 0x2d, 0x38, 0x64, 0x30, 0x63, 0x35, 0x34, 0x63,
	0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x05, 0x0a,
	0x06, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32,
	0x30, 0x31, 0x36, 0x30, 0x32, 0x32, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x02, 0x66, 0x31, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x36, 0x30,
	0x32, 0x32, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x32, 0x12,
	0x3a, 0x0a, 0x02, 0x66, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32, 0x30, 0x31, 0x36, 0x30, 0x35, 0x31, 0x39, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x33, 0x12, 0x3a, 0x0a, 0x02, 0x66,
	0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x36, 0x30, 0x35, 0x31, 0x39, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x34, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x35, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32,
	0x30, 0x31, 0x38, 0x30, 0x31, 0x32, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x02, 0x66, 0x35, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x38, 0x30,
	0x31, 0x32, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x36, 0x12,
	0x3a, 0x0a, 0x02, 0x66, 0x37, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32, 0x30, 0x31, 0x38, 0x30, 0x34, 0x33, 0x30, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x37, 0x12, 0x3a, 0x0a, 0x02, 0x66,
	0x38, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x38, 0x30, 0x34, 0x33, 0x30, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x02, 0x66, 0x38, 0x12, 0x3a, 0x0a, 0x02, 0x66, 0x39, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32,
	0x30, 0x31, 0x38, 0x30, 0x38, 0x31, 0x34, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x02, 0x66, 0x39, 0x12, 0x3c, 0x0a, 0x03, 0x66, 0x31, 0x30, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x38,
	0x30, 0x38, 0x31, 0x34, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x66, 0x31,
	0x30, 0x12, 0x3c, 0x0a, 0x03, 0x66, 0x31, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5f, 0x32, 0x30, 0x31, 0x38, 0x31, 0x31,
	0x32, 0x36, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x66, 0x31, 0x31, 0x12,
	0x3c, 0x0a, 0x03, 0x66, 0x31, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x5f, 0x32, 0x30, 0x31, 0x38, 0x31, 0x31, 0x32, 0x36,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x66, 0x31, 0x32, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var xxx_File_legacy_legacy_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_legacy_legacy_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_legacy_legacy_proto protoreflect.FileDescriptor

var xxx_File_legacy_legacy_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_legacy_legacy_proto_goTypes = []interface{}{
	(*Legacy)(nil),               // 0: google.golang.org.Legacy
	(*proto2_v0_0.Message)(nil),  // 1: google.golang.org.proto2_20160225.Message
	(*proto3_v0_0.Message)(nil),  // 2: google.golang.org.proto3_20160225.Message
	(*proto2_v0_01.Message)(nil), // 3: google.golang.org.proto2_20160519.Message
	(*proto3_v0_01.Message)(nil), // 4: google.golang.org.proto3_20160519.Message
	(*proto2_v1_0.Message)(nil),  // 5: google.golang.org.proto2_20180125.Message
	(*proto3_v1_0.Message)(nil),  // 6: google.golang.org.proto3_20180125.Message
	(*proto2_v1_1.Message)(nil),  // 7: google.golang.org.proto2_20180430.Message
	(*proto3_v1_1.Message)(nil),  // 8: google.golang.org.proto3_20180430.Message
	(*proto2_v1_2.Message)(nil),  // 9: google.golang.org.proto2_20180814.Message
	(*proto3_v1_2.Message)(nil),  // 10: google.golang.org.proto3_20180814.Message
	(*proto2_v1_21.Message)(nil), // 11: google.golang.org.proto2_20181126.Message
	(*proto3_v1_21.Message)(nil), // 12: google.golang.org.proto3_20181126.Message
}
var xxx_File_legacy_legacy_proto_depIdxs = []int32{
	1,  // google.golang.org.Legacy.f1:type_name -> google.golang.org.proto2_20160225.Message
	2,  // google.golang.org.Legacy.f2:type_name -> google.golang.org.proto3_20160225.Message
	3,  // google.golang.org.Legacy.f3:type_name -> google.golang.org.proto2_20160519.Message
	4,  // google.golang.org.Legacy.f4:type_name -> google.golang.org.proto3_20160519.Message
	5,  // google.golang.org.Legacy.f5:type_name -> google.golang.org.proto2_20180125.Message
	6,  // google.golang.org.Legacy.f6:type_name -> google.golang.org.proto3_20180125.Message
	7,  // google.golang.org.Legacy.f7:type_name -> google.golang.org.proto2_20180430.Message
	8,  // google.golang.org.Legacy.f8:type_name -> google.golang.org.proto3_20180430.Message
	9,  // google.golang.org.Legacy.f9:type_name -> google.golang.org.proto2_20180814.Message
	10, // google.golang.org.Legacy.f10:type_name -> google.golang.org.proto3_20180814.Message
	11, // google.golang.org.Legacy.f11:type_name -> google.golang.org.proto2_20181126.Message
	12, // google.golang.org.Legacy.f12:type_name -> google.golang.org.proto3_20181126.Message
}

func init() { xxx_File_legacy_legacy_proto_init() }
func xxx_File_legacy_legacy_proto_init() {
	if File_legacy_legacy_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 1)
	File_legacy_legacy_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_legacy_legacy_proto_rawdesc,
		GoTypes:            xxx_File_legacy_legacy_proto_goTypes,
		DependencyIndexes:  xxx_File_legacy_legacy_proto_depIdxs,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_legacy_legacy_proto_goTypes[0:][:1]
	for i, mt := range messageTypes {
		xxx_File_legacy_legacy_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_legacy_legacy_proto_messageTypes[i].PBType = mt
	}
	xxx_File_legacy_legacy_proto_goTypes = nil
	xxx_File_legacy_legacy_proto_depIdxs = nil
}
