// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package testprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	prototype "google.golang.org/protobuf/reflect/prototype"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

type Enum1 int32

const (
	Enum1_ONE Enum1 = 1
)

var Enum1_name = map[int32]string{
	1: "ONE",
}

var Enum1_value = map[string]int32{
	"ONE": 1,
}

func (x Enum1) Enum() *Enum1 {
	p := new(Enum1)
	*p = x
	return p
}

func (x Enum1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum1) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[0].EnumDescriptor
}

func (Enum1) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[0]
}

func (x Enum1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Enum1) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Enum1(num)
	return nil
}

// Deprecated: Use Enum1.Descriptor instead.
func (Enum1) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

type Enum2 int32

const (
	Enum2_UNO Enum2 = 1
)

var Enum2_name = map[int32]string{
	1: "UNO",
}

var Enum2_value = map[string]int32{
	"UNO": 1,
}

func (x Enum2) Enum() *Enum2 {
	p := new(Enum2)
	*p = x
	return p
}

func (x Enum2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum2) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[1].EnumDescriptor
}

func (Enum2) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[1]
}

func (x Enum2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Enum2) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Enum2(num)
	return nil
}

// Deprecated: Use Enum2.Descriptor instead.
func (Enum2) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

type Enum3 int32

const (
	Enum3_YI Enum3 = 1
)

var Enum3_name = map[int32]string{
	1: "YI",
}

var Enum3_value = map[string]int32{
	"YI": 1,
}

func (x Enum3) Enum() *Enum3 {
	p := new(Enum3)
	*p = x
	return p
}

func (x Enum3) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum3) Descriptor() protoreflect.EnumDescriptor {
	return file_test_proto_enumTypes[2].EnumDescriptor
}

func (Enum3) Type() protoreflect.EnumType {
	return &file_test_proto_enumTypes[2]
}

func (x Enum3) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Enum3) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Enum3(num)
	return nil
}

// Deprecated: Use Enum3.Descriptor instead.
func (Enum3) EnumDescriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

type Message1 struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields
}

func (x *Message1) Reset() {
	*x = Message1{}
}

func (x *Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message1) ProtoMessage() {}

func (x *Message1) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message1.ProtoReflect.Descriptor instead.
func (*Message1) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

var extRange_Message1 = []protoiface.ExtensionRangeV1{
	{Start: 10, End: 536870911},
}

// Deprecated: Use Message1.ProtoReflect.Descriptor.ExtensionRanges instead.
func (*Message1) ExtensionRangeArray() []protoiface.ExtensionRangeV1 {
	return extRange_Message1
}

type Message2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message2) Reset() {
	*x = Message2{}
}

func (x *Message2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message2) ProtoMessage() {}

func (x *Message2) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message2.ProtoReflect.Descriptor instead.
func (*Message2) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

type Message3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message3) Reset() {
	*x = Message3{}
}

func (x *Message3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message3) ProtoMessage() {}

func (x *Message3) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message3.ProtoReflect.Descriptor instead.
func (*Message3) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

type Message4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	BoolField     *bool `protobuf:"varint,30,opt,name=bool_field,json=boolField" json:"bool_field,omitempty"`
}

func (x *Message4) Reset() {
	*x = Message4{}
}

func (x *Message4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message4) ProtoMessage() {}

func (x *Message4) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message4.ProtoReflect.Descriptor instead.
func (*Message4) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *Message4) GetBoolField() bool {
	if x != nil && x.BoolField != nil {
		return *x.BoolField
	}
	return false
}

