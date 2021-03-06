// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomic/atomic_proto/blog/category.proto

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

type CategoryListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryListRequest) Reset()         { *m = CategoryListRequest{} }
func (m *CategoryListRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryListRequest) ProtoMessage()    {}
func (*CategoryListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d183b1c83caeed69, []int{0}
}

func (m *CategoryListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryListRequest.Unmarshal(m, b)
}
func (m *CategoryListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryListRequest.Marshal(b, m, deterministic)
}
func (m *CategoryListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryListRequest.Merge(m, src)
}
func (m *CategoryListRequest) XXX_Size() int {
	return xxx_messageInfo_CategoryListRequest.Size(m)
}
func (m *CategoryListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryListRequest proto.InternalMessageInfo

type CategoryListResponse struct {
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	Categories           []*Category      `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CategoryListResponse) Reset()         { *m = CategoryListResponse{} }
func (m *CategoryListResponse) String() string { return proto.CompactTextString(m) }
func (*CategoryListResponse) ProtoMessage()    {}
func (*CategoryListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d183b1c83caeed69, []int{1}
}

func (m *CategoryListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryListResponse.Unmarshal(m, b)
}
func (m *CategoryListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryListResponse.Marshal(b, m, deterministic)
}
func (m *CategoryListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryListResponse.Merge(m, src)
}
func (m *CategoryListResponse) XXX_Size() int {
	return xxx_messageInfo_CategoryListResponse.Size(m)
}
func (m *CategoryListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryListResponse proto.InternalMessageInfo

func (m *CategoryListResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

func (m *CategoryListResponse) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CategoryCreateRequest struct {
	// 目录名称
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryCreateRequest) Reset()         { *m = CategoryCreateRequest{} }
func (m *CategoryCreateRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryCreateRequest) ProtoMessage()    {}
func (*CategoryCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d183b1c83caeed69, []int{2}
}

func (m *CategoryCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryCreateRequest.Unmarshal(m, b)
}
func (m *CategoryCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryCreateRequest.Marshal(b, m, deterministic)
}
func (m *CategoryCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryCreateRequest.Merge(m, src)
}
func (m *CategoryCreateRequest) XXX_Size() int {
	return xxx_messageInfo_CategoryCreateRequest.Size(m)
}
func (m *CategoryCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryCreateRequest proto.InternalMessageInfo

func (m *CategoryCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CategoryCreateResponse struct {
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CategoryCreateResponse) Reset()         { *m = CategoryCreateResponse{} }
func (m *CategoryCreateResponse) String() string { return proto.CompactTextString(m) }
func (*CategoryCreateResponse) ProtoMessage()    {}
func (*CategoryCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d183b1c83caeed69, []int{3}
}

func (m *CategoryCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryCreateResponse.Unmarshal(m, b)
}
func (m *CategoryCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryCreateResponse.Marshal(b, m, deterministic)
}
func (m *CategoryCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryCreateResponse.Merge(m, src)
}
func (m *CategoryCreateResponse) XXX_Size() int {
	return xxx_messageInfo_CategoryCreateResponse.Size(m)
}
func (m *CategoryCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryCreateResponse proto.InternalMessageInfo

func (m *CategoryCreateResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

type Category struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_d183b1c83caeed69, []int{4}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CategoryUpdateRequest struct {
	// 目录id 唯一标识
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 目录名称
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryUpdateRequest) Reset()         { *m = CategoryUpdateRequest{} }
func (m *CategoryUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryUpdateRequest) ProtoMessage()    {}
func (*CategoryUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d183b1c83caeed69, []int{5}
}

func (m *CategoryUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryUpdateRequest.Unmarshal(m, b)
}
func (m *CategoryUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryUpdateRequest.Marshal(b, m, deterministic)
}
func (m *CategoryUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryUpdateRequest.Merge(m, src)
}
func (m *CategoryUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_CategoryUpdateRequest.Size(m)
}
func (m *CategoryUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryUpdateRequest proto.InternalMessageInfo

func (m *CategoryUpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategoryUpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CategoryUpdateResponse struct {
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CategoryUpdateResponse) Reset()         { *m = CategoryUpdateResponse{} }
func (m *CategoryUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*CategoryUpdateResponse) ProtoMessage()    {}
func (*CategoryUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d183b1c83caeed69, []int{6}
}

func (m *CategoryUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryUpdateResponse.Unmarshal(m, b)
}
func (m *CategoryUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryUpdateResponse.Marshal(b, m, deterministic)
}
func (m *CategoryUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryUpdateResponse.Merge(m, src)
}
func (m *CategoryUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_CategoryUpdateResponse.Size(m)
}
func (m *CategoryUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryUpdateResponse proto.InternalMessageInfo

func (m *CategoryUpdateResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

type CategoryDeleteRequest struct {
	// 目录id 唯一标识
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryDeleteRequest) Reset()         { *m = CategoryDeleteRequest{} }
func (m *CategoryDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryDeleteRequest) ProtoMessage()    {}
func (*CategoryDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d183b1c83caeed69, []int{7}
}

func (m *CategoryDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryDeleteRequest.Unmarshal(m, b)
}
func (m *CategoryDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryDeleteRequest.Marshal(b, m, deterministic)
}
func (m *CategoryDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryDeleteRequest.Merge(m, src)
}
func (m *CategoryDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_CategoryDeleteRequest.Size(m)
}
func (m *CategoryDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryDeleteRequest proto.InternalMessageInfo

func (m *CategoryDeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CategoryDeleteResponse struct {
	Res                  *common.Response `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CategoryDeleteResponse) Reset()         { *m = CategoryDeleteResponse{} }
func (m *CategoryDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*CategoryDeleteResponse) ProtoMessage()    {}
func (*CategoryDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d183b1c83caeed69, []int{8}
}

func (m *CategoryDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryDeleteResponse.Unmarshal(m, b)
}
func (m *CategoryDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryDeleteResponse.Marshal(b, m, deterministic)
}
func (m *CategoryDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryDeleteResponse.Merge(m, src)
}
func (m *CategoryDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_CategoryDeleteResponse.Size(m)
}
func (m *CategoryDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryDeleteResponse proto.InternalMessageInfo

func (m *CategoryDeleteResponse) GetRes() *common.Response {
	if m != nil {
		return m.Res
	}
	return nil
}

func init() {
	proto.RegisterType((*CategoryListRequest)(nil), "blog.CategoryListRequest")
	proto.RegisterType((*CategoryListResponse)(nil), "blog.CategoryListResponse")
	proto.RegisterType((*CategoryCreateRequest)(nil), "blog.CategoryCreateRequest")
	proto.RegisterType((*CategoryCreateResponse)(nil), "blog.CategoryCreateResponse")
	proto.RegisterType((*Category)(nil), "blog.Category")
	proto.RegisterType((*CategoryUpdateRequest)(nil), "blog.CategoryUpdateRequest")
	proto.RegisterType((*CategoryUpdateResponse)(nil), "blog.CategoryUpdateResponse")
	proto.RegisterType((*CategoryDeleteRequest)(nil), "blog.CategoryDeleteRequest")
	proto.RegisterType((*CategoryDeleteResponse)(nil), "blog.CategoryDeleteResponse")
}

func init() {
	proto.RegisterFile("atomic/atomic_proto/blog/category.proto", fileDescriptor_d183b1c83caeed69)
}

var fileDescriptor_d183b1c83caeed69 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x65, 0xdd, 0x18, 0xfa, 0x0d, 0xa6, 0x44, 0x27, 0xb3, 0x7a, 0x18, 0xb9, 0x6c, 0x20, 0x64,
	0x30, 0x8f, 0xea, 0xa9, 0x1e, 0x3d, 0x55, 0x3c, 0x4b, 0xd7, 0x7d, 0x8c, 0xc8, 0xda, 0xd4, 0x26,
	0x0a, 0xfe, 0x1a, 0xff, 0xaa, 0xb4, 0x4d, 0x6c, 0x12, 0x2a, 0xd4, 0x53, 0xc6, 0xf7, 0xbd, 0xef,
	0xbd, 0xc7, 0x7b, 0x2b, 0x2c, 0x13, 0x25, 0x32, 0x9e, 0xae, 0x9b, 0xe7, 0xb5, 0x28, 0x85, 0x12,
	0xeb, 0xed, 0x41, 0xec, 0xd7, 0x69, 0xa2, 0x70, 0x2f, 0xca, 0x2f, 0x56, 0xcf, 0xc8, 0xa8, 0x1a,
	0x86, 0x9d, 0xf0, 0x54, 0x64, 0x99, 0xc8, 0xf5, 0xd3, 0xc0, 0xe9, 0x0c, 0xce, 0x22, 0x4d, 0xf0,
	0xc4, 0xa5, 0x8a, 0xf1, 0xfd, 0x03, 0xa5, 0xa2, 0x6f, 0x70, 0xee, 0x8e, 0x65, 0x21, 0x72, 0x89,
	0x84, 0xc2, 0xb0, 0x44, 0x39, 0x1f, 0x2c, 0x06, 0xab, 0xc9, 0xe6, 0x94, 0x69, 0x2a, 0xb3, 0x8e,
	0xab, 0x25, 0x61, 0x00, 0xda, 0x13, 0x47, 0x39, 0x0f, 0x16, 0xc3, 0xd5, 0x64, 0x33, 0x65, 0x95,
	0x2d, 0x66, 0x38, 0x63, 0x0b, 0x41, 0x6f, 0x60, 0x66, 0xe6, 0x51, 0x89, 0x89, 0x42, 0x6d, 0x82,
	0x10, 0x18, 0xe5, 0x49, 0x86, 0xb5, 0xda, 0x71, 0x5c, 0xff, 0xa6, 0xf7, 0x70, 0xe1, 0x83, 0xfb,
	0x5b, 0xa3, 0x0c, 0x8e, 0xcc, 0x35, 0x99, 0x42, 0xc0, 0x77, 0x35, 0x7c, 0x18, 0x07, 0x7c, 0xf7,
	0xab, 0x16, 0x58, 0x6a, 0x77, 0xad, 0xb5, 0x97, 0x62, 0x67, 0x59, 0xeb, 0x73, 0x6c, 0x59, 0x35,
	0xc7, 0xff, 0xb0, 0xba, 0x6c, 0xa5, 0x1f, 0xf1, 0x80, 0x7f, 0x4a, 0xdb, 0x32, 0x06, 0xd8, 0x5f,
	0x66, 0xf3, 0x1d, 0xc0, 0x89, 0x39, 0x7f, 0xc6, 0xf2, 0x93, 0xa7, 0x48, 0x1e, 0x60, 0x54, 0x95,
	0x4e, 0x2e, 0xdd, 0xd2, 0xac, 0xff, 0x47, 0x18, 0x76, 0xad, 0xb4, 0x6c, 0x04, 0xe3, 0xa6, 0x1a,
	0x72, 0xe5, 0xa2, 0x9c, 0x76, 0xc3, 0xeb, 0xee, 0x65, 0x4b, 0xd2, 0x84, 0xe6, 0x93, 0x38, 0x3d,
	0xf8, 0x24, 0x5e, 0xce, 0x11, 0x8c, 0x9b, 0x48, 0x7c, 0x12, 0x27, 0x51, 0x9f, 0xc4, 0x4d, 0x71,
	0x3b, 0xae, 0x3f, 0x94, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x0b, 0xad, 0x56, 0x82,
	0x03, 0x00, 0x00,
}
