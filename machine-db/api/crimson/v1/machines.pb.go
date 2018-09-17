// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/machines.proto

package crimson

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "go.chromium.org/luci/machine-db/api/common/v1"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// A machine in the database.
type Machine struct {
	// The name of this machine. Uniquely identifies this machine.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of platform this machine is.
	Platform string `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	// The rack this machine belongs to.
	Rack string `protobuf:"bytes,3,opt,name=rack,proto3" json:"rack,omitempty"`
	// A description of this machine.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// The asset tag associated with this machine.
	AssetTag string `protobuf:"bytes,5,opt,name=asset_tag,json=assetTag,proto3" json:"asset_tag,omitempty"`
	// The service tag associated with this machine.
	ServiceTag string `protobuf:"bytes,6,opt,name=service_tag,json=serviceTag,proto3" json:"service_tag,omitempty"`
	// The deployment ticket associated with this machine.
	DeploymentTicket string `protobuf:"bytes,7,opt,name=deployment_ticket,json=deploymentTicket,proto3" json:"deployment_ticket,omitempty"`
	// The state of this machine.
	State v1.State `protobuf:"varint,8,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	// The datacenter this machine belongs to.
	// When creating a machine, omit this field. It will be inferred from the rack.
	Datacenter string `protobuf:"bytes,9,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	// The DRAC password associated with this machine.
	DracPassword         string   `protobuf:"bytes,10,opt,name=drac_password,json=dracPassword,proto3" json:"drac_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Machine) Reset()         { *m = Machine{} }
func (m *Machine) String() string { return proto.CompactTextString(m) }
func (*Machine) ProtoMessage()    {}
func (*Machine) Descriptor() ([]byte, []int) {
	return fileDescriptor_17aa8e73038b4213, []int{0}
}
func (m *Machine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Machine.Unmarshal(m, b)
}
func (m *Machine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Machine.Marshal(b, m, deterministic)
}
func (m *Machine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Machine.Merge(m, src)
}
func (m *Machine) XXX_Size() int {
	return xxx_messageInfo_Machine.Size(m)
}
func (m *Machine) XXX_DiscardUnknown() {
	xxx_messageInfo_Machine.DiscardUnknown(m)
}

var xxx_messageInfo_Machine proto.InternalMessageInfo

func (m *Machine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Machine) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *Machine) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *Machine) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Machine) GetAssetTag() string {
	if m != nil {
		return m.AssetTag
	}
	return ""
}

func (m *Machine) GetServiceTag() string {
	if m != nil {
		return m.ServiceTag
	}
	return ""
}

func (m *Machine) GetDeploymentTicket() string {
	if m != nil {
		return m.DeploymentTicket
	}
	return ""
}

func (m *Machine) GetState() v1.State {
	if m != nil {
		return m.State
	}
	return v1.State_STATE_UNSPECIFIED
}

func (m *Machine) GetDatacenter() string {
	if m != nil {
		return m.Datacenter
	}
	return ""
}

func (m *Machine) GetDracPassword() string {
	if m != nil {
		return m.DracPassword
	}
	return ""
}

// A request to create a new machine in the database.
type CreateMachineRequest struct {
	// The machine to create in the database.
	Machine              *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMachineRequest) Reset()         { *m = CreateMachineRequest{} }
