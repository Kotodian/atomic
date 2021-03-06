// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomic/atomic_proto/blog/blog.proto

package blog

import (
	common "atomic/atomic_proto/common"
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

type BlogCreateRequest struct {
	// 标题 6到10个字之内
	// @inject_tag: validate:"required,max=11,min=8"
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// 文章内容
	// @inject_tag: validate:"required"
	Content []string `protobuf:"bytes,2,rep,name=content,proto3" json:"content,omitempty"`
	// 用户唯一标识
	// @inject_tag: validate:"required,max=12,min=7"
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogCreateRequest) Reset()         { *m = BlogCreateRequest{} }
func (m *BlogCreateRequest) String() string { return proto.CompactTextString(m) }
func (*BlogCreateRequest) ProtoMessage()    {}
func (*BlogCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92db37cf00cb4a6f, []int{0}
}

func (m *BlogCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogCreateRequest.Unmarshal(m, b)
}
func (m *BlogCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogCreateRequest.Marshal(b, m, deterministic)
}
func (m *BlogCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogCreateRequest.Merge(m, src)
}
func (m *BlogCreateRequest) XXX_Size() int {
	return xxx_messageInfo_BlogCreateRequest.Size(m)
}
func (m *BlogCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlogCreateRequest proto.InternalMessageInfo

func (m *BlogCreateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BlogCreateRequest) GetContent() []string {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *BlogCreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type BlogCreateResponse struct {
	// 基本返回信息
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlogCreateResponse) Reset()         { *m = BlogCreateResponse{} }
func (m *BlogCreateResponse) String() string { return proto.CompactTextString(m) }
func (*BlogCreateResponse) ProtoMessage()    {}
func (*BlogCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92db37cf00cb4a6f, []int{1}
}

func (m *BlogCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogCreateResponse.Unmarshal(m, b)
}
func (m *BlogCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogCreateResponse.Marshal(b, m, deterministic)
}
func (m *BlogCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogCreateResponse.Merge(m, src)
}
func (m *BlogCreateResponse) XXX_Size() int {
	return xxx_messageInfo_BlogCreateResponse.Size(m)
}
func (m *BlogCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlogCreateResponse proto.InternalMessageInfo

func (m *BlogCreateResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

type CommonBlog struct {
	// 唯一标识符
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 博客标题
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// 文章内容
	Content []string `protobuf:"bytes,3,rep,name=content,proto3" json:"content,omitempty"`
	// 属于某个用户
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonBlog) Reset()         { *m = CommonBlog{} }
func (m *CommonBlog) String() string { return proto.CompactTextString(m) }
func (*CommonBlog) ProtoMessage()    {}
func (*CommonBlog) Descriptor() ([]byte, []int) {
	return fileDescriptor_92db37cf00cb4a6f, []int{2}
}

func (m *CommonBlog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonBlog.Unmarshal(m, b)
}
func (m *CommonBlog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonBlog.Marshal(b, m, deterministic)
}
func (m *CommonBlog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonBlog.Merge(m, src)
}
func (m *CommonBlog) XXX_Size() int {
	return xxx_messageInfo_CommonBlog.Size(m)
}
func (m *CommonBlog) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonBlog.DiscardUnknown(m)
}

var xxx_messageInfo_CommonBlog proto.InternalMessageInfo

func (m *CommonBlog) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CommonBlog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CommonBlog) GetContent() []string {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *CommonBlog) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type BlogDeleteRequest struct {
	// 用户名
	// @inject_tag: validate:"required,max=12,min=7"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 博客唯一标识
	// @inject_tag: validate:"required"
	BlogId               int64    `protobuf:"varint,2,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogDeleteRequest) Reset()         { *m = BlogDeleteRequest{} }
func (m *BlogDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*BlogDeleteRequest) ProtoMessage()    {}
func (*BlogDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92db37cf00cb4a6f, []int{3}
}

func (m *BlogDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogDeleteRequest.Unmarshal(m, b)
}
func (m *BlogDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogDeleteRequest.Marshal(b, m, deterministic)
}
func (m *BlogDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogDeleteRequest.Merge(m, src)
}
func (m *BlogDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_BlogDeleteRequest.Size(m)
}
func (m *BlogDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlogDeleteRequest proto.InternalMessageInfo

func (m *BlogDeleteRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *BlogDeleteRequest) GetBlogId() int64 {
	if m != nil {
		return m.BlogId
	}
	return 0
}

type BlogDeleteResponse struct {
	// 基本返回信息
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlogDeleteResponse) Reset()         { *m = BlogDeleteResponse{} }
func (m *BlogDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*BlogDeleteResponse) ProtoMessage()    {}
func (*BlogDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92db37cf00cb4a6f, []int{4}
}

func (m *BlogDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogDeleteResponse.Unmarshal(m, b)
}
func (m *BlogDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogDeleteResponse.Marshal(b, m, deterministic)
}
func (m *BlogDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogDeleteResponse.Merge(m, src)
}
func (m *BlogDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_BlogDeleteResponse.Size(m)
}
func (m *BlogDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlogDeleteResponse proto.InternalMessageInfo

func (m *BlogDeleteResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

func init() {
	proto.RegisterType((*BlogCreateRequest)(nil), "blog.BlogCreateRequest")
	proto.RegisterType((*BlogCreateResponse)(nil), "blog.BlogCreateResponse")
	proto.RegisterType((*CommonBlog)(nil), "blog.CommonBlog")
	proto.RegisterType((*BlogDeleteRequest)(nil), "blog.BlogDeleteRequest")
	proto.RegisterType((*BlogDeleteResponse)(nil), "blog.BlogDeleteResponse")
}

func init() {
	proto.RegisterFile("atomic/atomic_proto/blog/blog.proto", fileDescriptor_92db37cf00cb4a6f)
}

var fileDescriptor_92db37cf00cb4a6f = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0x83, 0x30,
	0x14, 0xc7, 0x03, 0x9d, 0xcc, 0xbd, 0x25, 0x46, 0x5f, 0x4c, 0x46, 0x38, 0x2d, 0x78, 0x70, 0x27,
	0x96, 0xcc, 0x8b, 0xc9, 0x6e, 0xce, 0x83, 0x5e, 0xeb, 0x07, 0x20, 0x0c, 0x5e, 0x66, 0x13, 0xa0,
	0x93, 0x76, 0x7e, 0x05, 0xbf, 0xb6, 0x69, 0x0b, 0x0e, 0x9c, 0x4b, 0xbc, 0x50, 0xfe, 0xbc, 0xfe,
	0xfa, 0xfe, 0xef, 0x5f, 0xe0, 0x2e, 0xd3, 0xb2, 0x12, 0xf9, 0xd2, 0x2d, 0xe9, 0xbe, 0x91, 0x5a,
	0x2e, 0xb7, 0xa5, 0xdc, 0xd9, 0x47, 0x62, 0x35, 0x8e, 0xcc, 0x7b, 0x74, 0xff, 0xd7, 0xd6, 0x5c,
	0x56, 0x95, 0xac, 0xdb, 0xc5, 0x6d, 0x8f, 0x53, 0xb8, 0x79, 0x2a, 0xe5, 0x6e, 0xd3, 0x50, 0xa6,
	0x89, 0xd3, 0xc7, 0x81, 0x94, 0xc6, 0x5b, 0xb8, 0xd0, 0x42, 0x97, 0x14, 0x7a, 0x73, 0x6f, 0x31,
	0xe1, 0x4e, 0x60, 0x08, 0xe3, 0x5c, 0xd6, 0x9a, 0x6a, 0x1d, 0xfa, 0x73, 0xb6, 0x98, 0xf0, 0x4e,
	0x62, 0x04, 0x97, 0x07, 0x45, 0x4d, 0x9d, 0x55, 0x14, 0x32, 0x8b, 0xfc, 0xe8, 0xf8, 0x11, 0xb0,
	0xdf, 0x40, 0xed, 0x65, 0xad, 0x08, 0x63, 0x60, 0x0d, 0x29, 0x7b, 0xfe, 0x74, 0x75, 0x9d, 0xb4,
	0x96, 0xba, 0x32, 0x37, 0xc5, 0xf8, 0x1d, 0x60, 0x63, 0xbf, 0x1b, 0x1e, 0xaf, 0xc0, 0x17, 0x85,
	0x05, 0x18, 0xf7, 0x45, 0x71, 0xf4, 0xe8, 0x9f, 0xf1, 0xc8, 0xce, 0x7b, 0x1c, 0xfd, 0xf2, 0xf8,
	0xe2, 0x42, 0x78, 0xa6, 0x92, 0x8e, 0x21, 0xf4, 0x01, 0x6f, 0x08, 0xe0, 0x0c, 0xc6, 0x26, 0xe6,
	0x54, 0x14, 0xb6, 0x3d, 0xe3, 0x81, 0x91, 0xaf, 0x45, 0x37, 0x6d, 0x77, 0xd2, 0xff, 0xa7, 0x5d,
	0x7d, 0x79, 0x30, 0x35, 0xe8, 0x1b, 0x35, 0x9f, 0x22, 0x27, 0x5c, 0x43, 0xe0, 0x32, 0xc3, 0x59,
	0x62, 0xaf, 0xf7, 0xe4, 0x9a, 0xa2, 0xf0, 0xb4, 0xd0, 0x36, 0x5c, 0x43, 0xe0, 0x2c, 0xf4, 0xe1,
	0xc1, 0x78, 0x7d, 0x78, 0xe8, 0x76, 0x1b, 0xd8, 0x3f, 0xe3, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xc3, 0xeb, 0x10, 0x7e, 0x6f, 0x02, 0x00, 0x00,
}
