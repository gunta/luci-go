// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/results/proto/v1/common.proto

package resultspb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Status associated with a result.
type Status int32

const (
	Status_STATUS_UNSPECIFIED Status = 0
	Status_PASS               Status = 1
	Status_FAIL               Status = 2
	Status_CRASH              Status = 3
	Status_ABORT              Status = 4
	Status_SKIP               Status = 5
)

var Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "PASS",
	2: "FAIL",
	3: "CRASH",
	4: "ABORT",
	5: "SKIP",
}

var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"PASS":               1,
	"FAIL":               2,
	"CRASH":              3,
	"ABORT":              4,
	"SKIP":               5,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6721ae6a14ce615e, []int{0}
}

// Bucket definition for specifying ACLs to invocations and results.
// TODO(jchinlee): This will change to "realm".
type Bucket struct {
	Project              string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721ae6a14ce615e, []int{0}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Bucket) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

// Result of the lowest granularity we store.
// This is often a test result, e.g. the result of a functional test run
// within a swarming task.
type Result struct {
	// If unset on upload, set to (Status == PASS).
	IsExpected *wrappers.BoolValue  `protobuf:"bytes,1,opt,name=is_expected,json=isExpected,proto3" json:"is_expected,omitempty"`
	Status     Status               `protobuf:"varint,2,opt,name=status,proto3,enum=luci.resultsdb.Status" json:"status,omitempty"`
	Summary    *Markdown            `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	StartTime  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Duration   *duration.Duration   `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	Tags       []*StringPair        `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	// note: avoid map<> because that would be tens of
	// thousands of tiny hash tables for no good reason.
	Artifacts            []*Artifact `protobuf:"bytes,7,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721ae6a14ce615e, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetIsExpected() *wrappers.BoolValue {
	if m != nil {
		return m.IsExpected
	}
	return nil
}

func (m *Result) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (m *Result) GetSummary() *Markdown {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *Result) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Result) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Result) GetTags() []*StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Result) GetArtifacts() []*Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type VariantDef struct {
	// key regex: ^[a-z][a-z0-9_]*(/[a-z][a-z0-9_]*)*$
	// Max key length: 64
	// value regex: ^[a-z][a-z0-9_]*$
	// Max value length: 64
	//
	// Up to 16 pairs.
	Def map[string]string `protobuf:"bytes,1,rep,name=def,proto3" json:"def,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// hash func:
	// hex(sha256(''.join('%s:%s\n' % (k, v) for k, v in sorted(def.items()))))
	Digest               string   `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VariantDef) Reset()         { *m = VariantDef{} }
func (m *VariantDef) String() string { return proto.CompactTextString(m) }
func (*VariantDef) ProtoMessage()    {}
func (*VariantDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721ae6a14ce615e, []int{2}
}

func (m *VariantDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VariantDef.Unmarshal(m, b)
}
func (m *VariantDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VariantDef.Marshal(b, m, deterministic)
}
func (m *VariantDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VariantDef.Merge(m, src)
}
func (m *VariantDef) XXX_Size() int {
	return xxx_messageInfo_VariantDef.Size(m)
}
func (m *VariantDef) XXX_DiscardUnknown() {
	xxx_messageInfo_VariantDef.DiscardUnknown(m)
}

var xxx_messageInfo_VariantDef proto.InternalMessageInfo

func (m *VariantDef) GetDef() map[string]string {
	if m != nil {
		return m.Def
	}
	return nil
}

func (m *VariantDef) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

