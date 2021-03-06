// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: atomic/atomic_proto/blog/category.proto

package blog

import (
	_ "atomic/atomic_proto/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CategoryService service

type CategoryService interface {
	// 目录列表
	List(ctx context.Context, in *CategoryListRequest, opts ...client.CallOption) (*CategoryListResponse, error)
	// 目录创建
	Create(ctx context.Context, in *CategoryCreateRequest, opts ...client.CallOption) (*CategoryCreateResponse, error)
	// 目录更新
	Update(ctx context.Context, in *CategoryUpdateRequest, opts ...client.CallOption) (*CategoryUpdateResponse, error)
	// 目录删除
	Delete(ctx context.Context, in *CategoryDeleteRequest, opts ...client.CallOption) (*CategoryDeleteResponse, error)
}

type categoryService struct {
	c    client.Client
	name string
}

func NewCategoryService(name string, c client.Client) CategoryService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "blog"
	}
	return &categoryService{
		c:    c,
		name: name,
	}
}

func (c *categoryService) List(ctx context.Context, in *CategoryListRequest, opts ...client.CallOption) (*CategoryListResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.List", in)
	out := new(CategoryListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Create(ctx context.Context, in *CategoryCreateRequest, opts ...client.CallOption) (*CategoryCreateResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Create", in)
	out := new(CategoryCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Update(ctx context.Context, in *CategoryUpdateRequest, opts ...client.CallOption) (*CategoryUpdateResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Update", in)
	out := new(CategoryUpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryService) Delete(ctx context.Context, in *CategoryDeleteRequest, opts ...client.CallOption) (*CategoryDeleteResponse, error) {
	req := c.c.NewRequest(c.name, "CategoryService.Delete", in)
	out := new(CategoryDeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CategoryService service

type CategoryServiceHandler interface {
	// 目录列表
	List(context.Context, *CategoryListRequest, *CategoryListResponse) error
	// 目录创建
	Create(context.Context, *CategoryCreateRequest, *CategoryCreateResponse) error
	// 目录更新
	Update(context.Context, *CategoryUpdateRequest, *CategoryUpdateResponse) error
	// 目录删除
	Delete(context.Context, *CategoryDeleteRequest, *CategoryDeleteResponse) error
}

func RegisterCategoryServiceHandler(s server.Server, hdlr CategoryServiceHandler, opts ...server.HandlerOption) error {
	type categoryService interface {
		List(ctx context.Context, in *CategoryListRequest, out *CategoryListResponse) error
		Create(ctx context.Context, in *CategoryCreateRequest, out *CategoryCreateResponse) error
		Update(ctx context.Context, in *CategoryUpdateRequest, out *CategoryUpdateResponse) error
		Delete(ctx context.Context, in *CategoryDeleteRequest, out *CategoryDeleteResponse) error
	}
	type CategoryService struct {
		categoryService
	}
	h := &categoryServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CategoryService{h}, opts...))
}

type categoryServiceHandler struct {
	CategoryServiceHandler
}

func (h *categoryServiceHandler) List(ctx context.Context, in *CategoryListRequest, out *CategoryListResponse) error {
	return h.CategoryServiceHandler.List(ctx, in, out)
}

func (h *categoryServiceHandler) Create(ctx context.Context, in *CategoryCreateRequest, out *CategoryCreateResponse) error {
	return h.CategoryServiceHandler.Create(ctx, in, out)
}

func (h *categoryServiceHandler) Update(ctx context.Context, in *CategoryUpdateRequest, out *CategoryUpdateResponse) error {
	return h.CategoryServiceHandler.Update(ctx, in, out)
}

func (h *categoryServiceHandler) Delete(ctx context.Context, in *CategoryDeleteRequest, out *CategoryDeleteResponse) error {
	return h.CategoryServiceHandler.Delete(ctx, in, out)
}
