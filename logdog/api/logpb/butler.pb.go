// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/logdog/api/logpb/butler.proto

package logpb

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

//
// This enumerates the possible contents of published Butler data.
type ButlerMetadata_ContentType int32

const (
	// An invalid content type. Do not use.
	ButlerMetadata_Invalid ButlerMetadata_ContentType = 0
	// The published data is a ButlerLogBundle protobuf message.
	ButlerMetadata_ButlerLogBundle ButlerMetadata_ContentType = 1
)

var ButlerMetadata_ContentType_name = map[int32]string{
	0: "Invalid",
	1: "ButlerLogBundle",
}

var ButlerMetadata_ContentType_value = map[string]int32{
	"Invalid":         0,
	"ButlerLogBundle": 1,
}

func (x ButlerMetadata_ContentType) String() string {
	return proto.EnumName(ButlerMetadata_ContentType_name, int32(x))
}

func (ButlerMetadata_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_25f3936477fa1e93, []int{0, 0}
}

// Compression scheme of attached data.
type ButlerMetadata_Compression int32

const (
	ButlerMetadata_NONE ButlerMetadata_Compression = 0
	ButlerMetadata_ZLIB ButlerMetadata_Compression = 1
)

var ButlerMetadata_Compression_name = map[int32]string{
	0: "NONE",
	1: "ZLIB",
}

var ButlerMetadata_Compression_value = map[string]int32{
	"NONE": 0,
	"ZLIB": 1,
}

func (x ButlerMetadata_Compression) String() string {
	return proto.EnumName(ButlerMetadata_Compression_name, int32(x))
}

func (ButlerMetadata_Compression) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_25f3936477fa1e93, []int{0, 1}
}

