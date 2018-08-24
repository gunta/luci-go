// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/physical_hosts.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v1 "go.chromium.org/luci/machine-db/api/common/v1"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A physical host in the database.
type PhysicalHost struct {
	// The name of this host on the network. Uniquely identifies this host.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The VLAN this host belongs to.
	// When creating a host, omit this field. It will be inferred from the IPv4 address.
	Vlan int64 `protobuf:"varint,2,opt,name=vlan,proto3" json:"vlan,omitempty"`
	// The machine backing this host.
	Machine string `protobuf:"bytes,3,opt,name=machine,proto3" json:"machine,omitempty"`
	// The operating system backing this host.
	Os string `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	// The number of VMs which can be deployed on this host.
	VmSlots int32 `protobuf:"varint,5,opt,name=vm_slots,json=vmSlots,proto3" json:"vm_slots,omitempty"`
	// A description of this host.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// The deployment ticket associated with this host.
	DeploymentTicket string `protobuf:"bytes,7,opt,name=deployment_ticket,json=deploymentTicket,proto3" json:"deployment_ticket,omitempty"`
	// The IPv4 address associated with this host.
	Ipv4 string `protobuf:"bytes,8,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// The state of the machine backing this host.
	State v1.State `protobuf:"varint,9,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	// The virtual datacenter VMs deployed on this host belong to.
	VirtualDatacenter string `protobuf:"bytes,10,opt,name=virtual_datacenter,json=virtualDatacenter,proto3" json:"virtual_datacenter,omitempty"`
	// The primary NIC for this host.
	Nic string `protobuf:"bytes,11,opt,name=nic,proto3" json:"nic,omitempty"`
	// The MAC address associated with the primary NIC.
	MacAddress           string   `protobuf:"bytes,12,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhysicalHost) Reset()         { *m = PhysicalHost{} }
func (m *PhysicalHost) String() string { return proto.CompactTextString(m) }
func (*PhysicalHost) ProtoMessage()    {}
func (*PhysicalHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_physical_hosts_1c486a130640ac96, []int{0}
}
func (m *PhysicalHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhysicalHost.Unmarshal(m, b)
}
func (m *PhysicalHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhysicalHost.Marshal(b, m, deterministic)
}
func (dst *PhysicalHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhysicalHost.Merge(dst, src)
}
func (m *PhysicalHost) XXX_Size() int {
	return xxx_messageInfo_PhysicalHost.Size(m)
}
func (m *PhysicalHost) XXX_DiscardUnknown() {
	xxx_messageInfo_PhysicalHost.DiscardUnknown(m)
}

var xxx_messageInfo_PhysicalHost proto.InternalMessageInfo

func (m *PhysicalHost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PhysicalHost) GetVlan() int64 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *PhysicalHost) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *PhysicalHost) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *PhysicalHost) GetVmSlots() int32 {
	if m != nil {
		return m.VmSlots
	}
	return 0
}

func (m *PhysicalHost) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PhysicalHost) GetDeploymentTicket() string {
	if m != nil {
		return m.DeploymentTicket
	}
	return ""
}

func (m *PhysicalHost) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *PhysicalHost) GetState() v1.State {
	if m != nil {
		return m.State
	}
	return v1.State_STATE_UNSPECIFIED
}

func (m *PhysicalHost) GetVirtualDatacenter() string {
	if m != nil {
		return m.VirtualDatacenter
	}
	return ""
}

func (m *PhysicalHost) GetNic() string {
	if m != nil {
		return m.Nic
	}
	return ""
}

func (m *PhysicalHost) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

