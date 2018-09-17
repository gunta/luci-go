// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/vpython/api/vpython/spec.proto

package vpython

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

// Spec is a "vpython" environment specification.
type Spec struct {
	// The Python version to use. This should be of the form:
	// "Major[.Minor[.Patch]]"
	//
	// If specified,
	// - The Major version will be enforced absolutely. Python 3 will not be
	//   preferred over Python 2 because '3' is greater than '2'.
	// - The remaining versions, if specified, will be regarded as *minimum*
	//   versions. In other words, if "2.7.4" is specified and the system has
	//   "2.7.12", that will suffice. Similarly, "2.6" would accept a "2.7"
	//   interpreter.
	//
	// If empty, the default Python interpreter ("python") will be used.
	PythonVersion string          `protobuf:"bytes,1,opt,name=python_version,json=pythonVersion,proto3" json:"python_version,omitempty"`
	Wheel         []*Spec_Package `protobuf:"bytes,2,rep,name=wheel,proto3" json:"wheel,omitempty"`
	// The VirtualEnv package.
	//
	// This should be left empty to use the `vpython` default package
	// (recommended).
	Virtualenv *Spec_Package `protobuf:"bytes,3,opt,name=virtualenv,proto3" json:"virtualenv,omitempty"`
	// Specification-provided PEP425 verification tags.
	//
	// By default, verification will be performed against a default set of
	// environment parameters. However, a given specification may offer its own
	// set of PEP425 tags representing the systems that it wants to be verified
	// against.
	VerifyPep425Tag      []*PEP425Tag `protobuf:"bytes,4,rep,name=verify_pep425_tag,json=verifyPep425Tag,proto3" json:"verify_pep425_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_12b41745b49e8c72, []int{0}
}
func (m *Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spec.Unmarshal(m, b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
}
func (m *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(m, src)
}
func (m *Spec) XXX_Size() int {
	return xxx_messageInfo_Spec.Size(m)
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

func (m *Spec) GetPythonVersion() string {
	if m != nil {
		return m.PythonVersion
	}
	return ""
}

func (m *Spec) GetWheel() []*Spec_Package {
	if m != nil {
		return m.Wheel
	}
	return nil
}

func (m *Spec) GetVirtualenv() *Spec_Package {
	if m != nil {
		return m.Virtualenv
	}
	return nil
}

func (m *Spec) GetVerifyPep425Tag() []*PEP425Tag {
	if m != nil {
		return m.VerifyPep425Tag
	}
	return nil
}

// A definition for a remote package. The type of package depends on the
// configured package resolver.
type Spec_Package struct {
	// The name of the package.
	//
	// - For CIPD, this is the package name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The package version.
	//
	// - For CIPD, this will be any recognized CIPD version (i.e., ID, tag, or
	//   ref).
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Optional PEP425 tags to determine whether this package is included on the
	// target system. If no match tags are specified, this package will always
	// be included. If match tags are specified, the package will be included if
	// any system PEP425 tags match at least one of the match tags.
	//
	// A match will succeed if any system PEP425 tag field matches the
	// corresponding field in the PEP425 tag. If the match tag omits a field
	// (partial), that field will not be considered. For example, if a match
	// tag specifies just an ABI field, any system PEP425 tag with that ABI will
	// be considered a successful match, regardless of other field values.
	MatchTag []*PEP425Tag `protobuf:"bytes,3,rep,name=match_tag,json=matchTag,proto3" json:"match_tag,omitempty"`
	// Optional PEP425 tags to determine whether this package is NOT included on
	// the target system. This has the opposite behavior as "match_tag": if any
	// host tags match any tags in this list, the package will not be installed
	// on this host.
	//
	// A "not_match_tag" overrides a "match_tag", so if a host has tags that
	// match entries in both, the package will be not considered a match.
	NotMatchTag          []*PEP425Tag `protobuf:"bytes,4,rep,name=not_match_tag,json=notMatchTag,proto3" json:"not_match_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Spec_Package) Reset()         { *m = Spec_Package{} }
func (m *Spec_Package) String() string { return proto.CompactTextString(m) }
func (*Spec_Package) ProtoMessage()    {}
func (*Spec_Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_12b41745b49e8c72, []int{0, 0}
}
func (m *Spec_Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spec_Package.Unmarshal(m, b)
}
func (m *Spec_Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spec_Package.Marshal(b, m, deterministic)
}
func (m *Spec_Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec_Package.Merge(m, src)
}
func (m *Spec_Package) XXX_Size() int {
	return xxx_messageInfo_Spec_Package.Size(m)
}
func (m *Spec_Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Spec_Package proto.InternalMessageInfo

func (m *Spec_Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Spec_Package) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Spec_Package) GetMatchTag() []*PEP425Tag {
	if m != nil {
		return m.MatchTag
	}
	return nil
}

func (m *Spec_Package) GetNotMatchTag() []*PEP425Tag {
	if m != nil {
		return m.NotMatchTag
	}
	return nil
}

func init() {
	proto.RegisterType((*Spec)(nil), "vpython.Spec")
	proto.RegisterType((*Spec_Package)(nil), "vpython.Spec.Package")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/vpython/api/vpython/spec.proto", fileDescriptor_12b41745b49e8c72)
}

var fileDescriptor_12b41745b49e8c72 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xe9, 0x9f, 0xef, 0xab, 0x73, 0x87, 0x51, 0x0c, 0x08, 0x65, 0x56, 0x45, 0x10, 0x06,
	0x84, 0x14, 0x3a, 0x53, 0x97, 0xee, 0x5c, 0x0a, 0xa5, 0x16, 0xb7, 0x25, 0x86, 0x98, 0x16, 0xdb,
	0x24, 0x64, 0xd2, 0xca, 0xbc, 0x8d, 0x0f, 0xe9, 0x03, 0x88, 0x49, 0xc7, 0x71, 0x53, 0x70, 0x77,
	0x7b, 0x7a, 0x7e, 0xe7, 0xdc, 0x1b, 0xd8, 0x72, 0x89, 0x69, 0xa3, 0x65, 0xdf, 0x0e, 0x3d, 0x96,
	0x9a, 0xa7, 0xdd, 0x40, 0xdb, 0x74, 0x54, 0x07, 0xd3, 0x48, 0x91, 0x12, 0x75, 0x9a, 0xf7, 0x8a,
	0x51, 0xac, 0xb4, 0x34, 0x12, 0x45, 0x93, 0xb6, 0xce, 0xff, 0x4c, 0x2b, 0xa6, 0x76, 0x59, 0xee,
	0xf8, 0xeb, 0x4f, 0x1f, 0xc2, 0x27, 0xc5, 0x28, 0xba, 0x81, 0x73, 0xf7, 0xbf, 0x1e, 0x99, 0xde,
	0xb7, 0x52, 0xc4, 0x5e, 0xe2, 0x6d, 0x16, 0xe5, 0xca, 0xa9, 0xcf, 0x4e, 0x44, 0xb7, 0xf0, 0xef,
	0xbd, 0x61, 0xac, 0x8b, 0xfd, 0x24, 0xd8, 0x2c, 0xb3, 0x2b, 0x3c, 0xa5, 0xe2, 0xef, 0x10, 0x5c,
	0x10, 0xfa, 0x46, 0x38, 0x2b, 0x9d, 0x07, 0xe5, 0x00, 0x63, 0xab, 0xcd, 0x40, 0x3a, 0x26, 0xc6,
	0x38, 0x48, 0xbc, 0x79, 0xe2, 0x97, 0x11, 0xdd, 0xc3, 0xe5, 0xc8, 0x74, 0xfb, 0x7a, 0xa8, 0xdd,
	0xaa, 0xb5, 0x21, 0x3c, 0x0e, 0x6d, 0x1f, 0xfa, 0xa1, 0x8b, 0x87, 0x62, 0x97, 0xe5, 0x15, 0xe1,
	0xe5, 0x85, 0x33, 0x17, 0xd6, 0x5b, 0x11, 0xbe, 0xfe, 0xf0, 0x20, 0x9a, 0x72, 0x11, 0x82, 0x50,
	0x90, 0x9e, 0x4d, 0xc7, 0xd8, 0x19, 0xc5, 0x10, 0x1d, 0x6f, 0xf4, 0xad, 0x7c, 0xfc, 0x44, 0x29,
	0x2c, 0x7a, 0x62, 0x68, 0x63, 0x1b, 0x83, 0xd9, 0xc6, 0x33, 0x6b, 0xaa, 0x08, 0x47, 0x77, 0xb0,
	0x12, 0xd2, 0xd4, 0x27, 0x68, 0x7e, 0xcd, 0xa5, 0x90, 0xe6, 0x71, 0xe2, 0x5e, 0xfe, 0xdb, 0xd7,
	0xdf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x80, 0x86, 0x0f, 0xf4, 0x01, 0x00, 0x00,
}
