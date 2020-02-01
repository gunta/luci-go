// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/proto/access/access.proto

package access

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// DescriptionResponse is the response message from Access.Description.
type DescriptionResponse struct {
	// Resources is a list of resource types presented on the given service.
	Resources            []*DescriptionResponse_ResourceDescription `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *DescriptionResponse) Reset()         { *m = DescriptionResponse{} }
func (m *DescriptionResponse) String() string { return proto.CompactTextString(m) }
func (*DescriptionResponse) ProtoMessage()    {}
func (*DescriptionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a2d397fb03f91ad, []int{0}
}

func (m *DescriptionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescriptionResponse.Unmarshal(m, b)
}
func (m *DescriptionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescriptionResponse.Marshal(b, m, deterministic)
}
func (m *DescriptionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescriptionResponse.Merge(m, src)
}
func (m *DescriptionResponse) XXX_Size() int {
	return xxx_messageInfo_DescriptionResponse.Size(m)
}
func (m *DescriptionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescriptionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescriptionResponse proto.InternalMessageInfo

func (m *DescriptionResponse) GetResources() []*DescriptionResponse_ResourceDescription {
	if m != nil {
		return m.Resources
	}
	return nil
}

// ResourceDescription is one resource type, e.g. buildbucket bucket
// or swarming pool.
type DescriptionResponse_ResourceDescription struct {
	// Kind identifies the resource type presented on the service.
	// Access.PermittedActions accepts one of resource kinds.
	// Example: "bucket" for buildbucket bucket, "package" for CIPD package.
	//
	// For implementers:
	// Kind must match regexp `^[a-z\-/]+$`.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Comment provides more info about the resource.
	Comment string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	// Actions defines all possible actions that can be performed on this type
	// of resource.
	//
	// Map key is an action ID, unique within the resource.
	// It is referenced from Role.allowed_actions.
	//
	// For implementers:
	// ActionId must match regexp `^[A-Z\_]+$`.
	// Recommendations:
	// - "READ", not "GET"
	// - "DELETE", not "REMOVE"
	// - prefer concrete actions ("ADD_BUILD", "CHANGE_ACL", "INCREMENT") to
	//   abstract ones ("MODIFY", "WRITE", "UPDATE").
	Actions map[string]*DescriptionResponse_ResourceDescription_Action `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Roles maps a role id to a set of actions.
	// Access configurations are typically expressed with roles, not actions.
	//
	// For implementers:
	// Role IDs must match regexp `^[A-Z\_]+$`.
	// Recommendataion: if it makes sense, make role ID close to the action
	// names, e.g. READER can READ, SCHEDULER can SCHEDULE.
	Roles                map[string]*DescriptionResponse_ResourceDescription_Role `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *DescriptionResponse_ResourceDescription) Reset() {
	*m = DescriptionResponse_ResourceDescription{}
}
func (m *DescriptionResponse_ResourceDescription) String() string { return proto.CompactTextString(m) }
func (*DescriptionResponse_ResourceDescription) ProtoMessage()    {}
func (*DescriptionResponse_ResourceDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a2d397fb03f91ad, []int{0, 0}
}

func (m *DescriptionResponse_ResourceDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescriptionResponse_ResourceDescription.Unmarshal(m, b)
}
func (m *DescriptionResponse_ResourceDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescriptionResponse_ResourceDescription.Marshal(b, m, deterministic)
}
func (m *DescriptionResponse_ResourceDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescriptionResponse_ResourceDescription.Merge(m, src)
}
func (m *DescriptionResponse_ResourceDescription) XXX_Size() int {
	return xxx_messageInfo_DescriptionResponse_ResourceDescription.Size(m)
}
func (m *DescriptionResponse_ResourceDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_DescriptionResponse_ResourceDescription.DiscardUnknown(m)
}

var xxx_messageInfo_DescriptionResponse_ResourceDescription proto.InternalMessageInfo

func (m *DescriptionResponse_ResourceDescription) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *DescriptionResponse_ResourceDescription) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *DescriptionResponse_ResourceDescription) GetActions() map[string]*DescriptionResponse_ResourceDescription_Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *DescriptionResponse_ResourceDescription) GetRoles() map[string]*DescriptionResponse_ResourceDescription_Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

// Action describes what a user can do with a resource.
type DescriptionResponse_ResourceDescription_Action struct {
	// Comment provides more human-readable info about the action.
	Comment              string   `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescriptionResponse_ResourceDescription_Action) Reset() {
	*m = DescriptionResponse_ResourceDescription_Action{}
}
func (m *DescriptionResponse_ResourceDescription_Action) String() string {
	return proto.CompactTextString(m)
}
func (*DescriptionResponse_ResourceDescription_Action) ProtoMessage() {}
func (*DescriptionResponse_ResourceDescription_Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a2d397fb03f91ad, []int{0, 0, 0}
}

