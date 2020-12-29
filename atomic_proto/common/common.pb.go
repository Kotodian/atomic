// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomic_server/proto/common/common.proto

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

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1244107002d070d, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "common.Response")
}

func init() {
	proto.RegisterFile("atomic_server/proto/common/common.proto", fileDescriptor_c1244107002d070d)
}

var fileDescriptor_c1244107002d070d = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0x2c, 0xc9, 0xcf,
	0xcd, 0x4c, 0x8e, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7,
	0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x52, 0x7a, 0x60, 0x31, 0x21, 0x36, 0x08, 0x4f, 0xc9,
	0x80, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x88, 0x8b, 0x25, 0x39,
	0x3f, 0x25, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x08, 0xcc, 0x16, 0x12, 0xe0, 0x62, 0xce,
	0x2d, 0x4e, 0x97, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x93, 0xd8, 0xc0, 0x06, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x67, 0xb7, 0xda, 0x69, 0x6b, 0x00, 0x00, 0x00,
}