var file_test_proto_extDescs = []protoiface.ExtensionDescV1{
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*string)(nil),
		Field:         11,
		Name:          "testprotos.string_field",
		Tag:           "bytes,11,opt,name=string_field",
		Filename:      "test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*Enum1)(nil),
		Field:         12,
		Name:          "testprotos.enum_field",
		Tag:           "varint,12,opt,name=enum_field,enum=testprotos.Enum1",
		Filename:      "test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*Message2)(nil),
		Field:         13,
		Name:          "testprotos.message_field",
		Tag:           "bytes,13,opt,name=message_field",
		Filename:      "test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*Message2)(nil),
		Field:         21,
		Name:          "testprotos.Message4.message_field",
		Tag:           "bytes,21,opt,name=message_field",
		Filename:      "test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*Enum1)(nil),
		Field:         22,
		Name:          "testprotos.Message4.enum_field",
		Tag:           "varint,22,opt,name=enum_field,enum=testprotos.Enum1",
		Filename:      "test.proto",
	},
	{
		ExtendedType:  (*Message1)(nil),
		ExtensionType: (*string)(nil),
		Field:         23,
		Name:          "testprotos.Message4.string_field",
		Tag:           "bytes,23,opt,name=string_field",
		Filename:      "test.proto",
	},
}
var (
	// extend testprotos.Message1 { optional string string_field = 11; }
	E_StringField = &file_test_proto_extDescs[0]

	// extend testprotos.Message1 { optional testprotos.Enum1 enum_field = 12; }
	E_EnumField = &file_test_proto_extDescs[1]

	// extend testprotos.Message1 { optional testprotos.Message2 message_field = 13; }
	E_MessageField = &file_test_proto_extDescs[2]

	// extend testprotos.Message1 { optional testprotos.Message2 message_field = 21; }
	E_Message4_MessageField = &file_test_proto_extDescs[3]

	// extend testprotos.Message1 { optional testprotos.Enum1 enum_field = 22; }
	E_Message4_EnumField = &file_test_proto_extDescs[4]

	// extend testprotos.Message1 { optional string string_field = 23; }
	E_Message4_StringField = &file_test_proto_extDescs[5]
)
var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x14, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x2a, 0x08, 0x08, 0x0a, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x22, 0x0a,
	0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x22, 0x0a, 0x0a, 0x08, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x22, 0xfb, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x34, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x4f, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x46, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31,
	0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x0a, 0x0c, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x31, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x2a, 0x10, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x07, 0x0a,
	0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x2a, 0x10, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x32, 0x12,
	0x07, 0x0a, 0x03, 0x55, 0x4e, 0x4f, 0x10, 0x01, 0x2a, 0x0f, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d,
	0x33, 0x12, 0x06, 0x0a, 0x02, 0x59, 0x49, 0x10, 0x01, 0x3a, 0x37, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x3a, 0x46, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x31, 0x52,
	0x09, 0x65, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x4f, 0x0a, 0x0d, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x31, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x52, 0x0c, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_enumTypes = make([]prototype.Enum, 3)
var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_test_proto_goTypes = []interface{}{
	(Enum1)(0),       // 0: testprotos.Enum1
	(Enum2)(0),       // 1: testprotos.Enum2
	(Enum3)(0),       // 2: testprotos.Enum3
	(*Message1)(nil), // 3: testprotos.Message1
	(*Message2)(nil), // 4: testprotos.Message2
	(*Message3)(nil), // 5: testprotos.Message3
	(*Message4)(nil), // 6: testprotos.Message4
}
var file_test_proto_depIdxs = []int32{
	3,  // testprotos.string_field:extendee -> testprotos.Message1
	3,  // testprotos.enum_field:extendee -> testprotos.Message1
	3,  // testprotos.message_field:extendee -> testprotos.Message1
	3,  // testprotos.Message4.message_field:extendee -> testprotos.Message1
	3,  // testprotos.Message4.enum_field:extendee -> testprotos.Message1
	3,  // testprotos.Message4.string_field:extendee -> testprotos.Message1
	0,  // testprotos.enum_field:type_name -> testprotos.Enum1
	4,  // testprotos.message_field:type_name -> testprotos.Message2
	4,  // testprotos.Message4.message_field:type_name -> testprotos.Message2
	0,  // testprotos.Message4.enum_field:type_name -> testprotos.Enum1
	10, // starting offset of method output_type sub-list
	10, // starting offset of method input_type sub-list
	6,  // starting offset of extension type_name sub-list
	0,  // starting offset of extension extendee sub-list
	0,  // starting offset of field type_name sub-list
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message3); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message4); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
		LegacyExtensions:  file_test_proto_extDescs,
	}.Build()
	File_test_proto = out.File
	file_test_proto_enumTypes = out.Enums
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