func (m *DescriptionResponse_ResourceDescription_Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescriptionResponse_ResourceDescription_Action.Unmarshal(m, b)
}
func (m *DescriptionResponse_ResourceDescription_Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescriptionResponse_ResourceDescription_Action.Marshal(b, m, deterministic)
}
func (m *DescriptionResponse_ResourceDescription_Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescriptionResponse_ResourceDescription_Action.Merge(m, src)
}
func (m *DescriptionResponse_ResourceDescription_Action) XXX_Size() int {
	return xxx_messageInfo_DescriptionResponse_ResourceDescription_Action.Size(m)
}
func (m *DescriptionResponse_ResourceDescription_Action) XXX_DiscardUnknown() {
	xxx_messageInfo_DescriptionResponse_ResourceDescription_Action.DiscardUnknown(m)
}

var xxx_messageInfo_DescriptionResponse_ResourceDescription_Action proto.InternalMessageInfo

func (m *DescriptionResponse_ResourceDescription_Action) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

// Role is a named set of allowed actions.
type DescriptionResponse_ResourceDescription_Role struct {
	// AllowedActions is a set of action IDs.
	// It defines what a role bearer can do with the resource.
	AllowedActions []string `protobuf:"bytes,1,rep,name=allowed_actions,json=allowedActions,proto3" json:"allowed_actions,omitempty"`
	// Comment provides more info about the role.
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescriptionResponse_ResourceDescription_Role) Reset() {
	*m = DescriptionResponse_ResourceDescription_Role{}
}
func (m *DescriptionResponse_ResourceDescription_Role) String() string {
	return proto.CompactTextString(m)
}
func (*DescriptionResponse_ResourceDescription_Role) ProtoMessage() {}
func (*DescriptionResponse_ResourceDescription_Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a2d397fb03f91ad, []int{0, 0, 1}
}

func (m *DescriptionResponse_ResourceDescription_Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescriptionResponse_ResourceDescription_Role.Unmarshal(m, b)
}
func (m *DescriptionResponse_ResourceDescription_Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescriptionResponse_ResourceDescription_Role.Marshal(b, m, deterministic)
}
func (m *DescriptionResponse_ResourceDescription_Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescriptionResponse_ResourceDescription_Role.Merge(m, src)
}
func (m *DescriptionResponse_ResourceDescription_Role) XXX_Size() int {
	return xxx_messageInfo_DescriptionResponse_ResourceDescription_Role.Size(m)
}
func (m *DescriptionResponse_ResourceDescription_Role) XXX_DiscardUnknown() {
	xxx_messageInfo_DescriptionResponse_ResourceDescription_Role.DiscardUnknown(m)
}

var xxx_messageInfo_DescriptionResponse_ResourceDescription_Role proto.InternalMessageInfo

func (m *DescriptionResponse_ResourceDescription_Role) GetAllowedActions() []string {
	if m != nil {
		return m.AllowedActions
	}
	return nil
}

func (m *DescriptionResponse_ResourceDescription_Role) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