func (m *CreateMachineRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMachineRequest) ProtoMessage()    {}
func (*CreateMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17aa8e73038b4213, []int{1}
}
func (m *CreateMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMachineRequest.Unmarshal(m, b)
}
func (m *CreateMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMachineRequest.Marshal(b, m, deterministic)
}
func (m *CreateMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMachineRequest.Merge(m, src)
}
func (m *CreateMachineRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMachineRequest.Size(m)
}
func (m *CreateMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMachineRequest proto.InternalMessageInfo

func (m *CreateMachineRequest) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

// A request to delete a machine from the database.
type DeleteMachineRequest struct {
	// The name of the machine to delete.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMachineRequest) Reset()         { *m = DeleteMachineRequest{} }
func (m *DeleteMachineRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMachineRequest) ProtoMessage()    {}
func (*DeleteMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17aa8e73038b4213, []int{2}
}
func (m *DeleteMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMachineRequest.Unmarshal(m, b)
}
func (m *DeleteMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMachineRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMachineRequest.Merge(m, src)
}
func (m *DeleteMachineRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMachineRequest.Size(m)
}
func (m *DeleteMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMachineRequest proto.InternalMessageInfo

func (m *DeleteMachineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request to list machines in the database.
type ListMachinesRequest struct {
	// The names of machines to get.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// The platforms to filter retrieved machines on.
	Platforms []string `protobuf:"bytes,2,rep,name=platforms,proto3" json:"platforms,omitempty"`
	// The racks to filter retrieved machines on.
	Racks []string `protobuf:"bytes,3,rep,name=racks,proto3" json:"racks,omitempty"`
	// The states to filter retrieved machines on.
	States []v1.State `protobuf:"varint,4,rep,packed,name=states,proto3,enum=common.State" json:"states,omitempty"`
	// The datacenters to filter retrieved machines on.
	Datacenters          []string `protobuf:"bytes,5,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMachinesRequest) Reset()         { *m = ListMachinesRequest{} }
func (m *ListMachinesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMachinesRequest) ProtoMessage()    {}
func (*ListMachinesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17aa8e73038b4213, []int{3}
}
func (m *ListMachinesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMachinesRequest.Unmarshal(m, b)
}
func (m *ListMachinesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMachinesRequest.Marshal(b, m, deterministic)
}
func (m *ListMachinesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMachinesRequest.Merge(m, src)
}
func (m *ListMachinesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMachinesRequest.Size(m)
}
func (m *ListMachinesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMachinesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMachinesRequest proto.InternalMessageInfo

func (m *ListMachinesRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListMachinesRequest) GetPlatforms() []string {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func (m *ListMachinesRequest) GetRacks() []string {
	if m != nil {
		return m.Racks
	}
	return nil
}

func (m *ListMachinesRequest) GetStates() []v1.State {
	if m != nil {
		return m.States
	}
	return nil
}

func (m *ListMachinesRequest) GetDatacenters() []string {
	if m != nil {
		return m.Datacenters
	}
	return nil
}

// A response containing a list of machines in the database.
type ListMachinesResponse struct {
	// The machines matching this request.
	Machines             []*Machine `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListMachinesResponse) Reset()         { *m = ListMachinesResponse{} }
func (m *ListMachinesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMachinesResponse) ProtoMessage()    {}
func (*ListMachinesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17aa8e73038b4213, []int{4}
}
func (m *ListMachinesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMachinesResponse.Unmarshal(m, b)
}
func (m *ListMachinesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMachinesResponse.Marshal(b, m, deterministic)
}
func (m *ListMachinesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMachinesResponse.Merge(m, src)
}
func (m *ListMachinesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMachinesResponse.Size(m)
}
func (m *ListMachinesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMachinesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMachinesResponse proto.InternalMessageInfo

func (m *ListMachinesResponse) GetMachines() []*Machine {
	if m != nil {
		return m.Machines
	}
	return nil
}

// A request to rename a machine in the database.
type RenameMachineRequest struct {
	// The name of the machine to rename.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The new name to give this machine.
	NewName              string   `protobuf:"bytes,2,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameMachineRequest) Reset()         { *m = RenameMachineRequest{} }
func (m *RenameMachineRequest) String() string { return proto.CompactTextString(m) }
func (*RenameMachineRequest) ProtoMessage()    {}
func (*RenameMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17aa8e73038b4213, []int{5}
}
func (m *RenameMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameMachineRequest.Unmarshal(m, b)
}
func (m *RenameMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameMachineRequest.Marshal(b, m, deterministic)
}
func (m *RenameMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameMachineRequest.Merge(m, src)
}
func (m *RenameMachineRequest) XXX_Size() int {
	return xxx_messageInfo_RenameMachineRequest.Size(m)
}
func (m *RenameMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenameMachineRequest proto.InternalMessageInfo

func (m *RenameMachineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RenameMachineRequest) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

// A request to update a machine in the database.
type UpdateMachineRequest struct {
	// The machine to update in the database.
	Machine *Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	// The fields to update in the machine.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateMachineRequest) Reset()         { *m = UpdateMachineRequest{} }
func (m *UpdateMachineRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMachineRequest) ProtoMessage()    {}
func (*UpdateMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17aa8e73038b4213, []int{6}
}
func (m *UpdateMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMachineRequest.Unmarshal(m, b)
}
func (m *UpdateMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMachineRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMachineRequest.Merge(m, src)
}
func (m *UpdateMachineRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMachineRequest.Size(m)
}
func (m *UpdateMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMachineRequest proto.InternalMessageInfo

func (m *UpdateMachineRequest) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

func (m *UpdateMachineRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func init() {
	proto.RegisterType((*Machine)(nil), "crimson.Machine")
	proto.RegisterType((*CreateMachineRequest)(nil), "crimson.CreateMachineRequest")
	proto.RegisterType((*DeleteMachineRequest)(nil), "crimson.DeleteMachineRequest")
	proto.RegisterType((*ListMachinesRequest)(nil), "crimson.ListMachinesRequest")
	proto.RegisterType((*ListMachinesResponse)(nil), "crimson.ListMachinesResponse")
	proto.RegisterType((*RenameMachineRequest)(nil), "crimson.RenameMachineRequest")
	proto.RegisterType((*UpdateMachineRequest)(nil), "crimson.UpdateMachineRequest")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/machines.proto", fileDescriptor_17aa8e73038b4213)
}

var fileDescriptor_17aa8e73038b4213 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x55, 0xd7, 0xf5, 0xdf, 0xed, 0x6f, 0x3f, 0x0d, 0x93, 0x07, 0x53, 0x10, 0x44, 0x99, 0x90,
	0xaa, 0x01, 0x89, 0x28, 0x6f, 0x20, 0x5e, 0x60, 0xf0, 0xc4, 0x10, 0x0a, 0xe3, 0xb9, 0x72, 0x93,
	0xdb, 0x2c, 0x6a, 0x12, 0x07, 0xdb, 0x59, 0xc5, 0x13, 0x5f, 0x87, 0x0f, 0xc2, 0x07, 0x43, 0xbe,
	0x4e, 0xda, 0xa1, 0x4d, 0x02, 0x89, 0x37, 0xfb, 0x9c, 0xe3, 0x7b, 0xed, 0x73, 0x7c, 0xe1, 0x75,
	0x26, 0xc3, 0xe4, 0x52, 0xc9, 0x32, 0x6f, 0xca, 0x50, 0xaa, 0x2c, 0x2a, 0x9a, 0x24, 0x8f, 0x4a,
	0x91, 0x5c, 0xe6, 0x15, 0x3e, 0x4b, 0x57, 0x91, 0xa8, 0xf3, 0x28, 0x51, 0x79, 0xa9, 0x65, 0x15,
	0x5d, 0x3d, 0xef, 0x18, 0x1d, 0xd6, 0x4a, 0x1a, 0xc9, 0x46, 0x2d, 0x35, 0xf3, 0x33, 0x29, 0xb3,
	0x02, 0x23, 0x82, 0x57, 0xcd, 0x3a, 0x5a, 0xe7, 0x58, 0xa4, 0xcb, 0x52, 0xe8, 0x8d, 0x93, 0xce,
	0x5e, 0xfe, 0x55, 0x27, 0x59, 0x96, 0xae, 0x91, 0x36, 0xc2, 0x74, 0x6d, 0x82, 0x9f, 0x07, 0x30,
	0x3a, 0x77, 0x4a, 0xc6, 0xe0, 0xb0, 0x12, 0x25, 0xf2, 0x9e, 0xdf, 0x9b, 0x4f, 0x62, 0x5a, 0xb3,
	0x19, 0x8c, 0xeb, 0x42, 0x98, 0xb5, 0x54, 0x25, 0x3f, 0x20, 0x7c, 0xb7, 0xb7, 0x7a, 0x25, 0x92,
	0x0d, 0xef, 0x3b, 0xbd, 0x5d, 0x33, 0x1f, 0xa6, 0x29, 0xea, 0x44, 0xe5, 0xb5, 0xc9, 0x65, 0xc5,
	0x0f, 0x89, 0xba, 0x0e, 0xb1, 0xfb, 0x30, 0x11, 0x5a, 0xa3, 0x59, 0x1a, 0x91, 0xf1, 0x81, 0x2b,
	0x49, 0xc0, 0x85, 0xc8, 0xd8, 0x23, 0x98, 0x6a, 0x54, 0x57, 0x79, 0x82, 0x44, 0x0f, 0x89, 0x86,
	0x16, 0xb2, 0x82, 0x27, 0x70, 0x27, 0xc5, 0xba, 0x90, 0xdf, 0x4a, 0xac, 0xcc, 0xd2, 0xe4, 0xc9,
	0x06, 0x0d, 0x1f, 0x91, 0xec, 0x78, 0x4f, 0x5c, 0x10, 0xce, 0x4e, 0x60, 0x40, 0x8f, 0xe5, 0x63,
	0xbf, 0x37, 0xff, 0x7f, 0x71, 0x14, 0x3a, 0x13, 0xc2, 0xcf, 0x16, 0x8c, 0x1d, 0xc7, 0x1e, 0x02,
	0xa4, 0xc2, 0x88, 0x04, 0x2b, 0x83, 0x8a, 0x4f, 0x5c, 0xc7, 0x3d, 0xc2, 0x4e, 0xe0, 0x28, 0x55,
	0x22, 0x59, 0xd6, 0x42, 0xeb, 0xad, 0x54, 0x29, 0x07, 0x92, 0xfc, 0x67, 0xc1, 0x4f, 0x2d, 0x16,
	0xbc, 0x01, 0xef, 0xad, 0x42, 0x61, 0xb0, 0xf5, 0x32, 0xc6, 0xaf, 0x0d, 0x6a, 0xc3, 0x4e, 0x61,
	0xd4, 0xe6, 0x40, 0xae, 0x4e, 0x17, 0xc7, 0x61, 0x9b, 0x6b, 0xd8, 0x29, 0x3b, 0x41, 0x70, 0x0a,
	0xde, 0x19, 0x16, 0x78, 0xa3, 0xc6, 0x2d, 0xb1, 0x04, 0x3f, 0x7a, 0x70, 0xf7, 0x43, 0xae, 0x4d,
	0x2b, 0xd5, 0x9d, 0xd6, 0x83, 0x81, 0xe5, 0x35, 0xef, 0xf9, 0xfd, 0xf9, 0x24, 0x76, 0x1b, 0xf6,
	0x00, 0x26, 0x5d, 0x68, 0x9a, 0x1f, 0x10, 0xb3, 0x07, 0xec, 0x19, 0x1b, 0x9d, 0xe6, 0x7d, 0x77,
	0x86, 0x36, 0xec, 0x31, 0x0c, 0xdd, 0x47, 0xe1, 0x87, 0x7e, 0xff, 0xa6, 0x79, 0x2d, 0x49, 0x79,
	0xef, 0xbc, 0xd2, 0x7c, 0x40, 0x25, 0xae, 0x43, 0xc1, 0x19, 0x78, 0xbf, 0xdf, 0x54, 0xd7, 0xb2,
	0xd2, 0xc8, 0x9e, 0xc2, 0xb8, 0xfb, 0xf2, 0x74, 0xdb, 0xdb, 0xbc, 0xd9, 0x29, 0x82, 0x77, 0xe0,
	0xc5, 0x68, 0x5f, 0xf3, 0x67, 0x73, 0xd8, 0x3d, 0x18, 0x57, 0xb8, 0x5d, 0x12, 0xee, 0xfe, 0xec,
	0xa8, 0xc2, 0xed, 0x47, 0xeb, 0xdb, 0x77, 0xf0, 0xbe, 0xd4, 0xe9, 0x3f, 0xe5, 0xc4, 0x5e, 0xc1,
	0xb4, 0xa1, 0x1a, 0x34, 0x83, 0xd4, 0x61, 0xba, 0x98, 0x85, 0x6e, 0x4c, 0xc3, 0x6e, 0x4c, 0xc3,
	0xf7, 0x76, 0x4c, 0xcf, 0x85, 0xde, 0xc4, 0xd0, 0xb4, 0x2d, 0xf5, 0x66, 0x35, 0x24, 0xfe, 0xc5,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xeb, 0xc4, 0xd6, 0x1e, 0x04, 0x00, 0x00,
}