//
// ButlerMetadata appears as a frame at the beginning of Butler published data
// to describe the remainder of the contents.
type ButlerMetadata struct {
	// This is the type of data in the subsequent frame.
	Type        ButlerMetadata_ContentType `protobuf:"varint,1,opt,name=type,proto3,enum=logpb.ButlerMetadata_ContentType" json:"type,omitempty"`
	Compression ButlerMetadata_Compression `protobuf:"varint,2,opt,name=compression,proto3,enum=logpb.ButlerMetadata_Compression" json:"compression,omitempty"`
	// The protobuf version string (see version.go).
	ProtoVersion         string   `protobuf:"bytes,3,opt,name=proto_version,json=protoVersion,proto3" json:"proto_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ButlerMetadata) Reset()         { *m = ButlerMetadata{} }
func (m *ButlerMetadata) String() string { return proto.CompactTextString(m) }
func (*ButlerMetadata) ProtoMessage()    {}
func (*ButlerMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_25f3936477fa1e93, []int{0}
}
func (m *ButlerMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ButlerMetadata.Unmarshal(m, b)
}
func (m *ButlerMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ButlerMetadata.Marshal(b, m, deterministic)
}
func (m *ButlerMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ButlerMetadata.Merge(m, src)
}
func (m *ButlerMetadata) XXX_Size() int {
	return xxx_messageInfo_ButlerMetadata.Size(m)
}
func (m *ButlerMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ButlerMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ButlerMetadata proto.InternalMessageInfo

func (m *ButlerMetadata) GetType() ButlerMetadata_ContentType {
	if m != nil {
		return m.Type
	}
	return ButlerMetadata_Invalid
}

func (m *ButlerMetadata) GetCompression() ButlerMetadata_Compression {
	if m != nil {
		return m.Compression
	}
	return ButlerMetadata_NONE
}

func (m *ButlerMetadata) GetProtoVersion() string {
	if m != nil {
		return m.ProtoVersion
	}
	return ""
}

//
// A message containing log data in transit from the Butler.
//
// The Butler is capable of conserving bandwidth by bundling collected log
// messages together into this protocol buffer. Based on Butler bundling
// settings, this message can represent anything from a single LogRecord to
// multiple LogRecords belonging to several different streams.
//
// Entries in a Log Bundle are fully self-descriptive: no additional information
// is needed to fully associate the contained data with its proper place in
// the source log stream.
type ButlerLogBundle struct {
	//
	// (DEPRECATED) Stream source information. Now supplied during prefix
	// registration.
	DeprecatedSource string `protobuf:"bytes,1,opt,name=deprecated_source,json=deprecatedSource,proto3" json:"deprecated_source,omitempty"`
	// The timestamp when this bundle was generated.
	//
	// This field will be used for debugging and internal accounting.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// *
	// Each Entry is an individual set of log records for a given log stream.
	Entries []*ButlerLogBundle_Entry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
	// * Project specifies which luci-config project this stream belongs to.
	Project string `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
	// *
	// The log stream prefix that is shared by all bundled streams.
	//
	// This prefix is valid within the supplied project scope.
	Prefix string `protobuf:"bytes,5,opt,name=prefix,proto3" json:"prefix,omitempty"`
	//
	// The log prefix's secret value (required).
	//
	// The secret is bound to all log streams that share the supplied Prefix, and
	// The Coordinator will record the secret associated with a given log Prefix,
	// but will not expose the secret to users.
	//
	// The Collector will check the secret prior to ingesting logs. If the
	// secret doesn't match the value recorded by the Coordinator, the log
	// will be discarded.
	//
	// This ensures that only the Butler instance that generated the log stream
	// can emit log data for that stream. It also ensures that only authenticated
	// users can write to a Prefix.
	Secret               []byte   `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ButlerLogBundle) Reset()         { *m = ButlerLogBundle{} }
func (m *ButlerLogBundle) String() string { return proto.CompactTextString(m) }
func (*ButlerLogBundle) ProtoMessage()    {}
func (*ButlerLogBundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_25f3936477fa1e93, []int{1}
}
func (m *ButlerLogBundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ButlerLogBundle.Unmarshal(m, b)
}
func (m *ButlerLogBundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ButlerLogBundle.Marshal(b, m, deterministic)
}
func (m *ButlerLogBundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ButlerLogBundle.Merge(m, src)
}
func (m *ButlerLogBundle) XXX_Size() int {
	return xxx_messageInfo_ButlerLogBundle.Size(m)
}
func (m *ButlerLogBundle) XXX_DiscardUnknown() {
	xxx_messageInfo_ButlerLogBundle.DiscardUnknown(m)
}

var xxx_messageInfo_ButlerLogBundle proto.InternalMessageInfo

func (m *ButlerLogBundle) GetDeprecatedSource() string {
	if m != nil {
		return m.DeprecatedSource
	}
	return ""
}

func (m *ButlerLogBundle) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ButlerLogBundle) GetEntries() []*ButlerLogBundle_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *ButlerLogBundle) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ButlerLogBundle) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ButlerLogBundle) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

//
// A bundle Entry describes a set of LogEntry messages originating from the
// same log stream.
type ButlerLogBundle_Entry struct {
	//
	// The descriptor for this entry's log stream.
	//
	// Each LogEntry in the "logs" field is shares this common descriptor.
	Desc *LogStreamDescriptor `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	// (DEPRECATED) Per-entry secret replaced with Butler-wide secret.
	DeprecatedEntrySecret []byte `protobuf:"bytes,2,opt,name=deprecated_entry_secret,json=deprecatedEntrySecret,proto3" json:"deprecated_entry_secret,omitempty"`
	//
	// Whether this log entry terminates its stream.
	//
	// If present and "true", this field declares that this Entry is the last
	// such entry in the stream. This fact is recorded by the Collector and
	// registered with the Coordinator. The largest stream prefix in this Entry
	// will be bound the stream's LogEntry records to [0:largest_prefix]. Once
	// all messages in that range have been received, the log may be archived.
	//
	// Further log entries belonging to this stream with stream indices
	// exceeding the terminal log's index will be discarded.
	Terminal bool `protobuf:"varint,3,opt,name=terminal,proto3" json:"terminal,omitempty"`
	//
	// If terminal is true, this is the terminal stream index; that is, the last
	// message index in the stream.
	TerminalIndex uint64 `protobuf:"varint,4,opt,name=terminal_index,json=terminalIndex,proto3" json:"terminal_index,omitempty"`
	//
	// Log entries attached to this record. These MUST be sequential.
	//
	// This is the main log entry content.
	Logs                 []*LogEntry `protobuf:"bytes,5,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ButlerLogBundle_Entry) Reset()         { *m = ButlerLogBundle_Entry{} }
func (m *ButlerLogBundle_Entry) String() string { return proto.CompactTextString(m) }
func (*ButlerLogBundle_Entry) ProtoMessage()    {}
func (*ButlerLogBundle_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_25f3936477fa1e93, []int{1, 0}
}
func (m *ButlerLogBundle_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ButlerLogBundle_Entry.Unmarshal(m, b)
}
func (m *ButlerLogBundle_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ButlerLogBundle_Entry.Marshal(b, m, deterministic)
}
func (m *ButlerLogBundle_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ButlerLogBundle_Entry.Merge(m, src)
}
func (m *ButlerLogBundle_Entry) XXX_Size() int {
	return xxx_messageInfo_ButlerLogBundle_Entry.Size(m)
}
func (m *ButlerLogBundle_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_ButlerLogBundle_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_ButlerLogBundle_Entry proto.InternalMessageInfo

func (m *ButlerLogBundle_Entry) GetDesc() *LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *ButlerLogBundle_Entry) GetDeprecatedEntrySecret() []byte {
	if m != nil {
		return m.DeprecatedEntrySecret
	}
	return nil
}

func (m *ButlerLogBundle_Entry) GetTerminal() bool {
	if m != nil {
		return m.Terminal
	}
	return false
}

func (m *ButlerLogBundle_Entry) GetTerminalIndex() uint64 {
	if m != nil {
		return m.TerminalIndex
	}
	return 0
}

func (m *ButlerLogBundle_Entry) GetLogs() []*LogEntry {
	if m != nil {
		return m.Logs
	}
	return nil
}

func init() {
	proto.RegisterType((*ButlerMetadata)(nil), "logpb.ButlerMetadata")
	proto.RegisterType((*ButlerLogBundle)(nil), "logpb.ButlerLogBundle")
	proto.RegisterType((*ButlerLogBundle_Entry)(nil), "logpb.ButlerLogBundle.Entry")
	proto.RegisterEnum("logpb.ButlerMetadata_ContentType", ButlerMetadata_ContentType_name, ButlerMetadata_ContentType_value)
	proto.RegisterEnum("logpb.ButlerMetadata_Compression", ButlerMetadata_Compression_name, ButlerMetadata_Compression_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/logdog/api/logpb/butler.proto", fileDescriptor_25f3936477fa1e93)
}

var fileDescriptor_25f3936477fa1e93 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x8d, 0x6c, 0xf9, 0x6b, 0x9c, 0x38, 0xee, 0x96, 0xb6, 0xc2, 0x14, 0xea, 0x38, 0x14, 0x0c,
	0x85, 0x15, 0xb8, 0x34, 0xf4, 0xec, 0x34, 0x07, 0x83, 0x9b, 0x82, 0x1c, 0x7a, 0xe8, 0xc5, 0xc8,
	0xd2, 0x64, 0xbb, 0x45, 0xd2, 0x2e, 0xab, 0x55, 0x88, 0x7f, 0x40, 0xff, 0x61, 0xa1, 0x7f, 0xa7,
	0x68, 0x64, 0xd9, 0x0e, 0x85, 0xd2, 0xdb, 0xce, 0x9b, 0x37, 0xf3, 0xde, 0xbc, 0x85, 0x99, 0x50,
	0x3c, 0xfa, 0x6e, 0x54, 0x2a, 0x8b, 0x94, 0x2b, 0x23, 0xfc, 0xa4, 0x88, 0xa4, 0x9f, 0x28, 0x11,
	0x2b, 0xe1, 0x87, 0x9a, 0x9e, 0x7a, 0xe3, 0x6f, 0x0a, 0x9b, 0xa0, 0xe1, 0xda, 0x28, 0xab, 0x58,
	0x8b, 0xb0, 0x91, 0xff, 0x7f, 0xa3, 0x89, 0x12, 0xd5, 0xdc, 0xe8, 0x8d, 0x50, 0x4a, 0x24, 0xe8,
	0x53, 0xb5, 0x29, 0xee, 0x7d, 0x2b, 0x53, 0xcc, 0x6d, 0x98, 0xea, 0x8a, 0x30, 0xf9, 0xd9, 0x80,
	0xc1, 0x9c, 0x94, 0x3e, 0xa3, 0x0d, 0xe3, 0xd0, 0x86, 0xec, 0x03, 0xb8, 0x76, 0xab, 0xd1, 0x73,
	0xc6, 0xce, 0x74, 0x30, 0xbb, 0xe0, 0xb4, 0x93, 0x3f, 0x25, 0xf1, 0x6b, 0x95, 0x59, 0xcc, 0xec,
	0xdd, 0x56, 0x63, 0x40, 0x74, 0x76, 0x0d, 0xfd, 0x48, 0xa5, 0xda, 0x60, 0x9e, 0x4b, 0x95, 0x79,
	0x8d, 0x7f, 0x4f, 0xef, 0x89, 0xc1, 0xf1, 0x14, 0xbb, 0x84, 0x33, 0xf2, 0xb5, 0x7e, 0x40, 0x43,
	0x6b, 0x9a, 0x63, 0x67, 0xda, 0x0b, 0x4e, 0x09, 0xfc, 0x5a, 0x61, 0x13, 0x1f, 0xfa, 0x47, 0xf2,
	0xac, 0x0f, 0x9d, 0x45, 0xf6, 0x10, 0x26, 0x32, 0x1e, 0x9e, 0xb0, 0xe7, 0x70, 0x5e, 0x69, 0x2d,
	0x95, 0x98, 0x17, 0x59, 0x9c, 0xe0, 0xd0, 0x99, 0x5c, 0x94, 0x03, 0x07, 0x91, 0x2e, 0xb8, 0xb7,
	0x5f, 0x6e, 0x6f, 0x86, 0x27, 0xe5, 0xeb, 0xdb, 0x72, 0x31, 0x1f, 0x3a, 0x93, 0x5f, 0xcd, 0xbf,
	0x06, 0xd9, 0x3b, 0x78, 0x16, 0xa3, 0x36, 0x18, 0x85, 0x16, 0xe3, 0x75, 0xae, 0x0a, 0x13, 0x55,
	0xa9, 0xf4, 0x82, 0xe1, 0xa1, 0xb1, 0x22, 0x9c, 0x7d, 0x84, 0xde, 0x3e, 0x5b, 0x3a, 0xbe, 0x3f,
	0x1b, 0xf1, 0x2a, 0x7d, 0x5e, 0xa7, 0xcf, 0xef, 0x6a, 0x46, 0x70, 0x20, 0xb3, 0x2b, 0xe8, 0x60,
	0x66, 0x8d, 0xc4, 0xdc, 0x6b, 0x8e, 0x9b, 0xd3, 0xfe, 0xec, 0xf5, 0x93, 0xd0, 0xf6, 0x7e, 0xf8,
	0x4d, 0x66, 0xcd, 0x36, 0xa8, 0xc9, 0xcc, 0x83, 0x8e, 0x36, 0xea, 0x07, 0x46, 0xd6, 0x73, 0xc9,
	0x54, 0x5d, 0xb2, 0x97, 0xd0, 0xd6, 0x06, 0xef, 0xe5, 0xa3, 0xd7, 0xa2, 0xc6, 0xae, 0x2a, 0xf1,
	0x1c, 0x23, 0x83, 0xd6, 0x6b, 0x8f, 0x9d, 0xe9, 0x69, 0xb0, 0xab, 0x46, 0xbf, 0x1d, 0x68, 0xd1,
	0x72, 0xc6, 0xc1, 0x8d, 0x31, 0x8f, 0xe8, 0xca, 0xf2, 0x80, 0xca, 0xc8, 0x52, 0x89, 0x95, 0x35,
	0x18, 0xa6, 0x9f, 0x30, 0x8f, 0x8c, 0xd4, 0x56, 0x99, 0x80, 0x78, 0xec, 0x0a, 0x5e, 0x1d, 0x45,
	0x54, 0x3a, 0xdb, 0xae, 0x77, 0x12, 0x0d, 0x92, 0x78, 0x71, 0x68, 0x93, 0xc2, 0x8a, 0x9a, 0x6c,
	0x04, 0x5d, 0x8b, 0x26, 0x95, 0x59, 0x98, 0xd0, 0x17, 0x77, 0x83, 0x7d, 0xcd, 0xde, 0xc2, 0xa0,
	0x7e, 0xaf, 0x65, 0x16, 0xe3, 0x23, 0x9d, 0xe7, 0x06, 0x67, 0x35, 0xba, 0x28, 0x41, 0x76, 0x09,
	0x6e, 0xa2, 0x44, 0xee, 0xb5, 0x28, 0xb3, 0xf3, 0x83, 0xd5, 0x2a, 0x26, 0x6a, 0x6e, 0xda, 0x14,
	0xfd, 0xfb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0x53, 0x1b, 0x37, 0x74, 0x03, 0x00, 0x00,
}