// PermittedActionsRequest is a request message to Access.PermittedActions.
//
// Besides explicit fields in the message, there is an implicit parameter: the
// current identity which is defined by the "Authorization" OAuth 2.0 HTTP
// header and, optionally, LUCI-specific delegation token header.
type PermittedActionsRequest struct {
	// ResourceKind is one of Resource.kind values returned by Access.Description.
	// It identifies the type of the resource being checked.
	ResourceKind string `protobuf:"bytes,1,opt,name=resource_kind,json=resourceKind,proto3" json:"resource_kind,omitempty"`
	// ResourceIds identifies the resources presented on this service.
	// For example, for a buildbucket bucket it would be a bucket name
	// ("luci.chromium.try").
	// For a CIPD package it would be a full package name,
	// "infra/git/linux-amd64".
	ResourceIds          []string `protobuf:"bytes,2,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermittedActionsRequest) Reset()         { *m = PermittedActionsRequest{} }
func (m *PermittedActionsRequest) String() string { return proto.CompactTextString(m) }
func (*PermittedActionsRequest) ProtoMessage()    {}
func (*PermittedActionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a2d397fb03f91ad, []int{1}
}

func (m *PermittedActionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermittedActionsRequest.Unmarshal(m, b)
}
func (m *PermittedActionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermittedActionsRequest.Marshal(b, m, deterministic)
}
func (m *PermittedActionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermittedActionsRequest.Merge(m, src)
}
func (m *PermittedActionsRequest) XXX_Size() int {
	return xxx_messageInfo_PermittedActionsRequest.Size(m)
}
func (m *PermittedActionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PermittedActionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PermittedActionsRequest proto.InternalMessageInfo

func (m *PermittedActionsRequest) GetResourceKind() string {
	if m != nil {
		return m.ResourceKind
	}
	return ""
}

func (m *PermittedActionsRequest) GetResourceIds() []string {
	if m != nil {
		return m.ResourceIds
	}
	return nil
}

// PermittedActionsResponse is the response message of the
// Accses.PermittedActions.
type PermittedActionsResponse struct {
	// Permitted maps a resource id to resource permissions.
	Permitted map[string]*PermittedActionsResponse_ResourcePermissions `protobuf:"bytes,1,rep,name=permitted,proto3" json:"permitted,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ValiditiyDuration specifies for how long clients may cache this
	// information.
	ValidityDuration     *duration.Duration `protobuf:"bytes,2,opt,name=validity_duration,json=validityDuration,proto3" json:"validity_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PermittedActionsResponse) Reset()         { *m = PermittedActionsResponse{} }
func (m *PermittedActionsResponse) String() string { return proto.CompactTextString(m) }
func (*PermittedActionsResponse) ProtoMessage()    {}
func (*PermittedActionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a2d397fb03f91ad, []int{2}
}

func (m *PermittedActionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermittedActionsResponse.Unmarshal(m, b)
}
func (m *PermittedActionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermittedActionsResponse.Marshal(b, m, deterministic)
}
func (m *PermittedActionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermittedActionsResponse.Merge(m, src)
}
func (m *PermittedActionsResponse) XXX_Size() int {
	return xxx_messageInfo_PermittedActionsResponse.Size(m)
}
func (m *PermittedActionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PermittedActionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PermittedActionsResponse proto.InternalMessageInfo

func (m *PermittedActionsResponse) GetPermitted() map[string]*PermittedActionsResponse_ResourcePermissions {
	if m != nil {
		return m.Permitted
	}
	return nil
}

func (m *PermittedActionsResponse) GetValidityDuration() *duration.Duration {
	if m != nil {
		return m.ValidityDuration
	}
	return nil
}

// ResourcePermissions describes what is permitted on a single resource.
type PermittedActionsResponse_ResourcePermissions struct {
	// Actions is a list of action ids that the user can do on the resource.
	// For resources that do not exist, this list must be empty.
	Actions              []string `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PermittedActionsResponse_ResourcePermissions) Reset() {
	*m = PermittedActionsResponse_ResourcePermissions{}
}
func (m *PermittedActionsResponse_ResourcePermissions) String() string {
	return proto.CompactTextString(m)
}
func (*PermittedActionsResponse_ResourcePermissions) ProtoMessage() {}
func (*PermittedActionsResponse_ResourcePermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a2d397fb03f91ad, []int{2, 0}
}

