// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/proto/srcman/manifest_diff.proto

package srcman

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	git "go.chromium.org/luci/common/proto/git"
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

// Stat indicates how a given item has changed.
type ManifestDiff_Stat int32

const (
	// These two items are identical
	ManifestDiff_EQUAL ManifestDiff_Stat = 0
	// The item was added in `new` compared to `old`
	ManifestDiff_ADDED ManifestDiff_Stat = 1
	// The item was removed in `new` compared to `old`
	ManifestDiff_REMOVED ManifestDiff_Stat = 2
	// The item is in both, but is incomparable (e.g. repo_url changed from
	// `old` to `new`).
	ManifestDiff_MODIFIED ManifestDiff_Stat = 4
	// The item is in both, and is directly comparable (e.g. different
	// revisions of the same repo_url). This only applies to the revision fields
	// of SCM messages.
	//
	// This is 0x8 | MODIFIED, so that users who don't care about DIFF v.
	// MODIFIED can check `Status & MODIFIED`.
	ManifestDiff_DIFF ManifestDiff_Stat = 12
)

var ManifestDiff_Stat_name = map[int32]string{
	0:  "EQUAL",
	1:  "ADDED",
	2:  "REMOVED",
	4:  "MODIFIED",
	12: "DIFF",
}

var ManifestDiff_Stat_value = map[string]int32{
	"EQUAL":    0,
	"ADDED":    1,
	"REMOVED":  2,
	"MODIFIED": 4,
	"DIFF":     12,
}

func (x ManifestDiff_Stat) String() string {
	return proto.EnumName(ManifestDiff_Stat_name, int32(x))
}

func (ManifestDiff_Stat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36818c0267d64e2d, []int{0, 0}
}

