// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: utils/gogo/graphql.gen.proto

package gogo

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//oneof和enum的区别就是一个可以多种类型一个只能是一种类型
type Operation int32

const (
	Operation_Query    Operation = 0
	Operation_Mutation Operation = 1
	Operation_Default  Operation = 2
)

var Operation_name = map[int32]string{
	0: "Query",
	1: "Mutation",
	2: "Default",
}

var Operation_value = map[string]int32{
	"Query":    0,
	"Mutation": 1,
	"Default":  2,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}

func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fc7562cb32fcf10, []int{0}
}

type Field struct {
	Required             bool     `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty"`
	Params               string   `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Dirs                 string   `protobuf:"bytes,3,opt,name=dirs,proto3" json:"dirs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fc7562cb32fcf10, []int{0}
}
func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *Field) GetParams() string {
	if m != nil {
		return m.Params
	}
	return ""
}

func (m *Field) GetDirs() string {
	if m != nil {
		return m.Dirs
	}
	return ""
}

var E_GraphqlOperation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Operation)(nil),
	Field:         1001,
	Name:          "gogo.graphql_operation",
	Tag:           "varint,1001,opt,name=graphql_operation,enum=gogo.Operation",
	Filename:      "utils/gogo/graphql.gen.proto",
}

var E_Operation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*Operation)(nil),
	Field:         1001,
	Name:          "gogo.operation",
	Tag:           "varint,1001,opt,name=operation,enum=gogo.Operation",
	Filename:      "utils/gogo/graphql.gen.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Field)(nil),
	Field:         1001,
	Name:          "gogo.field",
	Tag:           "bytes,1001,opt,name=field",
	Filename:      "utils/gogo/graphql.gen.proto",
}

func init() {
	proto.RegisterEnum("gogo.Operation", Operation_name, Operation_value)
	proto.RegisterType((*Field)(nil), "gogo.Field")
	proto.RegisterExtension(E_GraphqlOperation)
	proto.RegisterExtension(E_Operation)
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("utils/gogo/graphql.gen.proto", fileDescriptor_2fc7562cb32fcf10) }

var fileDescriptor_2fc7562cb32fcf10 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x39, 0x4f, 0xf3, 0x40,
	0x10, 0x86, 0x3f, 0xe7, 0xcb, 0xe5, 0x09, 0x87, 0xd9, 0x02, 0x59, 0x11, 0x87, 0x05, 0x4d, 0x44,
	0xb1, 0x16, 0xa1, 0x4b, 0x47, 0x84, 0xa0, 0x8a, 0x0c, 0xa6, 0x46, 0xc8, 0x89, 0x27, 0x9b, 0x95,
	0x4c, 0x76, 0xb3, 0x07, 0x22, 0xfc, 0x63, 0xfe, 0x05, 0xf2, 0x3a, 0x38, 0x42, 0x29, 0xe8, 0x3c,
	0xcf, 0x78, 0x9e, 0x7d, 0x67, 0xe0, 0xc4, 0x1a, 0x5e, 0xe8, 0x98, 0x09, 0x26, 0x62, 0xa6, 0x32,
	0xb9, 0x58, 0x15, 0x94, 0xe1, 0x92, 0x4a, 0x25, 0x8c, 0x20, 0xcd, 0x92, 0xf7, 0x23, 0x26, 0x04,
	0x2b, 0x30, 0x76, 0x6c, 0x6a, 0xe7, 0x71, 0x8e, 0x7a, 0xa6, 0xb8, 0x34, 0x42, 0x55, 0xff, 0x5d,
	0x24, 0xd0, 0xba, 0xe7, 0x58, 0xe4, 0xa4, 0x0f, 0x5d, 0x85, 0x2b, 0xcb, 0x15, 0xe6, 0xa1, 0x17,
	0x79, 0x83, 0x6e, 0x5a, 0xd7, 0xe4, 0x18, 0xda, 0x32, 0x53, 0xd9, 0x9b, 0x0e, 0x1b, 0x91, 0x37,
	0xf0, 0xd3, 0x4d, 0x45, 0x08, 0x34, 0x73, 0xae, 0x74, 0xf8, 0xdf, 0x51, 0xf7, 0x7d, 0x75, 0x0d,
	0x7e, 0x22, 0x51, 0x65, 0x86, 0x8b, 0x25, 0xf1, 0xa1, 0xf5, 0x64, 0x51, 0xad, 0x83, 0x7f, 0x64,
	0x0f, 0xba, 0x13, 0x6b, 0x1c, 0x0e, 0x3c, 0xd2, 0x83, 0xce, 0x1d, 0xce, 0x33, 0x5b, 0x98, 0xa0,
	0x31, 0x7a, 0x81, 0xa3, 0xcd, 0x02, 0xaf, 0xa2, 0x1e, 0x3d, 0xa3, 0x55, 0x76, 0xfa, 0x93, 0x9d,
	0x4e, 0xd0, 0x2c, 0x44, 0x9e, 0xc8, 0xb2, 0xad, 0xc3, 0xaf, 0x4e, 0xe4, 0x0d, 0x0e, 0x86, 0x87,
	0xb4, 0x5c, 0x94, 0xd6, 0x4f, 0xa6, 0xc1, 0x46, 0x55, 0x93, 0x51, 0x02, 0xfe, 0x56, 0x7b, 0xbe,
	0xa3, 0x7d, 0x46, 0xf5, 0xce, 0x67, 0xf8, 0x97, 0x77, 0xeb, 0x18, 0xdd, 0x42, 0x6b, 0xee, 0x6e,
	0x76, 0xba, 0x23, 0x73, 0xb7, 0xfc, 0xa5, 0xea, 0x0d, 0x7b, 0x95, 0xca, 0xb5, 0xd2, 0x6a, 0x72,
	0x7c, 0x09, 0xe4, 0x63, 0xfd, 0x49, 0x17, 0xa5, 0xb4, 0x9e, 0x1e, 0xef, 0x3f, 0x54, 0xd9, 0x1f,
	0x4b, 0xa0, 0xa7, 0x6d, 0xd7, 0xb8, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xac, 0x53, 0xab, 0xa1,
	0xea, 0x01, 0x00, 0x00,
}