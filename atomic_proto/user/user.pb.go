// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomic/atomic_proto/user/user.proto

package user

import (
	common "atomic/atomic_proto/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type Status int32

const (
	// 离线
	Status_offline Status = 0
	// 在线
	Status_online Status = 1
	// 被禁用
	Status_disabled Status = 2
)

var Status_name = map[int32]string{
	0: "offline",
	1: "online",
	2: "disabled",
}

var Status_value = map[string]int32{
	"offline":  0,
	"online":   1,
	"disabled": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{0}
}

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
	// 用户名
	// @inject_tag: validate:"required"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 密码
	// @inject_tag: validate:"required"
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
	// 基本返回信息
	Res *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	// 用户token
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// 用户信息
	UserInfo             *User    `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetUserInfo() *User {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type RegisterRequest struct {
	// 用户名 长度必须大于7位小于13位
	// @inject_tag: validate:"required"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 密码 长度必须大于7位小于13位
	// @inject_tag: validate:"required"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// @inject_tag: validate:"required"
	//别名 长度必须大于8位小于13位
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 邮箱
	// @inject_tag: validate:"required,email"
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// 手机号码或者电话号码
	// @inject_tag: validate:"required"
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
	// 基本返回信息
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

type UpdateRequest struct {
	// 用户名
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 别名
	// @inject_tag: validate:"required"
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 邮箱
	// @inject_tag: validate:"required,email"
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// 手机或者电话
	// @inject_tag: validate:"required"
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	// 密码
	// @inject_tag: validate:"required"
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UpdateRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UpdateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UpdateRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UpdateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UpdateResponse struct {
	// 基本返回信息
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

type LogoutRequest struct {
	// 用户唯一标识
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// token
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{7}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LogoutRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutResponse struct {
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_604476ffdad030e1, []int{8}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

func (m *LogoutResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

func init() {
	proto.RegisterEnum("user.Status", Status_name, Status_value)
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "user.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "user.RegisterResponse")
	proto.RegisterType((*UpdateRequest)(nil), "user.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "user.UpdateResponse")
	proto.RegisterType((*LogoutRequest)(nil), "user.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "user.LogoutResponse")
}

func init() {
	proto.RegisterFile("atomic/atomic_proto/user/user.proto", fileDescriptor_604476ffdad030e1)
}

var fileDescriptor_604476ffdad030e1 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x4e, 0x62, 0xd2, 0x49, 0x5b, 0xa2, 0xa5, 0x20, 0xcb, 0xa7, 0xc8, 0x1c, 0x5a, 0x71,
	0x88, 0x51, 0x8a, 0x90, 0xb8, 0x72, 0x40, 0x42, 0xe2, 0x82, 0xab, 0x5e, 0xb8, 0x54, 0x8e, 0x3d,
	0x0e, 0xab, 0xc6, 0x3b, 0x66, 0x77, 0x0d, 0x9f, 0x81, 0xb8, 0xf0, 0x85, 0x7c, 0x08, 0xda, 0x5d,
	0x3b, 0xa9, 0x43, 0x23, 0xa5, 0x82, 0x4b, 0xbc, 0x6f, 0x66, 0xf7, 0xf9, 0xcd, 0xdb, 0x17, 0xc3,
	0x8b, 0x4c, 0x53, 0xc5, 0xf3, 0xc4, 0x3d, 0x6e, 0x6a, 0x49, 0x9a, 0x92, 0x46, 0xa1, 0xb4, 0x3f,
	0x73, 0x8b, 0xd9, 0xd0, 0xac, 0xa3, 0xd9, 0x8a, 0x68, 0xb5, 0xc6, 0xc4, 0xd6, 0x96, 0x4d, 0x99,
	0x14, 0xa8, 0x72, 0xc9, 0x6b, 0x4d, 0xed, 0xbe, 0xe8, 0xfc, 0x3e, 0xb2, 0x9c, 0xaa, 0x8a, 0x44,
	0xfb, 0x70, 0x1b, 0xe3, 0x5f, 0x1e, 0x0c, 0xaf, 0x15, 0x4a, 0x76, 0x0a, 0x3e, 0x2f, 0x42, 0x6f,
	0xe6, 0x5d, 0x0c, 0x52, 0x9f, 0x17, 0x2c, 0x82, 0xb1, 0x79, 0x97, 0xc8, 0x2a, 0x0c, 0xfd, 0x99,
	0x77, 0x71, 0x94, 0x6e, 0xb0, 0xe9, 0x09, 0x9e, 0xdf, 0xda, 0xde, 0xc0, 0xf5, 0x3a, 0xcc, 0xce,
	0x60, 0x84, 0x55, 0xc6, 0xd7, 0xe1, 0xd0, 0x36, 0x1c, 0x30, 0xd5, 0xfa, 0x0b, 0x09, 0x0c, 0x47,
	0xae, 0x6a, 0x81, 0xe1, 0xa9, 0x33, 0xa5, 0xbe, 0x93, 0x2c, 0xc2, 0xc0, 0xf1, 0x74, 0x38, 0x7e,
	0x0f, 0xc7, 0x1f, 0x69, 0xc5, 0x45, 0x8a, 0x5f, 0x1b, 0x54, 0xba, 0xa7, 0xc7, 0xfb, 0x5b, 0xcf,
	0x86, 0xc7, 0xdf, 0xe1, 0x91, 0x70, 0xd2, 0xf2, 0xa8, 0x9a, 0x84, 0x42, 0x16, 0xc3, 0x40, 0xa2,
	0xb2, 0x1c, 0x93, 0xc5, 0x74, 0xde, 0xba, 0xd1, 0xb5, 0x53, 0xd3, 0x34, 0x72, 0x35, 0xdd, 0xa2,
	0x68, 0xd9, 0x1c, 0x60, 0xe7, 0x70, 0x64, 0x5e, 0x79, 0xc3, 0x45, 0x49, 0x76, 0xee, 0xc9, 0x02,
	0xe6, 0xf6, 0x72, 0x8c, 0x83, 0x4e, 0xcf, 0x07, 0x51, 0x52, 0xfc, 0xd3, 0x83, 0x27, 0x29, 0xae,
	0xb8, 0xd2, 0x28, 0xff, 0x51, 0xff, 0xff, 0xf2, 0x3a, 0x7e, 0x03, 0xd3, 0xad, 0xa4, 0xc3, 0xad,
	0x88, 0x7f, 0x78, 0x70, 0x72, 0x5d, 0x17, 0x99, 0xc6, 0x03, 0x27, 0xd9, 0xa8, 0xf5, 0xf7, 0xa9,
	0x1d, 0xdc, 0xab, 0x76, 0xb8, 0x2f, 0x19, 0xa3, 0x9d, 0x1b, 0x7d, 0x0d, 0xa7, 0x9d, 0xa0, 0x07,
	0xcc, 0xf1, 0xc9, 0xe6, 0x80, 0x1a, 0xdd, 0x8d, 0xf1, 0x90, 0xc0, 0x6f, 0xf2, 0x30, 0xb8, 0x93,
	0x07, 0x23, 0xa4, 0xa3, 0x3c, 0x5c, 0xc8, 0xcb, 0x04, 0x82, 0x2b, 0x9d, 0xe9, 0x46, 0xb1, 0x09,
	0x3c, 0xa6, 0xb2, 0x5c, 0x73, 0x81, 0xd3, 0x47, 0x0c, 0x20, 0x20, 0x61, 0xd7, 0x1e, 0x3b, 0x86,
	0x71, 0xc1, 0x55, 0xb6, 0x5c, 0x63, 0x31, 0xf5, 0x17, 0xbf, 0x3d, 0x98, 0x98, 0x80, 0x5d, 0xa1,
	0xfc, 0xc6, 0x73, 0x64, 0xaf, 0x60, 0x64, 0x13, 0xcd, 0x98, 0x0b, 0xdf, 0xdd, 0xbf, 0x49, 0xf4,
	0xb4, 0x57, 0x6b, 0x65, 0xbd, 0x85, 0x71, 0x77, 0xf7, 0xec, 0x99, 0xdb, 0xb0, 0x13, 0xcf, 0xe8,
	0xf9, 0x6e, 0xb9, 0x3d, 0x7a, 0x09, 0x81, 0x33, 0x9b, 0xb5, 0xcc, 0xbd, 0x2c, 0x44, 0x67, 0xfd,
	0xe2, 0xf6, 0x90, 0x33, 0x86, 0x6d, 0xe5, 0x6c, 0x9d, 0xef, 0x0e, 0xf5, 0xbd, 0x7b, 0x17, 0x7d,
	0x0e, 0xf7, 0x7d, 0x01, 0x97, 0x81, 0x5d, 0x5f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xc5,
	0x51, 0x59, 0x24, 0x05, 0x00, 0x00,
}
