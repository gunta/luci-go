// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/common.proto

package buildbucketpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Status of a build or a step.
type Status int32

const (
	// Unspecified state. Meaning depends on the context.
	Status_STATUS_UNSPECIFIED Status = 0
	// Build was scheduled, but did not start or end yet.
	Status_SCHEDULED Status = 1
	// Build/step has started.
	Status_STARTED Status = 2
	// A union of all terminal statuses.
	// Can be used in BuildPredicate.status.
	// A concrete build cannot have this status.
	// Can be used as a bitmask to check that a build ended.
	Status_ENDED_MASK Status = 4
	// A build/step ended successfully.
	Status_SUCCESS Status = 12
	// A build/step ended unsuccessfully due to its Build.Input,
	// e.g. tests failed, and NOT due to a build infrastructure failure.
	Status_FAILURE Status = 20
	// A build/step ended unsuccessfully due to a failure independent of the input,
	// e.g. swarming failed, or the recipe was unable to read the patch from gerrit..
	Status_INFRA_FAILURE Status = 36
	// A build was cancelled explicitly, e.g. via an RPC.
	Status_CANCELED Status = 68
)

var Status_name = map[int32]string{
	0:  "STATUS_UNSPECIFIED",
	1:  "SCHEDULED",
	2:  "STARTED",
	4:  "ENDED_MASK",
	12: "SUCCESS",
	20: "FAILURE",
	36: "INFRA_FAILURE",
	68: "CANCELED",
}

var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"SCHEDULED":          1,
	"STARTED":            2,
	"ENDED_MASK":         4,
	"SUCCESS":            12,
	"FAILURE":            20,
	"INFRA_FAILURE":      36,
	"CANCELED":           68,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{0}
}

// A boolean with an undefined value.
type Trinary int32

const (
	Trinary_UNSET Trinary = 0
	Trinary_YES   Trinary = 1
	Trinary_NO    Trinary = 2
)

var Trinary_name = map[int32]string{
	0: "UNSET",
	1: "YES",
	2: "NO",
}

var Trinary_value = map[string]int32{
	"UNSET": 0,
	"YES":   1,
	"NO":    2,
}

func (x Trinary) String() string {
	return proto.EnumName(Trinary_name, int32(x))
}

func (Trinary) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1}
}