// A request to create a new physical host in the database.
type CreatePhysicalHostRequest struct {
	// The host to create in the database.
	Host                 *PhysicalHost `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreatePhysicalHostRequest) Reset()         { *m = CreatePhysicalHostRequest{} }
func (m *CreatePhysicalHostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePhysicalHostRequest) ProtoMessage()    {}
func (*CreatePhysicalHostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_physical_hosts_1c486a130640ac96, []int{1}
}
func (m *CreatePhysicalHostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePhysicalHostRequest.Unmarshal(m, b)
}
func (m *CreatePhysicalHostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePhysicalHostRequest.Marshal(b, m, deterministic)
}
func (dst *CreatePhysicalHostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePhysicalHostRequest.Merge(dst, src)
}
func (m *CreatePhysicalHostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePhysicalHostRequest.Size(m)
}
func (m *CreatePhysicalHostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePhysicalHostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePhysicalHostRequest proto.InternalMessageInfo

func (m *CreatePhysicalHostRequest) GetHost() *PhysicalHost {
	if m != nil {
		return m.Host
	}
	return nil
}

// A request to list physical hosts in the database.
type ListPhysicalHostsRequest struct {
	// The names of hosts to get.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// The VLANs to filter retrieved hosts on.
	Vlans []int64 `protobuf:"varint,2,rep,packed,name=vlans,proto3" json:"vlans,omitempty"`
	// The IPv4 addresses to filter retrieved hosts on.
	Ipv4S []string `protobuf:"bytes,3,rep,name=ipv4s,proto3" json:"ipv4s,omitempty"`
	// The machines to filter retrieved hosts on.
	Machines []string `protobuf:"bytes,4,rep,name=machines,proto3" json:"machines,omitempty"`
	// The operating systems to filter retrieved hosts on.
	Oses []string `protobuf:"bytes,5,rep,name=oses,proto3" json:"oses,omitempty"`
	// The states to filter retrieved hosts on.
	States []v1.State `protobuf:"varint,6,rep,packed,name=states,proto3,enum=common.State" json:"states,omitempty"`
	// The platforms to filter retrieved hosts on.
	Platforms []string `protobuf:"bytes,7,rep,name=platforms,proto3" json:"platforms,omitempty"`
	// The racks to filter retrieved hosts on.
	Racks []string `protobuf:"bytes,8,rep,name=racks,proto3" json:"racks,omitempty"`
	// The datacenters to filter retrieved hosts on.
	Datacenters []string `protobuf:"bytes,9,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
	// The virtual datacenters to filter retrieved hosts on.
	VirtualDatacenters []string `protobuf:"bytes,10,rep,name=virtual_datacenters,json=virtualDatacenters,proto3" json:"virtual_datacenters,omitempty"`
	// The NICs to filter retrieved hosts on.
	Nics []string `protobuf:"bytes,11,rep,name=nics,proto3" json:"nics,omitempty"`
	// The MAC addresses to filter retrieved hosts on.
	MacAddresses         []string `protobuf:"bytes,12,rep,name=mac_addresses,json=macAddresses,proto3" json:"mac_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPhysicalHostsRequest) Reset()         { *m = ListPhysicalHostsRequest{} }
func (m *ListPhysicalHostsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPhysicalHostsRequest) ProtoMessage()    {}
func (*ListPhysicalHostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_physical_hosts_1c486a130640ac96, []int{2}
}
func (m *ListPhysicalHostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPhysicalHostsRequest.Unmarshal(m, b)
}
func (m *ListPhysicalHostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPhysicalHostsRequest.Marshal(b, m, deterministic)
}
func (dst *ListPhysicalHostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPhysicalHostsRequest.Merge(dst, src)
}
func (m *ListPhysicalHostsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPhysicalHostsRequest.Size(m)
}
func (m *ListPhysicalHostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPhysicalHostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPhysicalHostsRequest proto.InternalMessageInfo

func (m *ListPhysicalHostsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetVlans() []int64 {
	if m != nil {
		return m.Vlans
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetIpv4S() []string {
	if m != nil {
		return m.Ipv4S
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetMachines() []string {
	if m != nil {
		return m.Machines
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetOses() []string {
	if m != nil {
		return m.Oses
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetStates() []v1.State {
	if m != nil {
		return m.States
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetPlatforms() []string {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetRacks() []string {
	if m != nil {
		return m.Racks
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetDatacenters() []string {
	if m != nil {
		return m.Datacenters
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetVirtualDatacenters() []string {
	if m != nil {
		return m.VirtualDatacenters
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetNics() []string {
	if m != nil {
		return m.Nics
	}
	return nil
}

func (m *ListPhysicalHostsRequest) GetMacAddresses() []string {
	if m != nil {
		return m.MacAddresses
	}
	return nil
}

// A response containing a list of physical hosts in the database.
type ListPhysicalHostsResponse struct {
	// The hosts matching this request.
	Hosts                []*PhysicalHost `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListPhysicalHostsResponse) Reset()         { *m = ListPhysicalHostsResponse{} }
