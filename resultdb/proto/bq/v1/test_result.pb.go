// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/bq/v1/test_result.proto

package bqpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "go.chromium.org/luci/resultdb/proto/rpc/v1"
	_type "go.chromium.org/luci/resultdb/proto/type"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Represents a test result row in a BigQuery table.
type TestResultRow struct {
	// Invocation level information.
	Invocation *TestResultRow_Invocation `protobuf:"bytes,1,opt,name=invocation,proto3" json:"invocation,omitempty"`
	// A result of a functional test case.
	// Refer to ../../rpc/v1/test_result.proto for definition.
	Result *v1.TestResult `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	// Indicates that the test subject (e.g. a CL) is absolved from blame
	// for this result.
	Exoneration          *TestResultRow_TestExoneration `protobuf:"bytes,3,opt,name=exoneration,proto3" json:"exoneration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TestResultRow) Reset()         { *m = TestResultRow{} }
func (m *TestResultRow) String() string { return proto.CompactTextString(m) }
func (*TestResultRow) ProtoMessage()    {}
func (*TestResultRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_45e9665d0ffb1b0a, []int{0}
}

func (m *TestResultRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultRow.Unmarshal(m, b)
}
func (m *TestResultRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultRow.Marshal(b, m, deterministic)
}
func (m *TestResultRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultRow.Merge(m, src)
}
func (m *TestResultRow) XXX_Size() int {
	return xxx_messageInfo_TestResultRow.Size(m)
}
func (m *TestResultRow) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultRow.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultRow proto.InternalMessageInfo

func (m *TestResultRow) GetInvocation() *TestResultRow_Invocation {
	if m != nil {
		return m.Invocation
	}
	return nil
}

func (m *TestResultRow) GetResult() *v1.TestResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *TestResultRow) GetExoneration() *TestResultRow_TestExoneration {
	if m != nil {
		return m.Exoneration
	}
	return nil
}

