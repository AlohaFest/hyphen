
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: machine/machine.proto

package machine

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Instruction struct {
	Operator             string   `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Operand              int32    `protobuf:"varint,2,opt,name=operand,proto3" json:"operand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instruction) Reset()         { *m = Instruction{} }
func (m *Instruction) String() string { return proto.CompactTextString(m) }
func (*Instruction) ProtoMessage()    {}
func (*Instruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_84b4f59d98cc997c, []int{0}
}

func (m *Instruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instruction.Unmarshal(m, b)
}
func (m *Instruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instruction.Marshal(b, m, deterministic)
}
func (m *Instruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instruction.Merge(m, src)
}
func (m *Instruction) XXX_Size() int {
	return xxx_messageInfo_Instruction.Size(m)
}
func (m *Instruction) XXX_DiscardUnknown() {
	xxx_messageInfo_Instruction.DiscardUnknown(m)
}

var xxx_messageInfo_Instruction proto.InternalMessageInfo

func (m *Instruction) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *Instruction) GetOperand() int32 {
	if m != nil {
		return m.Operand
	}
	return 0
}

type Result struct {
	Output               float32  `protobuf:"fixed32,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_84b4f59d98cc997c, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetOutput() float32 {
	if m != nil {
		return m.Output
	}
	return 0
}

func init() {
	proto.RegisterType((*Instruction)(nil), "machine.Instruction")
	proto.RegisterType((*Result)(nil), "machine.Result")
}

func init() {
	proto.RegisterFile("machine/machine.proto", fileDescriptor_84b4f59d98cc997c)
}

var fileDescriptor_84b4f59d98cc997c = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4d, 0x4c, 0xce,
	0xc8, 0xcc, 0x4b, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x33, 0x17, 0xb7, 0x67, 0x5e, 0x71, 0x49, 0x51, 0x69, 0x72, 0x49, 0x66, 0x7e, 0x9e, 0x90,
	0x14, 0x17, 0x47, 0x7e, 0x41, 0x6a, 0x51, 0x62, 0x49, 0x7e, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x9c, 0x2f, 0x24, 0xc1, 0xc5, 0x0e, 0x66, 0xe7, 0xa5, 0x48, 0x30, 0x29, 0x30, 0x6a,
	0xb0, 0x06, 0xc1, 0xb8, 0x4a, 0x0a, 0x5c, 0x6c, 0x41, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x42, 0x62,
	0x5c, 0x6c, 0xf9, 0xa5, 0x25, 0x05, 0xa5, 0x25, 0x60, 0xdd, 0x4c, 0x41, 0x50, 0x9e, 0x91, 0x23,
	0x17, 0xbb, 0x2f, 0xc4, 0x46, 0x21, 0x33, 0x2e, 0x76, 0xd7, 0x8a, 0xd4, 0xe4, 0xd2, 0x92, 0x54,
	0x21, 0x11, 0x3d, 0x98, 0xab, 0x90, 0xdc, 0x20, 0xc5, 0x0f, 0x17, 0x85, 0x18, 0xaa, 0xc4, 0xa0,
	0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x76, 0xb9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x23, 0xc6,
	0x6c, 0xfd, 0xd2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface