// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/dm/api/service/v1/activate_execution.proto

package dm

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

// ActivateExecutionReq allows a currently-running Execution to activate itself.
// Doing this allows DM to know that the Execution has started, and also enables
// the Execution to access other APIs like WalkGraph, AddDeps, and
// FinishAttempt.
//
// ActivateExecution must be called with the ExecutionID and Activation token
// that DM provided when the Execution was started with the distributor.
//
// If the Execution has not been activated, the Execution will be marked as
// 'activating' and this will return an OK code. At this point, your client
// may use the ExecutionToken with any RPCs that have an ExecutionAuth field.
//
// This RPC may return:
//   * OK - The Execution is now activated.
//   * InvalidArgmument - The request was malformed. Retrying will not help.
//   * PermissionDenied - The provided activation token was incorrect.
//     Retrying will not help.
//   * AlreadyExists - The activation token was correct, but some other entity
//     already activated this Execution. The client should cease operations.
//     Retrying will not help.
//
// All other errors should be retried with the exact same ActivateExecutionReq
// data.
type ActivateExecutionReq struct {
	// Auth is the Execution_Auth containing the Activation Token, as provided
	// to the distributor when the Execution was created.
	Auth *Execution_Auth `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	// ExecutionToken should be randomly generated by the machine running the
	// execution, or by the distributor service such that if two racing Executions
	// both attempt to Activate with the same ExecutionID and ActivationToken, the
	// ExecutionToken will (probably) be different for them so that only one will
	// win.
	ExecutionToken       []byte   `protobuf:"bytes,2,opt,name=execution_token,json=executionToken,proto3" json:"execution_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivateExecutionReq) Reset()         { *m = ActivateExecutionReq{} }
func (m *ActivateExecutionReq) String() string { return proto.CompactTextString(m) }
func (*ActivateExecutionReq) ProtoMessage()    {}
func (*ActivateExecutionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9b5cee97bd0a18f, []int{0}
}
func (m *ActivateExecutionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateExecutionReq.Unmarshal(m, b)
}
func (m *ActivateExecutionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateExecutionReq.Marshal(b, m, deterministic)
}
func (m *ActivateExecutionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateExecutionReq.Merge(m, src)
}
func (m *ActivateExecutionReq) XXX_Size() int {
	return xxx_messageInfo_ActivateExecutionReq.Size(m)
}
func (m *ActivateExecutionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateExecutionReq.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateExecutionReq proto.InternalMessageInfo

func (m *ActivateExecutionReq) GetAuth() *Execution_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *ActivateExecutionReq) GetExecutionToken() []byte {
	if m != nil {
		return m.ExecutionToken
	}
	return nil
}

func init() {
	proto.RegisterType((*ActivateExecutionReq)(nil), "dm.ActivateExecutionReq")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/dm/api/service/v1/activate_execution.proto", fileDescriptor_c9b5cee97bd0a18f)
}

var fileDescriptor_c9b5cee97bd0a18f = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x46, 0x69, 0x11, 0x87, 0x28, 0x0a, 0xc1, 0xa1, 0x38, 0x15, 0x07, 0xed, 0x94, 0xa0, 0x0e,
	0x8e, 0xd2, 0xc1, 0x3f, 0x50, 0xdc, 0x4b, 0x4c, 0x42, 0x12, 0x34, 0xbd, 0x1a, 0x2f, 0xc5, 0x9f,
	0x2f, 0x56, 0xed, 0xec, 0xfa, 0x1e, 0xf7, 0xee, 0x23, 0x47, 0x03, 0x4c, 0xda, 0x00, 0xde, 0x45,
	0xcf, 0x20, 0x18, 0x7e, 0x8b, 0xd2, 0x71, 0xe5, 0xb9, 0x68, 0x1d, 0x7f, 0xe8, 0xd0, 0x39, 0xa9,
	0x79, 0xb7, 0xe5, 0x42, 0xa2, 0xeb, 0x04, 0xea, 0x5a, 0x3f, 0xb5, 0x8c, 0xe8, 0xa0, 0x61, 0x6d,
	0x00, 0x04, 0x9a, 0x2a, 0xbf, 0x3c, 0xfc, 0x19, 0x31, 0x41, 0xb4, 0xb6, 0x56, 0x02, 0xc5, 0xe7,
	0x78, 0x65, 0xc8, 0xa2, 0xfc, 0x86, 0x4f, 0xbf, 0x6e, 0xa5, 0xef, 0x74, 0x4d, 0x46, 0x22, 0xa2,
	0xcd, 0x92, 0x3c, 0x29, 0x26, 0x3b, 0xca, 0x94, 0x67, 0x83, 0x67, 0x65, 0x44, 0x5b, 0xf5, 0x9e,
	0x6e, 0xc8, 0x7c, 0xd8, 0x53, 0x23, 0x5c, 0x75, 0x93, 0xa5, 0x79, 0x52, 0x4c, 0xab, 0xd9, 0x80,
	0xcf, 0x6f, 0x7a, 0x19, 0xf7, 0xff, 0xf6, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xdd, 0xcd,
	0xcb, 0xef, 0x00, 0x00, 0x00,
}