// A subset of luci.resultdb.rpc.v1.Invocation message
// in ../../rpc/v1/invocation.proto.
type TestResultRow_Invocation struct {
	// Id of the exported invocation.
	// Note that it's possible that this invocation is not the result's
	// immediate parent invocation, but the including invocation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags                 []*_type.StringPair `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TestResultRow_Invocation) Reset()         { *m = TestResultRow_Invocation{} }
func (m *TestResultRow_Invocation) String() string { return proto.CompactTextString(m) }
func (*TestResultRow_Invocation) ProtoMessage()    {}
func (*TestResultRow_Invocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_45e9665d0ffb1b0a, []int{0, 0}
}

func (m *TestResultRow_Invocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultRow_Invocation.Unmarshal(m, b)
}
func (m *TestResultRow_Invocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultRow_Invocation.Marshal(b, m, deterministic)
}
func (m *TestResultRow_Invocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultRow_Invocation.Merge(m, src)
}
func (m *TestResultRow_Invocation) XXX_Size() int {
	return xxx_messageInfo_TestResultRow_Invocation.Size(m)
}
func (m *TestResultRow_Invocation) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultRow_Invocation.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultRow_Invocation proto.InternalMessageInfo

func (m *TestResultRow_Invocation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TestResultRow_Invocation) GetTags() []*_type.StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

type TestResultRow_TestExoneration struct {
	// True if the test subject (e.g. a CL) is absolved from blame for this
	// result, otherwise False.
	// For more details, refer to luci.resultdb.rpc.v1.TestExoneration
	// in ../../rpc/v1/test_result.proto.
	Exonerated           bool     `protobuf:"varint,1,opt,name=exonerated,proto3" json:"exonerated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResultRow_TestExoneration) Reset()         { *m = TestResultRow_TestExoneration{} }
func (m *TestResultRow_TestExoneration) String() string { return proto.CompactTextString(m) }
func (*TestResultRow_TestExoneration) ProtoMessage()    {}
func (*TestResultRow_TestExoneration) Descriptor() ([]byte, []int) {
	return fileDescriptor_45e9665d0ffb1b0a, []int{0, 1}
}

func (m *TestResultRow_TestExoneration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResultRow_TestExoneration.Unmarshal(m, b)
}
func (m *TestResultRow_TestExoneration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResultRow_TestExoneration.Marshal(b, m, deterministic)
}
func (m *TestResultRow_TestExoneration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResultRow_TestExoneration.Merge(m, src)
}
func (m *TestResultRow_TestExoneration) XXX_Size() int {
	return xxx_messageInfo_TestResultRow_TestExoneration.Size(m)
}
func (m *TestResultRow_TestExoneration) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResultRow_TestExoneration.DiscardUnknown(m)
}

var xxx_messageInfo_TestResultRow_TestExoneration proto.InternalMessageInfo

func (m *TestResultRow_TestExoneration) GetExonerated() bool {
	if m != nil {
		return m.Exonerated
	}
	return false
}

func init() {
	proto.RegisterType((*TestResultRow)(nil), "luci.resultdb.bq.v1.TestResultRow")
	proto.RegisterType((*TestResultRow_Invocation)(nil), "luci.resultdb.bq.v1.TestResultRow.Invocation")
	proto.RegisterType((*TestResultRow_TestExoneration)(nil), "luci.resultdb.bq.v1.TestResultRow.TestExoneration")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/bq/v1/test_result.proto", fileDescriptor_45e9665d0ffb1b0a)
}

var fileDescriptor_45e9665d0ffb1b0a = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x59, 0x27, 0x43, 0xdf, 0x50, 0x21, 0x5e, 0x46, 0x0f, 0x63, 0x78, 0xda, 0xc5, 0xc4,
	0x56, 0x04, 0x61, 0x9e, 0x04, 0x0f, 0x1e, 0x84, 0x11, 0x77, 0xf2, 0x22, 0x4d, 0x1a, 0x6a, 0x60,
	0xed, 0x6b, 0xd3, 0xb4, 0xea, 0xe7, 0xf6, 0x0b, 0x48, 0x53, 0xe7, 0x6a, 0xd9, 0xa1, 0xc7, 0xbe,
	0xfe, 0x7f, 0xbf, 0xf7, 0x4f, 0x08, 0xac, 0x12, 0xa4, 0xf2, 0xdd, 0x60, 0xaa, 0xab, 0x94, 0xa2,
	0x49, 0xd8, 0xb6, 0x92, 0x9a, 0x19, 0x55, 0x56, 0x5b, 0x1b, 0x0b, 0x96, 0x1b, 0xb4, 0xc8, 0x44,
	0xc1, 0xea, 0x80, 0x59, 0x55, 0xda, 0xb7, 0xf6, 0x0f, 0x75, 0x73, 0x72, 0xd1, 0x84, 0xe9, 0x2e,
	0x4c, 0x45, 0x41, 0xeb, 0xc0, 0xbf, 0x1f, 0x62, 0x34, 0xb9, 0x3c, 0xa8, 0xf4, 0x6f, 0x87, 0xd0,
	0xf6, 0x2b, 0x57, 0x4c, 0x62, 0x9a, 0x62, 0xd6, 0x62, 0x97, 0xdf, 0x1e, 0x9c, 0x6e, 0x54, 0x69,
	0xb9, 0x0b, 0x72, 0xfc, 0x20, 0xcf, 0x00, 0x3a, 0xab, 0x51, 0x46, 0x56, 0x63, 0x36, 0x1b, 0x2d,
	0x46, 0xcb, 0x69, 0x78, 0x45, 0x0f, 0x14, 0xa6, 0xff, 0x38, 0xfa, 0xf4, 0x07, 0xf1, 0x8e, 0x80,
	0xdc, 0xc1, 0xa4, 0xc5, 0x66, 0x9e, 0x53, 0x2d, 0x7a, 0x2a, 0x93, 0xcb, 0x9e, 0xeb, 0x37, 0x4f,
	0x36, 0x30, 0x55, 0x9f, 0x98, 0x29, 0xd3, 0x36, 0x19, 0x3b, 0x3c, 0x1c, 0xd0, 0xa4, 0xf9, 0x7a,
	0xdc, 0x93, 0xbc, 0xab, 0xf1, 0xd7, 0x00, 0xfb, 0xa6, 0xe4, 0x0c, 0x3c, 0x1d, 0xbb, 0x43, 0x9e,
	0x70, 0x4f, 0xc7, 0x24, 0x84, 0x23, 0x1b, 0x25, 0xe5, 0xcc, 0x5b, 0x8c, 0x97, 0xd3, 0x70, 0xde,
	0x5b, 0xd6, 0x5c, 0x1f, 0x7d, 0xb1, 0x46, 0x67, 0xc9, 0x3a, 0xd2, 0x86, 0xbb, 0xac, 0x1f, 0xc0,
	0x79, 0x6f, 0x23, 0x99, 0x03, 0xec, 0x76, 0xaa, 0x56, 0x7f, 0xcc, 0x3b, 0x93, 0x87, 0xeb, 0x57,
	0x3a, 0xf8, 0xf9, 0xac, 0x44, 0x91, 0x0b, 0x31, 0x71, 0x93, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0xab, 0x08, 0x43, 0x77, 0x02, 0x00, 0x00,
}