// A Gerrit patchset.
type GerritChange struct {
	// Gerrit hostname, e.g. "chromium-review.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Gerrit project, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Change number, e.g. 12345.
	Change int64 `protobuf:"varint,3,opt,name=change,proto3" json:"change,omitempty"`
	// Patch set number, e.g. 1.
	Patchset             int64    `protobuf:"varint,4,opt,name=patchset,proto3" json:"patchset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GerritChange) Reset()         { *m = GerritChange{} }
func (m *GerritChange) String() string { return proto.CompactTextString(m) }
func (*GerritChange) ProtoMessage()    {}
func (*GerritChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{0}
}
func (m *GerritChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GerritChange.Unmarshal(m, b)
}
func (m *GerritChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GerritChange.Marshal(b, m, deterministic)
}
func (m *GerritChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GerritChange.Merge(m, src)
}
func (m *GerritChange) XXX_Size() int {
	return xxx_messageInfo_GerritChange.Size(m)
}
func (m *GerritChange) XXX_DiscardUnknown() {
	xxx_messageInfo_GerritChange.DiscardUnknown(m)
}

var xxx_messageInfo_GerritChange proto.InternalMessageInfo

func (m *GerritChange) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GerritChange) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GerritChange) GetChange() int64 {
	if m != nil {
		return m.Change
	}
	return 0
}

func (m *GerritChange) GetPatchset() int64 {
	if m != nil {
		return m.Patchset
	}
	return 0
}

// A landed Git commit hosted on Gitiles.
type GitilesCommit struct {
	// Gitiles hostname, e.g. "chromium.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Repository name on the host, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Commit HEX SHA1.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Commit ref, e.g. "refs/heads/master".
	// NOT a branch name: if specified, must start with "refs/".
	Ref string `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
	// Defines a total order of commits on the ref. Requires ref field.
	// Typically 1-based, monotonically increasing, contiguous integer
	// defined by a Gerrit plugin, goto.google.com/git-numberer.
	// TODO(tandrii): make it a public doc.
	Position             uint32   `protobuf:"varint,5,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitilesCommit) Reset()         { *m = GitilesCommit{} }
func (m *GitilesCommit) String() string { return proto.CompactTextString(m) }
func (*GitilesCommit) ProtoMessage()    {}
func (*GitilesCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1}
}
func (m *GitilesCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitilesCommit.Unmarshal(m, b)
}
func (m *GitilesCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitilesCommit.Marshal(b, m, deterministic)
}
func (m *GitilesCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitilesCommit.Merge(m, src)
}
func (m *GitilesCommit) XXX_Size() int {
	return xxx_messageInfo_GitilesCommit.Size(m)
}
func (m *GitilesCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_GitilesCommit.DiscardUnknown(m)
}

var xxx_messageInfo_GitilesCommit proto.InternalMessageInfo

func (m *GitilesCommit) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GitilesCommit) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GitilesCommit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GitilesCommit) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *GitilesCommit) GetPosition() uint32 {
	if m != nil {
		return m.Position
	}
	return 0
}

// A key-value pair of strings.
type StringPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringPair) Reset()         { *m = StringPair{} }
func (m *StringPair) String() string { return proto.CompactTextString(m) }
func (*StringPair) ProtoMessage()    {}
func (*StringPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{2}
}
func (m *StringPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringPair.Unmarshal(m, b)
}
func (m *StringPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringPair.Marshal(b, m, deterministic)
}
func (m *StringPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringPair.Merge(m, src)
}
func (m *StringPair) XXX_Size() int {
	return xxx_messageInfo_StringPair.Size(m)
}
func (m *StringPair) XXX_DiscardUnknown() {
	xxx_messageInfo_StringPair.DiscardUnknown(m)
}

var xxx_messageInfo_StringPair proto.InternalMessageInfo

func (m *StringPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StringPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Half-open time range.
type TimeRange struct {
	// Inclusive lower boundary. Optional.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Exclusive upper boundary. Optional.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{3}
}
func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func init() {
	proto.RegisterType((*GerritChange)(nil), "buildbucket.v2.GerritChange")
	proto.RegisterType((*GitilesCommit)(nil), "buildbucket.v2.GitilesCommit")
	proto.RegisterType((*StringPair)(nil), "buildbucket.v2.StringPair")
	proto.RegisterType((*TimeRange)(nil), "buildbucket.v2.TimeRange")
	proto.RegisterEnum("buildbucket.v2.Status", Status_name, Status_value)
	proto.RegisterEnum("buildbucket.v2.Trinary", Trinary_name, Trinary_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/common.proto", fileDescriptor_a1a0c34bd7fcf0dc)
}

var fileDescriptor_a1a0c34bd7fcf0dc = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6b, 0xdb, 0x4a,
	0x14, 0xc5, 0x23, 0xf9, 0xaf, 0x6e, 0x6c, 0xa3, 0x37, 0x84, 0x20, 0xbc, 0x79, 0xc6, 0xb4, 0x60,
	0xb2, 0x90, 0x21, 0x49, 0x0b, 0xa5, 0x2b, 0x55, 0x1a, 0xa7, 0xa6, 0xa9, 0x1a, 0x66, 0xa4, 0x45,
	0xbb, 0x31, 0xb2, 0x3c, 0x91, 0xa7, 0xb1, 0x34, 0x62, 0x34, 0x0a, 0x84, 0xd2, 0x75, 0xbf, 0x76,
	0xd1, 0xd8, 0x0e, 0xd9, 0xb5, 0xdd, 0xdd, 0xdf, 0xb9, 0xf7, 0xe8, 0xa0, 0xc3, 0xc0, 0x55, 0x26,
	0xdc, 0x74, 0x2b, 0x45, 0xce, 0xeb, 0xdc, 0x15, 0x32, 0x9b, 0xef, 0xea, 0x94, 0xcf, 0xd7, 0x35,
	0xdf, 0x6d, 0xd6, 0x75, 0xfa, 0xc0, 0xd4, 0xbc, 0x94, 0x42, 0x89, 0x79, 0x2a, 0xf2, 0x5c, 0x14,
	0xae, 0x06, 0x34, 0x7a, 0xb1, 0x77, 0x1f, 0x2f, 0xc7, 0xff, 0x67, 0x42, 0x64, 0x3b, 0xb6, 0x3f,
	0x5d, 0xd7, 0xf7, 0x73, 0xc5, 0x73, 0x56, 0xa9, 0x24, 0x2f, 0xf7, 0x86, 0x69, 0x09, 0x83, 0x1b,
	0x26, 0x25, 0x57, 0xfe, 0x36, 0x29, 0x32, 0x86, 0x10, 0xb4, 0xb7, 0xa2, 0x52, 0x8e, 0x31, 0x31,
	0x66, 0x16, 0xd1, 0x33, 0x72, 0xa0, 0x57, 0x4a, 0xf1, 0x9d, 0xa5, 0xca, 0x31, 0xb5, 0x7c, 0x44,
	0x74, 0x0e, 0xdd, 0x54, 0xfb, 0x9c, 0xd6, 0xc4, 0x98, 0xb5, 0xc8, 0x81, 0xd0, 0x18, 0xfa, 0x65,
	0xa2, 0xd2, 0x6d, 0xc5, 0x94, 0xd3, 0xd6, 0x9b, 0x67, 0x9e, 0xfe, 0x80, 0xe1, 0x0d, 0x57, 0x7c,
	0xc7, 0x2a, 0x5f, 0xe4, 0x39, 0x57, 0xff, 0x18, 0x39, 0x02, 0x93, 0x6f, 0x74, 0x9c, 0x45, 0x4c,
	0xbe, 0x41, 0x36, 0xb4, 0x24, 0xbb, 0xd7, 0x29, 0x16, 0x69, 0x46, 0x1d, 0x2e, 0x2a, 0xae, 0xb8,
	0x28, 0x9c, 0xce, 0xc4, 0x98, 0x0d, 0xc9, 0x33, 0x4f, 0xaf, 0x01, 0xa8, 0x92, 0xbc, 0xc8, 0xee,
	0x12, 0x2e, 0x1b, 0xef, 0x03, 0x7b, 0x3a, 0x04, 0x37, 0x23, 0x3a, 0x83, 0xce, 0x63, 0xb2, 0xab,
	0xd9, 0x21, 0x75, 0x0f, 0xd3, 0x9f, 0x60, 0x45, 0x3c, 0x67, 0x44, 0xff, 0xdb, 0x3b, 0x80, 0x4a,
	0x25, 0x52, 0xad, 0x9a, 0x2a, 0xb5, 0xf7, 0xf4, 0x72, 0xec, 0xee, 0x7b, 0x76, 0x8f, 0x3d, 0xbb,
	0xd1, 0xb1, 0x67, 0x62, 0xe9, 0xeb, 0x86, 0xd1, 0x1b, 0xe8, 0xb3, 0x62, 0xb3, 0x37, 0x9a, 0x7f,
	0x34, 0xf6, 0x58, 0xb1, 0x69, 0xe8, 0xe2, 0x97, 0x01, 0x5d, 0xaa, 0x12, 0x55, 0x57, 0xe8, 0x1c,
	0x10, 0x8d, 0xbc, 0x28, 0xa6, 0xab, 0x38, 0xa4, 0x77, 0xd8, 0x5f, 0x2e, 0x96, 0x38, 0xb0, 0x4f,
	0xd0, 0x10, 0x2c, 0xea, 0x7f, 0xc4, 0x41, 0x7c, 0x8b, 0x03, 0xdb, 0x40, 0xa7, 0xd0, 0xa3, 0x91,
	0x47, 0x22, 0x1c, 0xd8, 0x26, 0x1a, 0x01, 0xe0, 0x30, 0xc0, 0xc1, 0xea, 0xb3, 0x47, 0x3f, 0xd9,
	0x6d, 0xbd, 0x8c, 0x7d, 0x1f, 0x53, 0x6a, 0x0f, 0x1a, 0x58, 0x78, 0xcb, 0xdb, 0x98, 0x60, 0xfb,
	0x0c, 0xfd, 0x07, 0xc3, 0x65, 0xb8, 0x20, 0xde, 0xea, 0x28, 0xbd, 0x42, 0x03, 0xe8, 0xfb, 0x5e,
	0xe8, 0xe3, 0xe6, 0xbb, 0xc1, 0xc5, 0x6b, 0xe8, 0x45, 0x92, 0x17, 0x89, 0x7c, 0x42, 0x16, 0x74,
	0xe2, 0x90, 0xe2, 0xc8, 0x3e, 0x41, 0x3d, 0x68, 0x7d, 0xc5, 0xd4, 0x36, 0x50, 0x17, 0xcc, 0xf0,
	0x8b, 0x6d, 0x7e, 0x78, 0xfb, 0xed, 0xfa, 0xef, 0x1e, 0xef, 0xfb, 0x17, 0x4a, 0xb9, 0x5e, 0x77,
	0xb5, 0x78, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x91, 0xbd, 0xf4, 0xa2, 0xfb, 0x02, 0x00, 0x00,
}
