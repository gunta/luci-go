// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/lucicfg/testproto/test.proto

package testproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Msg struct {
	I                    int64    `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_f604da41cdec35eb, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetI() int64 {
	if m != nil {
		return m.I
	}
	return 0
}

func init() {
	proto.RegisterType((*Msg)(nil), "testproto.Msg")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/lucicfg/testproto/test.proto", fileDescriptor_f604da41cdec35eb)
}

var fileDescriptor_f604da41cdec35eb = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0x04, 0x13, 0xc9, 0x69, 0xe9, 0xfa, 0x25, 0xa9, 0xc5, 0x25, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x60,
	0x96, 0x1e, 0x98, 0x29, 0xc4, 0x09, 0x17, 0x55, 0x12, 0xe6, 0x62, 0xf6, 0x2d, 0x4e, 0x17, 0xe2,
	0xe1, 0x62, 0xcc, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0xcc, 0x4c, 0x62, 0x03, 0xcb,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x77, 0x0c, 0x7b, 0xaf, 0x5b, 0x00, 0x00, 0x00,
}