// Artifact generated as part of a Result.
//
// Examples include traces, logs, screenshots, memory dumps, profiler output.
type Artifact struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// must start with either of "rbe-cas://" "isolate://", "gs://" or "logdog://".
	FetchUrl string `protobuf:"bytes,2,opt,name=fetch_url,json=fetchUrl,proto3" json:"fetch_url,omitempty"`
	// URL to human-consumable HTTP page with the file content.
	ViewUrl              string   `protobuf:"bytes,3,opt,name=view_url,json=viewUrl,proto3" json:"view_url,omitempty"`
	ContentType          string   `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Size                 int64    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721ae6a14ce615e, []int{3}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetFetchUrl() string {
	if m != nil {
		return m.FetchUrl
	}
	return ""
}

func (m *Artifact) GetViewUrl() string {
	if m != nil {
		return m.ViewUrl
	}
	return ""
}

func (m *Artifact) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Artifact) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// Convenience proto to define key:value pairs.
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
	return fileDescriptor_6721ae6a14ce615e, []int{4}
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

// Markdown-formatted text (see https://spec.commonmark.org/0.28/).
type Markdown struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Markdown) Reset()         { *m = Markdown{} }
func (m *Markdown) String() string { return proto.CompactTextString(m) }
func (*Markdown) ProtoMessage()    {}
func (*Markdown) Descriptor() ([]byte, []int) {
	return fileDescriptor_6721ae6a14ce615e, []int{5}
}

func (m *Markdown) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Markdown.Unmarshal(m, b)
}
func (m *Markdown) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Markdown.Marshal(b, m, deterministic)
}
func (m *Markdown) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Markdown.Merge(m, src)
}
func (m *Markdown) XXX_Size() int {
	return xxx_messageInfo_Markdown.Size(m)
}
func (m *Markdown) XXX_DiscardUnknown() {
	xxx_messageInfo_Markdown.DiscardUnknown(m)
}

var xxx_messageInfo_Markdown proto.InternalMessageInfo

func (m *Markdown) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterEnum("luci.resultsdb.Status", Status_name, Status_value)
	proto.RegisterType((*Bucket)(nil), "luci.resultsdb.Bucket")
	proto.RegisterType((*Result)(nil), "luci.resultsdb.Result")
	proto.RegisterType((*VariantDef)(nil), "luci.resultsdb.VariantDef")
	proto.RegisterMapType((map[string]string)(nil), "luci.resultsdb.VariantDef.DefEntry")
	proto.RegisterType((*Artifact)(nil), "luci.resultsdb.Artifact")
	proto.RegisterType((*StringPair)(nil), "luci.resultsdb.StringPair")
	proto.RegisterType((*Markdown)(nil), "luci.resultsdb.Markdown")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/results/proto/v1/common.proto", fileDescriptor_6721ae6a14ce615e)
}

var fileDescriptor_6721ae6a14ce615e = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x4f, 0xdb, 0x30,
	0x14, 0xc7, 0x57, 0xd2, 0x86, 0xf4, 0x31, 0xa1, 0xc8, 0x9a, 0x50, 0xe8, 0x24, 0xc6, 0xba, 0xcb,
	0xb4, 0x43, 0x22, 0xba, 0x81, 0x36, 0x38, 0xb5, 0xb4, 0x68, 0xd5, 0x7e, 0x55, 0x4e, 0xe1, 0xb0,
	0x4b, 0xe5, 0xa6, 0x6e, 0xc8, 0x48, 0xe2, 0xc8, 0x76, 0x80, 0xee, 0xba, 0xf3, 0x4e, 0xfb, 0x87,
	0x27, 0x3b, 0x0e, 0x68, 0x54, 0x93, 0xb8, 0xbd, 0xe7, 0xf7, 0xf9, 0xbe, 0x67, 0xbf, 0xf7, 0x12,
	0xe8, 0xc5, 0xcc, 0x8f, 0x2e, 0x39, 0xcb, 0x92, 0x32, 0xf3, 0x19, 0x8f, 0x83, 0xb4, 0x8c, 0x92,
	0x80, 0x53, 0x51, 0xa6, 0x52, 0x04, 0x05, 0x67, 0x92, 0x05, 0xd7, 0x07, 0x41, 0xc4, 0xb2, 0x8c,
	0xe5, 0xbe, 0xf6, 0xd1, 0xb6, 0x62, 0x7c, 0xc3, 0x2c, 0xe6, 0x9d, 0xbd, 0x98, 0xb1, 0x38, 0xa5,
	0x15, 0x3d, 0x2f, 0x97, 0xc1, 0xa2, 0xe4, 0x44, 0x26, 0x35, 0xdf, 0x79, 0xf1, 0x30, 0x2e, 0x93,
	0x8c, 0x0a, 0x49, 0xb2, 0xc2, 0x00, 0x6b, 0x09, 0x6e, 0x38, 0x29, 0x0a, 0xca, 0x45, 0x15, 0xef,
	0x1e, 0x83, 0x3d, 0x28, 0xa3, 0x2b, 0x2a, 0x91, 0x07, 0x9b, 0x05, 0x67, 0x3f, 0x68, 0x24, 0xbd,
	0xc6, 0x7e, 0xe3, 0x75, 0x1b, 0xd7, 0x2e, 0xda, 0x01, 0x7b, 0xae, 0x19, 0x6f, 0x43, 0x07, 0x8c,
	0xd7, 0xfd, 0x65, 0x81, 0x8d, 0xf5, 0x55, 0xd1, 0x09, 0x6c, 0x25, 0x62, 0x46, 0x6f, 0x0b, 0x1a,
	0x49, 0xba, 0xd0, 0x09, 0xb6, 0x7a, 0x1d, 0xbf, 0x2a, 0xee, 0xd7, 0xc5, 0xfd, 0x01, 0x63, 0xe9,
	0x05, 0x49, 0x4b, 0x8a, 0x21, 0x11, 0x23, 0x43, 0x23, 0x1f, 0x6c, 0x21, 0x89, 0x2c, 0x85, 0xce,
	0xbf, 0xdd, 0xdb, 0xf1, 0xff, 0xed, 0x82, 0x1f, 0xea, 0x28, 0x36, 0x14, 0xea, 0xc1, 0xa6, 0x28,
	0xb3, 0x8c, 0xf0, 0x95, 0x67, 0xe9, 0x42, 0xde, 0x43, 0xc1, 0x17, 0xc2, 0xaf, 0x16, 0xec, 0x26,
	0xc7, 0x35, 0x88, 0x3e, 0x00, 0x08, 0x49, 0xb8, 0x9c, 0xa9, 0x06, 0x79, 0xcd, 0xff, 0xdc, 0x6f,
	0x5a, 0x77, 0x0f, 0xb7, 0x35, 0xad, 0x7c, 0x74, 0x08, 0x4e, 0xdd, 0x75, 0xaf, 0xa5, 0x85, 0xbb,
	0x6b, 0xc2, 0xa1, 0x01, 0xf0, 0x1d, 0x8a, 0x7c, 0x68, 0x4a, 0x12, 0x0b, 0xcf, 0xde, 0xb7, 0x74,
	0xad, 0xb5, 0x37, 0xf1, 0x24, 0x8f, 0x27, 0x24, 0xe1, 0x58, 0x73, 0xe8, 0x08, 0xda, 0x84, 0xcb,
	0x64, 0x49, 0x22, 0x29, 0xbc, 0x4d, 0x2d, 0x5a, 0x7b, 0x57, 0xdf, 0x00, 0xf8, 0x1e, 0xed, 0xfe,
	0x69, 0x00, 0x5c, 0x10, 0x9e, 0x90, 0x5c, 0x0e, 0xe9, 0x12, 0x1d, 0x82, 0xb5, 0xa0, 0x4b, 0xaf,
	0xa1, 0x13, 0xbc, 0x7a, 0x98, 0xe0, 0x1e, 0xf4, 0x87, 0x74, 0x39, 0xca, 0x25, 0x5f, 0x61, 0xc5,
	0xab, 0x19, 0x2f, 0x92, 0x98, 0x8a, 0xbb, 0x19, 0x57, 0x5e, 0xe7, 0x08, 0x9c, 0x1a, 0x44, 0x2e,
	0x58, 0x57, 0x74, 0x65, 0xb6, 0x43, 0x99, 0xe8, 0x19, 0xb4, 0xae, 0xd5, 0x38, 0x8d, 0xa8, 0x72,
	0x8e, 0x37, 0xde, 0x37, 0xba, 0xbf, 0x1b, 0xe0, 0xd4, 0xb7, 0x45, 0x08, 0x9a, 0x39, 0xc9, 0xa8,
	0x51, 0x6a, 0x1b, 0x3d, 0x87, 0xf6, 0x92, 0xca, 0xe8, 0x72, 0x56, 0xf2, 0xd4, 0xc8, 0x1d, 0x7d,
	0x70, 0xce, 0x53, 0xb4, 0x0b, 0xce, 0x75, 0x42, 0x6f, 0x74, 0xcc, 0xaa, 0x96, 0x51, 0xf9, 0x2a,
	0xf4, 0x12, 0x9e, 0x46, 0x2c, 0x97, 0x34, 0x97, 0x33, 0xb9, 0x2a, 0xaa, 0x51, 0xb6, 0xf1, 0x96,
	0x39, 0x9b, 0xae, 0x0a, 0xaa, 0xca, 0x89, 0xe4, 0x27, 0xd5, 0xc3, 0xb2, 0xb0, 0xb6, 0xbb, 0xef,
	0x00, 0xee, 0x3b, 0xfe, 0xd8, 0x97, 0x74, 0xf7, 0xc0, 0xa9, 0x57, 0x49, 0x65, 0x95, 0xf4, 0xb6,
	0xfe, 0x38, 0xb4, 0xfd, 0x66, 0x0a, 0x76, 0xb5, 0x9b, 0x68, 0x07, 0x50, 0x38, 0xed, 0x4f, 0xcf,
	0xc3, 0xd9, 0xf9, 0xd7, 0x70, 0x32, 0x3a, 0x1d, 0x9f, 0x8d, 0x47, 0x43, 0xf7, 0x09, 0x72, 0xa0,
	0x39, 0xe9, 0x87, 0xa1, 0xdb, 0x50, 0xd6, 0x59, 0x7f, 0xfc, 0xd9, 0xdd, 0x40, 0x6d, 0x68, 0x9d,
	0xe2, 0x7e, 0xf8, 0xd1, 0xb5, 0x94, 0xd9, 0x1f, 0x7c, 0xc3, 0x53, 0xb7, 0xa9, 0xe2, 0xe1, 0xa7,
	0xf1, 0xc4, 0x6d, 0x0d, 0x0e, 0xbe, 0x07, 0x8f, 0xfa, 0x75, 0x9c, 0x98, 0x83, 0x62, 0x3e, 0xb7,
	0xf5, 0xd9, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x3d, 0xc9, 0x23, 0x74, 0x04, 0x00,
	0x00,
}