// ManifestDiff holds basic difference information between two source manifests.
type ManifestDiff struct {
	// The older of the two manifests.
	Old *Manifest `protobuf:"bytes,1,opt,name=old,proto3" json:"old,omitempty"`
	// The newer of the two manifests.
	New *Manifest `protobuf:"bytes,2,opt,name=new,proto3" json:"new,omitempty"`
	// Indicates if there is some overall difference between old and new.
	Overall              ManifestDiff_Stat                  `protobuf:"varint,3,opt,name=overall,proto3,enum=srcman.ManifestDiff_Stat" json:"overall,omitempty"`
	Directories          map[string]*ManifestDiff_Directory `protobuf:"bytes,4,rep,name=directories,proto3" json:"directories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ManifestDiff) Reset()         { *m = ManifestDiff{} }
func (m *ManifestDiff) String() string { return proto.CompactTextString(m) }
func (*ManifestDiff) ProtoMessage()    {}
func (*ManifestDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_36818c0267d64e2d, []int{0}
}
func (m *ManifestDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestDiff.Unmarshal(m, b)
}
func (m *ManifestDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestDiff.Marshal(b, m, deterministic)
}
func (m *ManifestDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestDiff.Merge(m, src)
}
func (m *ManifestDiff) XXX_Size() int {
	return xxx_messageInfo_ManifestDiff.Size(m)
}
func (m *ManifestDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestDiff.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestDiff proto.InternalMessageInfo

func (m *ManifestDiff) GetOld() *Manifest {
	if m != nil {
		return m.Old
	}
	return nil
}

func (m *ManifestDiff) GetNew() *Manifest {
	if m != nil {
		return m.New
	}
	return nil
}

func (m *ManifestDiff) GetOverall() ManifestDiff_Stat {
	if m != nil {
		return m.Overall
	}
	return ManifestDiff_EQUAL
}

func (m *ManifestDiff) GetDirectories() map[string]*ManifestDiff_Directory {
	if m != nil {
		return m.Directories
	}
	return nil
}

type ManifestDiff_GitCheckout struct {
	// Indicates if there is some overall difference between old and new.
	Overall ManifestDiff_Stat `protobuf:"varint,1,opt,name=overall,proto3,enum=srcman.ManifestDiff_Stat" json:"overall,omitempty"`
	// Indicates the status for the `revision` field.
	//
	// If this is DIFF, it is sensible to compute
	//   `git log repo_url old.revision new.revision`
	Revision ManifestDiff_Stat `protobuf:"varint,2,opt,name=revision,proto3,enum=srcman.ManifestDiff_Stat" json:"revision,omitempty"`
	// Indicates the status for the `patch_revision` field. It evaluates
	// the patch_fetch_ref values to ensure that old and new are different
	// patches from the same CL.
	//
	// If this is DIFF, it is sensible to compute
	//   `git log repo_url old.patch_revision new.patch_revision`
	PatchRevision ManifestDiff_Stat `protobuf:"varint,3,opt,name=patch_revision,json=patchRevision,proto3,enum=srcman.ManifestDiff_Stat" json:"patch_revision,omitempty"`
	// The URL that should be used for RPCs. It may differ from the url in old
	// or new if the service computing this ManifestDiff knows of e.g. a repo
	// URL migration.
	RepoUrl string `protobuf:"bytes,4,opt,name=repo_url,json=repoUrl,proto3" json:"repo_url,omitempty"`
	// If revision==DIFF, this may be populated with git history occuring
	// between the two base revisions.
	History              []*git.Commit `protobuf:"bytes,5,rep,name=history,proto3" json:"history,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ManifestDiff_GitCheckout) Reset()         { *m = ManifestDiff_GitCheckout{} }
func (m *ManifestDiff_GitCheckout) String() string { return proto.CompactTextString(m) }
func (*ManifestDiff_GitCheckout) ProtoMessage()    {}
func (*ManifestDiff_GitCheckout) Descriptor() ([]byte, []int) {
	return fileDescriptor_36818c0267d64e2d, []int{0, 0}
}
func (m *ManifestDiff_GitCheckout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestDiff_GitCheckout.Unmarshal(m, b)
}
func (m *ManifestDiff_GitCheckout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestDiff_GitCheckout.Marshal(b, m, deterministic)
}
func (m *ManifestDiff_GitCheckout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestDiff_GitCheckout.Merge(m, src)
}
func (m *ManifestDiff_GitCheckout) XXX_Size() int {
	return xxx_messageInfo_ManifestDiff_GitCheckout.Size(m)
}
func (m *ManifestDiff_GitCheckout) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestDiff_GitCheckout.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestDiff_GitCheckout proto.InternalMessageInfo

func (m *ManifestDiff_GitCheckout) GetOverall() ManifestDiff_Stat {
	if m != nil {
		return m.Overall
	}
	return ManifestDiff_EQUAL
}

func (m *ManifestDiff_GitCheckout) GetRevision() ManifestDiff_Stat {
	if m != nil {
		return m.Revision
	}
	return ManifestDiff_EQUAL
}

func (m *ManifestDiff_GitCheckout) GetPatchRevision() ManifestDiff_Stat {
	if m != nil {
		return m.PatchRevision
	}
	return ManifestDiff_EQUAL
}

func (m *ManifestDiff_GitCheckout) GetRepoUrl() string {
	if m != nil {
		return m.RepoUrl
	}
	return ""
}

func (m *ManifestDiff_GitCheckout) GetHistory() []*git.Commit {
	if m != nil {
		return m.History
	}
	return nil
}

type ManifestDiff_Directory struct {
	// This is the overall status for this Directory.
	Overall        ManifestDiff_Stat         `protobuf:"varint,1,opt,name=overall,proto3,enum=srcman.ManifestDiff_Stat" json:"overall,omitempty"`
	GitCheckout    *ManifestDiff_GitCheckout `protobuf:"bytes,2,opt,name=git_checkout,json=gitCheckout,proto3" json:"git_checkout,omitempty"`
	CipdServerHost ManifestDiff_Stat         `protobuf:"varint,3,opt,name=cipd_server_host,json=cipdServerHost,proto3,enum=srcman.ManifestDiff_Stat" json:"cipd_server_host,omitempty"`
	// Note: this will only ever be MODIFIED, because we cannot (yet) determine
	// if two versions of a cipd package are diffable. We may later implement
	// DIFF detection (i.e. if both packages use `version:<sha1>` tags).
	CipdPackage        map[string]ManifestDiff_Stat `protobuf:"bytes,4,rep,name=cipd_package,json=cipdPackage,proto3" json:"cipd_package,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=srcman.ManifestDiff_Stat"`
	IsolatedServerHost ManifestDiff_Stat            `protobuf:"varint,5,opt,name=isolated_server_host,json=isolatedServerHost,proto3,enum=srcman.ManifestDiff_Stat" json:"isolated_server_host,omitempty"`
	// This merely indicates if the list of isolated hashes was the same or not;
	// there's not a good way to register the two lists.
	//
	// Since order-of-application for isolateds matters, this will indicate
	// MODIFIED if the order of isolated hashes changes.
	Isolated             ManifestDiff_Stat `protobuf:"varint,6,opt,name=isolated,proto3,enum=srcman.ManifestDiff_Stat" json:"isolated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ManifestDiff_Directory) Reset()         { *m = ManifestDiff_Directory{} }
func (m *ManifestDiff_Directory) String() string { return proto.CompactTextString(m) }
func (*ManifestDiff_Directory) ProtoMessage()    {}
func (*ManifestDiff_Directory) Descriptor() ([]byte, []int) {
	return fileDescriptor_36818c0267d64e2d, []int{0, 1}
}
func (m *ManifestDiff_Directory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestDiff_Directory.Unmarshal(m, b)
}
func (m *ManifestDiff_Directory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestDiff_Directory.Marshal(b, m, deterministic)
}
func (m *ManifestDiff_Directory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestDiff_Directory.Merge(m, src)
}
func (m *ManifestDiff_Directory) XXX_Size() int {
	return xxx_messageInfo_ManifestDiff_Directory.Size(m)
}
func (m *ManifestDiff_Directory) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestDiff_Directory.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestDiff_Directory proto.InternalMessageInfo

func (m *ManifestDiff_Directory) GetOverall() ManifestDiff_Stat {
	if m != nil {
		return m.Overall
	}
	return ManifestDiff_EQUAL
}

func (m *ManifestDiff_Directory) GetGitCheckout() *ManifestDiff_GitCheckout {
	if m != nil {
		return m.GitCheckout
	}
	return nil
}

func (m *ManifestDiff_Directory) GetCipdServerHost() ManifestDiff_Stat {
	if m != nil {
		return m.CipdServerHost
	}
	return ManifestDiff_EQUAL
}

func (m *ManifestDiff_Directory) GetCipdPackage() map[string]ManifestDiff_Stat {
	if m != nil {
		return m.CipdPackage
	}
	return nil
}

func (m *ManifestDiff_Directory) GetIsolatedServerHost() ManifestDiff_Stat {
	if m != nil {
		return m.IsolatedServerHost
	}
	return ManifestDiff_EQUAL
}

func (m *ManifestDiff_Directory) GetIsolated() ManifestDiff_Stat {
	if m != nil {
		return m.Isolated
	}
	return ManifestDiff_EQUAL
}

func init() {
	proto.RegisterType((*ManifestDiff)(nil), "srcman.ManifestDiff")
	proto.RegisterMapType((map[string]*ManifestDiff_Directory)(nil), "srcman.ManifestDiff.DirectoriesEntry")
	proto.RegisterType((*ManifestDiff_GitCheckout)(nil), "srcman.ManifestDiff.GitCheckout")
	proto.RegisterType((*ManifestDiff_Directory)(nil), "srcman.ManifestDiff.Directory")
	proto.RegisterMapType((map[string]ManifestDiff_Stat)(nil), "srcman.ManifestDiff.Directory.CipdPackageEntry")
	proto.RegisterEnum("srcman.ManifestDiff_Stat", ManifestDiff_Stat_name, ManifestDiff_Stat_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/proto/srcman/manifest_diff.proto", fileDescriptor_36818c0267d64e2d)
}

var fileDescriptor_36818c0267d64e2d = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0xcd, 0xf6, 0xf7, 0x4b, 0x2d, 0x61, 0xf0, 0x90, 0xcd, 0x41, 0x4a, 0x61, 0xa1, 0xa7,
	0x04, 0xba, 0x2e, 0x8a, 0x78, 0xb0, 0x34, 0xe9, 0x5a, 0xb4, 0xac, 0x66, 0x59, 0xc1, 0x8b, 0x21,
	0x4e, 0xa7, 0xe9, 0xd0, 0x24, 0x13, 0x26, 0xd3, 0x2e, 0xfd, 0xcb, 0x3d, 0x2d, 0x48, 0x26, 0xcd,
	0x6e, 0x2c, 0xb5, 0x51, 0x6f, 0xd3, 0xf7, 0xbe, 0x9f, 0xf7, 0xde, 0x7c, 0xe7, 0x35, 0xf0, 0x2e,
	0x60, 0x26, 0x5e, 0x71, 0x16, 0xd1, 0x4d, 0x64, 0x32, 0x1e, 0x58, 0xe1, 0x06, 0x53, 0x0b, 0xb3,
	0x28, 0x62, 0xb1, 0x95, 0x70, 0x26, 0x98, 0x95, 0x72, 0x1c, 0xf9, 0xb1, 0x15, 0xf9, 0x31, 0x5d,
	0x92, 0x54, 0x78, 0x0b, 0xba, 0x5c, 0x9a, 0x32, 0x85, 0x9a, 0x79, 0xce, 0x18, 0x55, 0x57, 0x09,
	0xa8, 0x90, 0x01, 0x2a, 0x72, 0xd6, 0x78, 0xfd, 0xcf, 0x9d, 0x73, 0x70, 0xf0, 0xb3, 0x0d, 0xdd,
	0xf9, 0x3e, 0x64, 0xd3, 0xe5, 0x12, 0x0d, 0xa0, 0xc6, 0xc2, 0x85, 0xae, 0xf4, 0x95, 0xa1, 0x3a,
	0xd2, 0xcc, 0x9c, 0x32, 0x0b, 0x89, 0x9b, 0x25, 0x33, 0x4d, 0x4c, 0xee, 0xf5, 0xb3, 0x3f, 0x69,
	0x62, 0x72, 0x8f, 0x2e, 0xa1, 0xc5, 0xb6, 0x84, 0xfb, 0x61, 0xa8, 0xd7, 0xfa, 0xca, 0xb0, 0x37,
	0x3a, 0x3f, 0xd4, 0x65, 0xed, 0xcc, 0x5b, 0xe1, 0x0b, 0xb7, 0x50, 0xa2, 0x6b, 0x50, 0x17, 0x94,
	0x13, 0x2c, 0x18, 0xa7, 0x24, 0xd5, 0xeb, 0xfd, 0xda, 0x50, 0x1d, 0x5d, 0x1c, 0x05, 0xed, 0x27,
	0x9d, 0x13, 0x0b, 0xbe, 0x73, 0xcb, 0xa4, 0xf1, 0xa0, 0x80, 0x7a, 0x4d, 0xc5, 0x64, 0x45, 0xf0,
	0x9a, 0x6d, 0x44, 0x79, 0x1a, 0xe5, 0xaf, 0xa7, 0xb9, 0x82, 0x36, 0x27, 0x5b, 0x9a, 0x52, 0x16,
	0xcb, 0xbb, 0x9e, 0xa4, 0x1e, 0xa5, 0xe8, 0x3d, 0xf4, 0x12, 0x5f, 0xe0, 0x95, 0xf7, 0x08, 0x57,
	0x1a, 0xf0, 0x5c, 0x02, 0x6e, 0x51, 0xe1, 0x3c, 0x6b, 0x9c, 0x30, 0x6f, 0xc3, 0x43, 0xbd, 0xde,
	0x57, 0x86, 0x1d, 0xb7, 0x95, 0xfd, 0xbe, 0xe3, 0x21, 0xba, 0x80, 0xd6, 0x8a, 0xa6, 0x82, 0xf1,
	0x9d, 0xde, 0x90, 0xee, 0xa8, 0x66, 0x40, 0x85, 0x39, 0x91, 0xcb, 0xe0, 0x16, 0x39, 0xe3, 0xa1,
	0x06, 0x9d, 0xc2, 0xa1, 0xdd, 0xff, 0xdd, 0x7e, 0x02, 0xdd, 0x80, 0x0a, 0x0f, 0xef, 0x2d, 0xdc,
	0xbf, 0x76, 0xff, 0x28, 0x59, 0xb2, 0xda, 0x55, 0x83, 0x92, 0xef, 0x13, 0xd0, 0x30, 0x4d, 0x16,
	0x5e, 0x4a, 0xf8, 0x96, 0x70, 0x6f, 0xc5, 0x52, 0x51, 0xed, 0x46, 0x2f, 0x43, 0x6e, 0x25, 0xf1,
	0x81, 0xa5, 0x02, 0xb9, 0xd0, 0x95, 0x45, 0x12, 0x1f, 0xaf, 0xfd, 0x80, 0xec, 0xd7, 0xc2, 0x3a,
	0xb9, 0x16, 0x3b, 0x73, 0x42, 0x93, 0xc5, 0xe7, 0x9c, 0xd8, 0x2f, 0x08, 0x7e, 0x8a, 0xa0, 0x8f,
	0xf0, 0x82, 0xa6, 0x2c, 0xf4, 0x05, 0xf9, 0x7d, 0xb8, 0x46, 0xd5, 0x70, 0xa8, 0xc0, 0x4a, 0x03,
	0x5e, 0x41, 0xbb, 0x88, 0xea, 0xcd, 0xca, 0x45, 0x29, 0xa4, 0xc6, 0x37, 0xd0, 0x0e, 0x87, 0x44,
	0x1a, 0xd4, 0xd6, 0x64, 0x27, 0x9f, 0xa9, 0xe3, 0x66, 0x47, 0x64, 0x41, 0x63, 0xeb, 0x87, 0x1b,
	0x52, 0xbd, 0x82, 0xb9, 0xee, 0xed, 0xd9, 0x1b, 0xc5, 0xf8, 0x0e, 0xda, 0xe1, 0x1f, 0xe4, 0x48,
	0xe9, 0x57, 0xe5, 0xd2, 0xea, 0xe8, 0xe5, 0x69, 0x47, 0x4b, 0xf5, 0x07, 0x63, 0xa8, 0x67, 0x2d,
	0x51, 0x07, 0x1a, 0xce, 0x97, 0xbb, 0xf1, 0x27, 0xed, 0x59, 0x76, 0x1c, 0xdb, 0xb6, 0x63, 0x6b,
	0x0a, 0x52, 0xa1, 0xe5, 0x3a, 0xf3, 0x9b, 0xaf, 0x8e, 0xad, 0x9d, 0xa1, 0x2e, 0xb4, 0xe7, 0x37,
	0xf6, 0x6c, 0x3a, 0x73, 0x6c, 0xad, 0x8e, 0xda, 0x50, 0xb7, 0x67, 0xd3, 0xa9, 0xd6, 0xfd, 0xd1,
	0x94, 0x1f, 0xa0, 0xcb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0xa5, 0x63, 0xa7, 0x35, 0x05,
	0x00, 0x00,
}
