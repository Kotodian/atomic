// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomic/atomic_proto/user/user.proto

package user

import (
	common "atomic/atomic_proto/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "github.com/mwitkow/go-proto-validators"
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

type User struct {
	// userId 唯一标识
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// 别名
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 邮箱
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// 电话
	Phone string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	// 密码
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{1}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{2}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

type RegisterRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{3}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type RegisterResponse struct {
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{4}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "user.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "user.RegisterResponse")
}

func init() {
	proto.RegisterFile("atomic/atomic_proto/user/user.proto", fileDescriptor_604476ffdad030e1)
}

var fileDescriptor_604476ffdad030e1 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcf, 0x4e, 0xf2, 0x40,
	0x1c, 0x4c, 0x4b, 0xe1, 0x2b, 0x0b, 0x7c, 0x1f, 0xd9, 0x4f, 0x4d, 0xd3, 0x83, 0x92, 0x4a, 0x22,
	0x17, 0x5a, 0x52, 0x12, 0x13, 0x6f, 0x86, 0xb3, 0xa7, 0x1a, 0x2f, 0x5c, 0x4c, 0x69, 0xd7, 0xba,
	0x81, 0x76, 0xeb, 0x6e, 0x01, 0xe3, 0x83, 0xf8, 0x2e, 0x3e, 0x8d, 0x89, 0x4f, 0x62, 0xf6, 0x4f,
	0x6b, 0x21, 0x1a, 0xbd, 0x78, 0x61, 0xf7, 0x37, 0x33, 0x3b, 0x3b, 0x0c, 0x0b, 0x38, 0x0d, 0x0b,
	0x92, 0xe2, 0xc8, 0x93, 0xcb, 0x6d, 0x4e, 0x49, 0x41, 0xbc, 0x35, 0x43, 0x54, 0x7c, 0xb8, 0x62,
	0x86, 0x06, 0xdf, 0xdb, 0x83, 0x84, 0x90, 0x64, 0x85, 0x3c, 0x81, 0x2d, 0xd6, 0x77, 0x5e, 0x8c,
	0x58, 0x44, 0x71, 0x5e, 0x10, 0xa5, 0xb3, 0xfd, 0x74, 0x8b, 0x8b, 0x25, 0xd9, 0x7a, 0x09, 0x19,
	0x0b, 0x64, 0xbc, 0x09, 0x57, 0x38, 0x0e, 0x0b, 0x42, 0xd9, 0xe5, 0x66, 0xe2, 0x4e, 0x5d, 0xdf,
	0xab, 0x10, 0x75, 0xe6, 0xec, 0xb3, 0x00, 0x11, 0x49, 0x53, 0x92, 0xa9, 0x45, 0x0a, 0x9d, 0x67,
	0x0d, 0x18, 0x37, 0x0c, 0x51, 0xf8, 0x17, 0xe8, 0x38, 0xb6, 0xb4, 0x81, 0x36, 0x6a, 0x04, 0x3a,
	0x8e, 0xa1, 0x0d, 0x4c, 0x9e, 0x2f, 0x0b, 0x53, 0x64, 0xe9, 0x03, 0x6d, 0xd4, 0x0e, 0xaa, 0x99,
	0x73, 0x19, 0x8e, 0x96, 0x82, 0x6b, 0x48, 0xae, 0x9c, 0xe1, 0x01, 0x68, 0xa2, 0x34, 0xc4, 0x2b,
	0xcb, 0x10, 0x84, 0x1c, 0x38, 0x9a, 0xdf, 0x93, 0x0c, 0x59, 0x4d, 0x89, 0x8a, 0x81, 0xfb, 0xe4,
	0x21, 0x63, 0x5b, 0x42, 0x63, 0xab, 0x25, 0x7d, 0xca, 0xd9, 0x99, 0x83, 0xee, 0x15, 0x49, 0x70,
	0x16, 0xa0, 0x87, 0x35, 0x62, 0x05, 0x1c, 0xd6, 0xf2, 0xf0, 0x94, 0xed, 0x99, 0xf9, 0xf6, 0x7a,
	0x62, 0xe4, 0x7f, 0x1e, 0x7b, 0xb5, 0x64, 0xc3, 0x9a, 0xa3, 0xbe, 0xaf, 0xaa, 0xbc, 0xa7, 0xa0,
	0xa7, 0xbc, 0x59, 0x4e, 0x32, 0x86, 0xa0, 0x03, 0x1a, 0x14, 0x31, 0xe1, 0xdb, 0xf1, 0xfb, 0xae,
	0x6a, 0xa8, 0xa4, 0x03, 0x4e, 0x3a, 0x2f, 0x1a, 0xf8, 0x17, 0xa0, 0x04, 0xb3, 0x02, 0xd1, 0x5f,
	0x08, 0xc5, 0x55, 0xbb, 0xa5, 0x2a, 0x95, 0xc9, 0x55, 0xdf, 0xd4, 0x7b, 0xbc, 0x53, 0x6f, 0x65,
	0xdf, 0x55, 0x45, 0x3b, 0xe7, 0xa0, 0xff, 0x11, 0xfd, 0xe7, 0xdf, 0xd9, 0x7f, 0x02, 0x1d, 0xfe,
	0x38, 0xae, 0x11, 0xdd, 0xe0, 0x08, 0xc1, 0x09, 0x68, 0x8a, 0xde, 0x20, 0x74, 0xc5, 0x3b, 0xae,
	0xff, 0x40, 0xf6, 0xff, 0x1d, 0x4c, 0x5d, 0x72, 0x01, 0xcc, 0xf2, 0x62, 0x78, 0x28, 0x05, 0x7b,
	0x1d, 0xda, 0x47, 0xfb, 0xb0, 0x3c, 0x3a, 0xb3, 0xe7, 0xd6, 0x57, 0xff, 0xa2, 0x45, 0x4b, 0xec,
	0xa7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x40, 0xf7, 0x9a, 0x68, 0x03, 0x00, 0x00,
}