// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/dm/api/acls/acls.proto

package acls

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

// TODO(iannucci): move these to per-project acls when implementing per-project
// namespaces.
type Acls struct {
	Readers              []string `protobuf:"bytes,1,rep,name=readers,proto3" json:"readers,omitempty"`
	Writers              []string `protobuf:"bytes,2,rep,name=writers,proto3" json:"writers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Acls) Reset()         { *m = Acls{} }
func (m *Acls) String() string { return proto.CompactTextString(m) }
func (*Acls) ProtoMessage()    {}
func (*Acls) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a85e1525a5d8e6b, []int{0}
}
func (m *Acls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Acls.Unmarshal(m, b)
}
func (m *Acls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Acls.Marshal(b, m, deterministic)
}
func (m *Acls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Acls.Merge(m, src)
}
func (m *Acls) XXX_Size() int {
	return xxx_messageInfo_Acls.Size(m)
}
func (m *Acls) XXX_DiscardUnknown() {
	xxx_messageInfo_Acls.DiscardUnknown(m)
}

var xxx_messageInfo_Acls proto.InternalMessageInfo

func (m *Acls) GetReaders() []string {
	if m != nil {
		return m.Readers
	}
	return nil
}

func (m *Acls) GetWriters() []string {
	if m != nil {
		return m.Writers
	}
	return nil
}

func init() {
	proto.RegisterType((*Acls)(nil), "acls.Acls")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/dm/api/acls/acls.proto", fileDescriptor_0a85e1525a5d8e6b)
}

var fileDescriptor_0a85e1525a5d8e6b = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0xc9, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0x4c, 0xce, 0x29, 0x06, 0x13, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x15, 0x17, 0x8b, 0x63, 0x72, 0x4e, 0xb1,
	0x90, 0x04, 0x17, 0x7b, 0x51, 0x6a, 0x62, 0x4a, 0x6a, 0x51, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06,
	0x67, 0x10, 0x8c, 0x0b, 0x92, 0x29, 0x2f, 0xca, 0x2c, 0x01, 0xc9, 0x30, 0x41, 0x64, 0xa0, 0xdc,
	0x24, 0x36, 0xb0, 0x41, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xe9, 0x95, 0x66, 0x77,
	0x00, 0x00, 0x00,
}
