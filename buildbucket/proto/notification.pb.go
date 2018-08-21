// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/notification.proto

package buildbucketpb // import "go.chromium.org/luci/buildbucket/proto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A notification about a build.
type Notification struct {
	// When this notification was created.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Cloud Project ID of the Buildbucket instance that sent this notification,
	// e.g. "cr-buildbucket".
	// Useful if a service listens to both prod and dev instances of buildbucket.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// Buildbucket build id.
	// Use GetBuild rpc to load the contents.
	BuildId int64 `protobuf:"varint,3,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// User-defined opaque blob specified in NotificationConfig.user_data.
	UserData             []byte   `protobuf:"bytes,4,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_notification_f13bc9baf5700659, []int{0}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (dst *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(dst, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Notification) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Notification) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

func (m *Notification) GetUserData() []byte {
	if m != nil {
		return m.UserData
	}
	return nil
}

// Configuration of notifications.
type NotificationConfig struct {
	// Target Cloud PubSub topic.
	// Usually has format "projects/{cloud project}/topics/{topic name}".
	//
	// The PubSub message data is a Notification message in binary format.
	//
	// <buildbucket-app-id>@appspot.gserviceaccount.com must have
	// "pubsub.topics.publish" permissions on the topic, where
	// <buildbucket-app-id> is usually "cr-buildbucket."
	PubsubTopic string `protobuf:"bytes,1,opt,name=pubsub_topic,json=pubsubTopic,proto3" json:"pubsub_topic,omitempty"`
	// Will be available in Notification.user_data.
	// Max length: 4096.
	UserData             []byte   `protobuf:"bytes,2,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationConfig) Reset()         { *m = NotificationConfig{} }
func (m *NotificationConfig) String() string { return proto.CompactTextString(m) }
func (*NotificationConfig) ProtoMessage()    {}
func (*NotificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_notification_f13bc9baf5700659, []int{1}
}
func (m *NotificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationConfig.Unmarshal(m, b)
}
func (m *NotificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationConfig.Marshal(b, m, deterministic)
}
func (dst *NotificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationConfig.Merge(dst, src)
}
func (m *NotificationConfig) XXX_Size() int {
	return xxx_messageInfo_NotificationConfig.Size(m)
}
func (m *NotificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationConfig proto.InternalMessageInfo

func (m *NotificationConfig) GetPubsubTopic() string {
	if m != nil {
		return m.PubsubTopic
	}
	return ""
}

func (m *NotificationConfig) GetUserData() []byte {
	if m != nil {
		return m.UserData
	}
	return nil
}

func init() {
	proto.RegisterType((*Notification)(nil), "buildbucket.v2.Notification")
	proto.RegisterType((*NotificationConfig)(nil), "buildbucket.v2.NotificationConfig")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/notification.proto", fileDescriptor_notification_f13bc9baf5700659)
}

var fileDescriptor_notification_f13bc9baf5700659 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0xe5, 0x16, 0x4a, 0xe3, 0x56, 0x0c, 0x96, 0x90, 0x42, 0x19, 0x08, 0x9d, 0x32, 0x39,
	0x52, 0x41, 0x08, 0xc4, 0x06, 0x2c, 0x5d, 0x18, 0xa2, 0x4c, 0x2c, 0x91, 0x1d, 0x27, 0xe6, 0x89,
	0x24, 0xcf, 0x4a, 0x6c, 0x7e, 0x85, 0xdf, 0x45, 0x71, 0x54, 0x48, 0x37, 0x46, 0x1f, 0xdf, 0x7b,
	0x74, 0xf5, 0xe8, 0xa3, 0x46, 0x5e, 0x7c, 0x74, 0xd8, 0x80, 0x6b, 0x38, 0x76, 0x3a, 0xa9, 0x5d,
	0x01, 0x89, 0x74, 0x50, 0x2b, 0xe9, 0x8a, 0xcf, 0xd2, 0x26, 0xa6, 0x43, 0x8b, 0x49, 0x8b, 0x16,
	0x2a, 0x28, 0x84, 0x05, 0x6c, 0xb9, 0x47, 0xec, 0x7c, 0x92, 0xe2, 0x5f, 0xbb, 0xcd, 0xb5, 0x46,
	0xd4, 0x75, 0x39, 0x16, 0xa4, 0xab, 0x12, 0x0b, 0x4d, 0xd9, 0x5b, 0xd1, 0x98, 0xb1, 0xb0, 0xfd,
	0x26, 0x74, 0xfd, 0x36, 0xf1, 0xb0, 0x07, 0x1a, 0xfc, 0x66, 0x42, 0x12, 0x91, 0x78, 0xb5, 0xdb,
	0xf0, 0xd1, 0xc2, 0x0f, 0x16, 0x9e, 0x1d, 0x12, 0xe9, 0x5f, 0x98, 0x5d, 0xd0, 0x85, 0x30, 0x26,
	0x07, 0x15, 0xce, 0x22, 0x12, 0x07, 0xe9, 0xa9, 0x30, 0x66, 0xaf, 0xd8, 0x25, 0x5d, 0xfa, 0x51,
	0xc3, 0xc7, 0x3c, 0x22, 0xf1, 0x3c, 0x3d, 0xf3, 0xef, 0xbd, 0x62, 0x57, 0x34, 0x70, 0x7d, 0xd9,
	0xe5, 0x4a, 0x58, 0x11, 0x9e, 0x44, 0x24, 0x5e, 0xa7, 0xcb, 0x01, 0xbc, 0x0a, 0x2b, 0xb6, 0x19,
	0x65, 0xd3, 0x61, 0x2f, 0xd8, 0x56, 0xa0, 0xd9, 0x0d, 0x5d, 0x1b, 0x27, 0x7b, 0x27, 0x73, 0x8b,
	0x06, 0x0a, 0xbf, 0x30, 0x48, 0x57, 0x23, 0xcb, 0x06, 0x74, 0x6c, 0x9d, 0x1d, 0x5b, 0x9f, 0xef,
	0xdf, 0xef, 0xfe, 0x77, 0xdd, 0xa7, 0x09, 0x31, 0x52, 0x2e, 0x3c, 0xbc, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x77, 0x43, 0x5c, 0x47, 0x9c, 0x01, 0x00, 0x00,
}