func (m *PermittedActionsResponse_ResourcePermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PermittedActionsResponse_ResourcePermissions.Unmarshal(m, b)
}
func (m *PermittedActionsResponse_ResourcePermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PermittedActionsResponse_ResourcePermissions.Marshal(b, m, deterministic)
}
func (m *PermittedActionsResponse_ResourcePermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermittedActionsResponse_ResourcePermissions.Merge(m, src)
}
func (m *PermittedActionsResponse_ResourcePermissions) XXX_Size() int {
	return xxx_messageInfo_PermittedActionsResponse_ResourcePermissions.Size(m)
}
func (m *PermittedActionsResponse_ResourcePermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_PermittedActionsResponse_ResourcePermissions.DiscardUnknown(m)
}

var xxx_messageInfo_PermittedActionsResponse_ResourcePermissions proto.InternalMessageInfo

func (m *PermittedActionsResponse_ResourcePermissions) GetActions() []string {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*DescriptionResponse)(nil), "access.DescriptionResponse")
	proto.RegisterType((*DescriptionResponse_ResourceDescription)(nil), "access.DescriptionResponse.ResourceDescription")
	proto.RegisterMapType((map[string]*DescriptionResponse_ResourceDescription_Action)(nil), "access.DescriptionResponse.ResourceDescription.ActionsEntry")
	proto.RegisterMapType((map[string]*DescriptionResponse_ResourceDescription_Role)(nil), "access.DescriptionResponse.ResourceDescription.RolesEntry")
	proto.RegisterType((*DescriptionResponse_ResourceDescription_Action)(nil), "access.DescriptionResponse.ResourceDescription.Action")
	proto.RegisterType((*DescriptionResponse_ResourceDescription_Role)(nil), "access.DescriptionResponse.ResourceDescription.Role")
	proto.RegisterType((*PermittedActionsRequest)(nil), "access.PermittedActionsRequest")
	proto.RegisterType((*PermittedActionsResponse)(nil), "access.PermittedActionsResponse")
	proto.RegisterMapType((map[string]*PermittedActionsResponse_ResourcePermissions)(nil), "access.PermittedActionsResponse.PermittedEntry")
	proto.RegisterType((*PermittedActionsResponse_ResourcePermissions)(nil), "access.PermittedActionsResponse.ResourcePermissions")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/proto/access/access.proto", fileDescriptor_4a2d397fb03f91ad)
}

