// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// In most cases, detect point should be `server` or `client`.
// Even in service mesh, this means `server`/`client` side sidecar
// `proxy` is reserved only.
type DetectPoint int32

const (
	DetectPoint_client DetectPoint = 0
	DetectPoint_server DetectPoint = 1
	DetectPoint_proxy  DetectPoint = 2
)

var DetectPoint_name = map[int32]string{
	0: "client",
	1: "server",
	2: "proxy",
}

var DetectPoint_value = map[string]int32{
	"client": 0,
	"server": 1,
	"proxy":  2,
}

func (x DetectPoint) String() string {
	return proto.EnumName(DetectPoint_name, int32(x))
}

func (DetectPoint) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type KeyStringValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyStringValuePair) Reset()         { *m = KeyStringValuePair{} }
func (m *KeyStringValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyStringValuePair) ProtoMessage()    {}
func (*KeyStringValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *KeyStringValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyStringValuePair.Unmarshal(m, b)
}
func (m *KeyStringValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyStringValuePair.Marshal(b, m, deterministic)
}
func (m *KeyStringValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyStringValuePair.Merge(m, src)
}
func (m *KeyStringValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyStringValuePair.Size(m)
}
func (m *KeyStringValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyStringValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyStringValuePair proto.InternalMessageInfo

func (m *KeyStringValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyStringValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KeyIntValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyIntValuePair) Reset()         { *m = KeyIntValuePair{} }
func (m *KeyIntValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyIntValuePair) ProtoMessage()    {}
func (*KeyIntValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *KeyIntValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyIntValuePair.Unmarshal(m, b)
}
func (m *KeyIntValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyIntValuePair.Marshal(b, m, deterministic)
}
func (m *KeyIntValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyIntValuePair.Merge(m, src)
}
func (m *KeyIntValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyIntValuePair.Size(m)
}
func (m *KeyIntValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyIntValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyIntValuePair proto.InternalMessageInfo

func (m *KeyIntValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyIntValuePair) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Commands struct {
	Commands             []*Command `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Commands) Reset()         { *m = Commands{} }
func (m *Commands) String() string { return proto.CompactTextString(m) }
func (*Commands) ProtoMessage()    {}
func (*Commands) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *Commands) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commands.Unmarshal(m, b)
}
func (m *Commands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commands.Marshal(b, m, deterministic)
}
func (m *Commands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commands.Merge(m, src)
}
func (m *Commands) XXX_Size() int {
	return xxx_messageInfo_Commands.Size(m)
}
func (m *Commands) XXX_DiscardUnknown() {
	xxx_messageInfo_Commands.DiscardUnknown(m)
}

var xxx_messageInfo_Commands proto.InternalMessageInfo

func (m *Commands) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

type Command struct {
	Command              string                `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Args                 []*KeyStringValuePair `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Command) GetArgs() []*KeyStringValuePair {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterEnum("DetectPoint", DetectPoint_name, DetectPoint_value)
	proto.RegisterType((*KeyStringValuePair)(nil), "KeyStringValuePair")
	proto.RegisterType((*KeyIntValuePair)(nil), "KeyIntValuePair")
	proto.RegisterType((*Commands)(nil), "Commands")
	proto.RegisterType((*Command)(nil), "Command")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x6d, 0xe7, 0xb6, 0xee, 0xed, 0x60, 0xc9, 0x44, 0x8a, 0xa7, 0x51, 0x04, 0x8b, 0x87,
	0x6c, 0x28, 0x08, 0x82, 0x27, 0xf5, 0x22, 0x13, 0x29, 0x1d, 0x28, 0x78, 0xcb, 0xe2, 0xa3, 0x96,
	0xb6, 0x49, 0x49, 0xe3, 0x66, 0xfe, 0x25, 0xff, 0x4a, 0x69, 0x93, 0x79, 0xf1, 0xe2, 0x25, 0xf9,
	0xde, 0xf7, 0xe5, 0x97, 0x90, 0xf7, 0x60, 0xc6, 0x65, 0x5d, 0x4b, 0xb1, 0xb0, 0x1b, 0x6d, 0x94,
	0xd4, 0x32, 0xbe, 0x05, 0xb2, 0x42, 0xb3, 0xd6, 0xaa, 0x10, 0xf9, 0x0b, 0xab, 0x3e, 0x31, 0x65,
	0x85, 0x22, 0x21, 0x0c, 0x4a, 0x34, 0x91, 0x37, 0xf7, 0x92, 0x49, 0xd6, 0x49, 0x72, 0x0c, 0xc3,
	0x6d, 0x17, 0x47, 0x7e, 0xef, 0xd9, 0x22, 0xbe, 0x81, 0xa3, 0x15, 0x9a, 0x47, 0xa1, 0xff, 0x8d,
	0x0e, 0xf7, 0xe8, 0x12, 0x82, 0x7b, 0x59, 0xd7, 0x4c, 0xbc, 0xb7, 0xe4, 0x0c, 0x02, 0xee, 0x74,
	0xe4, 0xcd, 0x07, 0xc9, 0xf4, 0x32, 0xa0, 0x2e, 0xcc, 0x7e, 0x93, 0xf8, 0x09, 0xc6, 0xce, 0x24,
	0x11, 0x8c, 0x9d, 0xed, 0x1e, 0xda, 0x97, 0xe4, 0x1c, 0x0e, 0x99, 0xca, 0xdb, 0xc8, 0xef, 0xaf,
	0x99, 0xd1, 0xbf, 0x9f, 0xcb, 0xfa, 0x03, 0x17, 0x4b, 0x98, 0x3e, 0xa0, 0x46, 0xae, 0x53, 0x59,
	0x08, 0x4d, 0x00, 0x46, 0xbc, 0x2a, 0x50, 0xe8, 0xf0, 0xa0, 0xd3, 0x2d, 0xaa, 0x2d, 0xaa, 0xd0,
	0x23, 0x13, 0x18, 0x36, 0x4a, 0x7e, 0x99, 0xd0, 0xbf, 0x63, 0x90, 0x48, 0x95, 0x53, 0xd6, 0x30,
	0xfe, 0x81, 0xb4, 0x2d, 0xcd, 0x8e, 0x55, 0x65, 0x21, 0x3a, 0xa7, 0xa6, 0x02, 0xf5, 0x4e, 0xaa,
	0x92, 0xda, 0xe6, 0xa6, 0xde, 0xdb, 0x09, 0xcb, 0x51, 0xe8, 0x85, 0x5d, 0x9b, 0xcd, 0xb5, 0x6b,
	0xfb, 0xb7, 0x7f, 0xba, 0x2e, 0xcd, 0xab, 0x23, 0x9f, 0x2d, 0x95, 0x76, 0xa3, 0xe0, 0xb2, 0xda,
	0x8c, 0xfa, 0xa1, 0x5c, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x61, 0xcd, 0x58, 0xab, 0x01,
	0x00, 0x00,
}