func (m *ListPhysicalHostsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPhysicalHostsResponse) ProtoMessage()    {}
func (*ListPhysicalHostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_physical_hosts_1c486a130640ac96, []int{3}
}
func (m *ListPhysicalHostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPhysicalHostsResponse.Unmarshal(m, b)
}
func (m *ListPhysicalHostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPhysicalHostsResponse.Marshal(b, m, deterministic)
}
func (dst *ListPhysicalHostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPhysicalHostsResponse.Merge(dst, src)
}
func (m *ListPhysicalHostsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPhysicalHostsResponse.Size(m)
}
func (m *ListPhysicalHostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPhysicalHostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPhysicalHostsResponse proto.InternalMessageInfo

func (m *ListPhysicalHostsResponse) GetHosts() []*PhysicalHost {
	if m != nil {
		return m.Hosts
	}
	return nil
}

// A request to update a physical host in the database.
type UpdatePhysicalHostRequest struct {
	// The host to update in the database.
	Host *PhysicalHost `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The fields to update in the host.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdatePhysicalHostRequest) Reset()         { *m = UpdatePhysicalHostRequest{} }
func (m *UpdatePhysicalHostRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePhysicalHostRequest) ProtoMessage()    {}
func (*UpdatePhysicalHostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_physical_hosts_1c486a130640ac96, []int{4}
}
func (m *UpdatePhysicalHostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePhysicalHostRequest.Unmarshal(m, b)
}
func (m *UpdatePhysicalHostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePhysicalHostRequest.Marshal(b, m, deterministic)
}
func (dst *UpdatePhysicalHostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePhysicalHostRequest.Merge(dst, src)
}
func (m *UpdatePhysicalHostRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePhysicalHostRequest.Size(m)
}
func (m *UpdatePhysicalHostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePhysicalHostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePhysicalHostRequest proto.InternalMessageInfo

func (m *UpdatePhysicalHostRequest) GetHost() *PhysicalHost {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *UpdatePhysicalHostRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func init() {
	proto.RegisterType((*PhysicalHost)(nil), "crimson.PhysicalHost")
	proto.RegisterType((*CreatePhysicalHostRequest)(nil), "crimson.CreatePhysicalHostRequest")
	proto.RegisterType((*ListPhysicalHostsRequest)(nil), "crimson.ListPhysicalHostsRequest")
	proto.RegisterType((*ListPhysicalHostsResponse)(nil), "crimson.ListPhysicalHostsResponse")
	proto.RegisterType((*UpdatePhysicalHostRequest)(nil), "crimson.UpdatePhysicalHostRequest")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/physical_hosts.proto", fileDescriptor_physical_hosts_1c486a130640ac96)
}

var fileDescriptor_physical_hosts_1c486a130640ac96 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6b, 0xdc, 0x3a,
	0x14, 0xc5, 0xe3, 0xcc, 0xd7, 0x75, 0x12, 0x12, 0xbd, 0xf7, 0x40, 0x33, 0x3c, 0x78, 0x66, 0xc2,
	0x83, 0x29, 0x21, 0x36, 0x4d, 0xbb, 0x6a, 0x57, 0x25, 0x25, 0x64, 0xd1, 0x42, 0x71, 0xda, 0xf5,
	0xa0, 0xc8, 0xca, 0x8c, 0x18, 0xcb, 0x72, 0x7d, 0x35, 0x03, 0x59, 0xf7, 0xb7, 0xf6, 0x17, 0xf4,
	0x0f, 0x14, 0x5d, 0x39, 0xc9, 0x94, 0xb4, 0x50, 0xe8, 0xee, 0xde, 0x73, 0x8e, 0x75, 0xa5, 0x73,
	0x64, 0xc1, 0xc5, 0xd2, 0x66, 0x72, 0xd5, 0x5a, 0xa3, 0x37, 0x26, 0xb3, 0xed, 0x32, 0xaf, 0x36,
	0x52, 0xe7, 0x46, 0xc8, 0x95, 0xae, 0xd5, 0x59, 0x79, 0x93, 0x8b, 0x46, 0xe7, 0xb2, 0xd5, 0x06,
	0x6d, 0x9d, 0x6f, 0x9f, 0xe7, 0xcd, 0xea, 0x0e, 0xb5, 0x14, 0xd5, 0x62, 0x65, 0xd1, 0x61, 0xd6,
	0xb4, 0xd6, 0x59, 0x36, 0xec, 0x04, 0xd3, 0x74, 0x69, 0xed, 0xb2, 0x52, 0x39, 0xc1, 0x37, 0x9b,
	0xdb, 0xfc, 0x56, 0xab, 0xaa, 0x5c, 0x18, 0x81, 0xeb, 0x20, 0x9d, 0xbe, 0xfa, 0xad, 0x79, 0xd6,
	0x98, 0x30, 0x0e, 0x9d, 0x70, 0xaa, 0x1b, 0x33, 0xfb, 0xda, 0x83, 0xfd, 0x0f, 0xdd, 0xfc, 0x2b,
	0x8b, 0x8e, 0x31, 0xd8, 0xab, 0x85, 0x51, 0x3c, 0x4a, 0xa3, 0xf9, 0xb8, 0xa0, 0xda, 0x63, 0xdb,
	0x4a, 0xd4, 0xbc, 0x97, 0x46, 0xf3, 0xb8, 0xa0, 0x9a, 0x71, 0x18, 0x76, 0x13, 0x78, 0x4c, 0xd2,
	0xfb, 0x96, 0x1d, 0x42, 0xcf, 0x22, 0xdf, 0x23, 0xb0, 0x67, 0x91, 0x4d, 0x60, 0xb4, 0x35, 0x0b,
	0xac, 0xac, 0x43, 0xde, 0x4f, 0xa3, 0x79, 0xbf, 0x18, 0x6e, 0xcd, 0xb5, 0x6f, 0x59, 0x0a, 0x49,
	0xa9, 0x50, 0xb6, 0xba, 0x71, 0xda, 0xd6, 0x7c, 0x40, 0xdf, 0xec, 0x42, 0xec, 0x14, 0x8e, 0x4b,
	0xd5, 0x54, 0xf6, 0xce, 0xa8, 0xda, 0x2d, 0x9c, 0x96, 0x6b, 0xe5, 0xf8, 0x90, 0x74, 0x47, 0x8f,
	0xc4, 0x47, 0xc2, 0xfd, 0x3e, 0x75, 0xb3, 0x7d, 0xc9, 0x47, 0x61, 0xef, 0xbe, 0x66, 0x27, 0xd0,
	0xa7, 0x03, 0xf3, 0x71, 0x1a, 0xcd, 0x0f, 0xcf, 0x0f, 0xb2, 0x60, 0x44, 0x76, 0xed, 0xc1, 0x22,
	0x70, 0xec, 0x0c, 0xd8, 0x56, 0xb7, 0x6e, 0x23, 0xaa, 0x45, 0x29, 0x9c, 0x90, 0xaa, 0x76, 0xaa,
	0xe5, 0x40, 0xcb, 0x1c, 0x77, 0xcc, 0xdb, 0x07, 0x82, 0x1d, 0x41, 0x5c, 0x6b, 0xc9, 0x13, 0xe2,
	0x7d, 0xc9, 0xfe, 0x83, 0xc4, 0x08, 0xb9, 0x10, 0x65, 0xd9, 0x2a, 0x44, 0xbe, 0x4f, 0x0c, 0x18,
	0x21, 0xdf, 0x04, 0x64, 0x76, 0x09, 0x93, 0x8b, 0x56, 0x09, 0xa7, 0x76, 0xcd, 0x2e, 0xd4, 0xe7,
	0x8d, 0x42, 0xc7, 0x9e, 0xc1, 0x9e, 0x8f, 0x9e, 0x3c, 0x4f, 0xce, 0xff, 0xc9, 0xba, 0xe8, 0xb3,
	0x1f, 0xb4, 0x24, 0x99, 0x7d, 0xeb, 0x01, 0x7f, 0xa7, 0xd1, 0xed, 0x52, 0x78, 0xbf, 0xce, 0xdf,
	0xd0, 0xf7, 0x79, 0x21, 0x8f, 0xd2, 0x78, 0x3e, 0x2e, 0x42, 0xe3, 0x51, 0x9f, 0x18, 0xf2, 0x5e,
	0x1a, 0xcf, 0xe3, 0x22, 0x34, 0x1e, 0xf5, 0xfe, 0x20, 0x8f, 0x83, 0x96, 0x1a, 0x36, 0x85, 0x51,
	0x17, 0xa3, 0x4f, 0xd0, 0x13, 0x0f, 0xbd, 0x77, 0xd7, 0xa2, 0xf2, 0x19, 0x7a, 0x9c, 0x6a, 0xf6,
	0x3f, 0x0c, 0xc2, 0x75, 0xe2, 0x83, 0x34, 0x7e, 0x6a, 0x6f, 0x47, 0xb2, 0x7f, 0x61, 0xdc, 0x54,
	0xc2, 0xdd, 0xda, 0xd6, 0x20, 0x1f, 0xd2, 0xf7, 0x8f, 0x80, 0xdf, 0x4a, 0x2b, 0xe4, 0x1a, 0xf9,
	0x28, 0x6c, 0x85, 0x1a, 0xba, 0x1b, 0x0f, 0x96, 0x23, 0x1f, 0x13, 0xb7, 0x0b, 0xb1, 0x1c, 0xfe,
	0x7a, 0x9a, 0x1a, 0x72, 0x20, 0x25, 0x7b, 0x12, 0x1b, 0x9d, 0xa0, 0xd6, 0x12, 0x79, 0x12, 0x4e,
	0xe0, 0x6b, 0x76, 0x02, 0x07, 0x3b, 0xc9, 0x29, 0x9f, 0x9d, 0x27, 0xf7, 0x1f, 0xb3, 0x53, 0x38,
	0xbb, 0x82, 0xc9, 0x4f, 0x4c, 0xc7, 0xc6, 0xd6, 0xa8, 0xd8, 0x29, 0xf4, 0xe9, 0xc7, 0x25, 0xd7,
	0x7f, 0x19, 0x5f, 0xd0, 0xcc, 0xbe, 0x44, 0x30, 0xf9, 0xd4, 0x94, 0x7f, 0x7c, 0x11, 0xd8, 0x6b,
	0x48, 0x36, 0xb4, 0x0e, 0xbd, 0x04, 0xf4, 0x6b, 0x26, 0xe7, 0xd3, 0x2c, 0x3c, 0x16, 0xd9, 0xfd,
	0x63, 0x91, 0x5d, 0xfa, 0xc7, 0xe2, 0xbd, 0xc0, 0x75, 0x01, 0x41, 0xee, 0xeb, 0x9b, 0x01, 0xf1,
	0x2f, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x18, 0xf1, 0xe9, 0x40, 0xaa, 0x04, 0x00, 0x00,
}