var fileDescriptor_4a2d397fb03f91ad = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x5e, 0xba, 0xb6, 0x53, 0x4f, 0xc7, 0x28, 0x9e, 0x04, 0xc1, 0x93, 0xa0, 0x84, 0x0b, 0x76,
	0x95, 0x48, 0xe5, 0x47, 0x68, 0xe2, 0x06, 0xa9, 0x43, 0x1a, 0x3f, 0xd2, 0xe4, 0x0b, 0xb8, 0xac,
	0xb2, 0xc4, 0x14, 0x6b, 0x49, 0x1c, 0x6c, 0x67, 0xa8, 0x4f, 0xb4, 0xd7, 0xe0, 0x29, 0x78, 0x1e,
	0x64, 0x3b, 0x5e, 0xd2, 0xae, 0x05, 0x75, 0x57, 0x89, 0xcf, 0xcf, 0xf7, 0x9d, 0xf3, 0xf9, 0x1c,
	0xc3, 0xeb, 0x39, 0x0f, 0x93, 0x1f, 0x82, 0xe7, 0xac, 0xca, 0x43, 0x2e, 0xe6, 0x51, 0x56, 0x25,
	0x2c, 0x4a, 0x78, 0x9e, 0xf3, 0x22, 0x2a, 0x05, 0x57, 0x3c, 0x8a, 0x93, 0x84, 0x4a, 0x59, 0x7f,
	0x42, 0x63, 0x43, 0x7d, 0x7b, 0xc2, 0x4f, 0xe6, 0x9c, 0xcf, 0x33, 0x6a, 0x23, 0x2f, 0xaa, 0xef,
	0x51, 0x5a, 0x89, 0x58, 0x31, 0x5e, 0xd8, 0x38, 0x7c, 0xb4, 0xea, 0xa7, 0x79, 0xa9, 0x16, 0xd6,
	0x19, 0x5c, 0xf7, 0xe0, 0x70, 0x4a, 0x65, 0x22, 0x58, 0xa9, 0x53, 0x08, 0x95, 0x25, 0x2f, 0x24,
	0x45, 0x5f, 0x60, 0x20, 0xa8, 0xe4, 0x95, 0x48, 0xa8, 0xf4, 0xbd, 0xf1, 0xee, 0xf1, 0x70, 0x12,
	0x85, 0x35, 0xfd, 0x9a, 0xf8, 0x90, 0xd4, 0xc1, 0x6d, 0x5f, 0x83, 0x80, 0x7f, 0x77, 0xe1, 0x70,
	0x4d, 0x08, 0x42, 0xd0, 0xbd, 0x64, 0x45, 0xea, 0x7b, 0x63, 0xef, 0x78, 0x40, 0xcc, 0x3f, 0xf2,
	0x61, 0x4f, 0xf7, 0x4e, 0x0b, 0xe5, 0x77, 0x8c, 0xd9, 0x1d, 0xd1, 0x57, 0xd8, 0x8b, 0x13, 0x9d,
	0x27, 0xfd, 0x5d, 0x53, 0xd2, 0xbb, 0x2d, 0x4b, 0x0a, 0xdf, 0xdb, 0xf4, 0xd3, 0x42, 0x89, 0x05,
	0x71, 0x60, 0xe8, 0x1c, 0x7a, 0x82, 0x67, 0x54, 0xfa, 0x5d, 0x83, 0x7a, 0xb2, 0x2d, 0x2a, 0xd1,
	0xc9, 0x16, 0xd3, 0x02, 0xe1, 0x00, 0xfa, 0x96, 0xaa, 0xdd, 0x8d, 0xb7, 0xd4, 0x0d, 0x3e, 0x83,
	0xae, 0x4e, 0x44, 0x2f, 0xe0, 0x7e, 0x9c, 0x65, 0xfc, 0x17, 0x4d, 0x67, 0xae, 0x3b, 0x2d, 0xf8,
	0x80, 0x1c, 0xd4, 0xe6, 0xba, 0xe8, 0xcd, 0xc2, 0x60, 0x01, 0xfb, 0xed, 0xce, 0xd0, 0x08, 0x76,
	0x2f, 0xe9, 0xa2, 0x26, 0xd4, 0xbf, 0xe8, 0x33, 0xf4, 0xae, 0xe2, 0xac, 0xa2, 0x26, 0x73, 0x38,
	0x79, 0x73, 0x37, 0xe1, 0x88, 0x05, 0x39, 0xe9, 0xbc, 0xf5, 0x70, 0x01, 0xd0, 0xf4, 0xbd, 0x86,
	0xf1, 0xe3, 0x32, 0xe3, 0xab, 0xbb, 0x88, 0xda, 0xe2, 0x0b, 0x62, 0x78, 0x74, 0x4e, 0x45, 0xce,
	0x94, 0xba, 0x51, 0x84, 0xd0, 0x9f, 0x15, 0x95, 0x0a, 0x3d, 0x87, 0x7b, 0x6e, 0xd4, 0x66, 0xad,
	0x71, 0xda, 0x77, 0xc6, 0x4f, 0x7a, 0xac, 0x9e, 0xc1, 0xcd, 0x79, 0xc6, 0x52, 0xe9, 0x77, 0x8c,
	0xc6, 0x43, 0x67, 0x3b, 0x4b, 0x65, 0xf0, 0xa7, 0x03, 0xfe, 0x6d, 0x8e, 0x66, 0x23, 0x4a, 0xe7,
	0x5b, 0xdd, 0x88, 0x4d, 0x49, 0x8d, 0xc3, 0x4e, 0x47, 0x83, 0x80, 0x3e, 0xc0, 0x83, 0xab, 0x38,
	0x63, 0x29, 0x53, 0x8b, 0x99, 0x5b, 0xd8, 0x5a, 0xaa, 0xc7, 0xa1, 0xdd, 0xd8, 0xd0, 0x6d, 0x6c,
	0x38, 0xad, 0x03, 0xc8, 0xc8, 0xe5, 0x38, 0x0b, 0x8e, 0x9a, 0xc5, 0x32, 0x64, 0x52, 0xba, 0x59,
	0x59, 0x1e, 0x26, 0x77, 0xc4, 0x02, 0x0e, 0x96, 0xab, 0xda, 0xe2, 0xee, 0x36, 0xf6, 0xb9, 0xa6,
	0x84, 0xd6, 0xdd, 0x4d, 0xae, 0x3d, 0xbd, 0x0f, 0x1a, 0x02, 0x7d, 0x83, 0xd1, 0x2a, 0x0a, 0x7a,
	0xba, 0x19, 0xdf, 0x5c, 0x30, 0x1e, 0xff, 0xaf, 0x80, 0x60, 0x07, 0x4d, 0x61, 0xd8, 0x7e, 0x59,
	0x1e, 0xde, 0x12, 0xf1, 0x54, 0x3f, 0x7b, 0xf8, 0xe8, 0x1f, 0x73, 0x18, 0xec, 0x5c, 0xf4, 0x4d,
	0xf8, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x35, 0xe4, 0x3d, 0x95, 0x94, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccessClient is the client API for Access service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessClient interface {
	// PermittedActions returns a list of actions the requester can perform
	// on a given resource.
	PermittedActions(ctx context.Context, in *PermittedActionsRequest, opts ...grpc.CallOption) (*PermittedActionsResponse, error)
	// Description returns types of resources and actions that this service
	// supports.
	// It is intended to be used as self-documentation, for humans that play
	// with the API.
	// If the concepts returned by this RPC are internal, it should be restricted.
	Description(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DescriptionResponse, error)
}
type accessPRPCClient struct {
	client *prpc.Client
}

func NewAccessPRPCClient(client *prpc.Client) AccessClient {
	return &accessPRPCClient{client}
}

func (c *accessPRPCClient) PermittedActions(ctx context.Context, in *PermittedActionsRequest, opts ...grpc.CallOption) (*PermittedActionsResponse, error) {
	out := new(PermittedActionsResponse)
	err := c.client.Call(ctx, "access.Access", "PermittedActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessPRPCClient) Description(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DescriptionResponse, error) {
	out := new(DescriptionResponse)
	err := c.client.Call(ctx, "access.Access", "Description", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type accessClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessClient(cc grpc.ClientConnInterface) AccessClient {
	return &accessClient{cc}
}

func (c *accessClient) PermittedActions(ctx context.Context, in *PermittedActionsRequest, opts ...grpc.CallOption) (*PermittedActionsResponse, error) {
	out := new(PermittedActionsResponse)
	err := c.cc.Invoke(ctx, "/access.Access/PermittedActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessClient) Description(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DescriptionResponse, error) {
	out := new(DescriptionResponse)
	err := c.cc.Invoke(ctx, "/access.Access/Description", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessServer is the server API for Access service.
type AccessServer interface {
	// PermittedActions returns a list of actions the requester can perform
	// on a given resource.
	PermittedActions(context.Context, *PermittedActionsRequest) (*PermittedActionsResponse, error)
	// Description returns types of resources and actions that this service
	// supports.
	// It is intended to be used as self-documentation, for humans that play
	// with the API.
	// If the concepts returned by this RPC are internal, it should be restricted.
	Description(context.Context, *empty.Empty) (*DescriptionResponse, error)
}

// UnimplementedAccessServer can be embedded to have forward compatible implementations.
type UnimplementedAccessServer struct {
}

func (*UnimplementedAccessServer) PermittedActions(ctx context.Context, req *PermittedActionsRequest) (*PermittedActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermittedActions not implemented")
}
func (*UnimplementedAccessServer) Description(ctx context.Context, req *empty.Empty) (*DescriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Description not implemented")
}

func RegisterAccessServer(s prpc.Registrar, srv AccessServer) {
	s.RegisterService(&_Access_serviceDesc, srv)
}

func _Access_PermittedActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermittedActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).PermittedActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.Access/PermittedActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).PermittedActions(ctx, req.(*PermittedActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Access_Description_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessServer).Description(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/access.Access/Description",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessServer).Description(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Access_serviceDesc = grpc.ServiceDesc{
	ServiceName: "access.Access",
	HandlerType: (*AccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PermittedActions",
			Handler:    _Access_PermittedActions_Handler,
		},
		{
			MethodName: "Description",
			Handler:    _Access_Description_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/common/proto/access/access.proto",
}
