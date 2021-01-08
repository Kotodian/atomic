package atomic_handler

import (
	"atomic/atomic_model/category"
	pbBlog "atomic/atomic_proto/blog"
	"atomic/atomic_proto/common"
	"atomic/atomic_server/atomic_service"
	"atomic/internal/proto_model"
	"context"
)

type CategoryHandler struct {
}

func (c *CategoryHandler) List(ctx context.Context, req *pbBlog.ListReq, resp *pbBlog.ListResp) (err error) {
	categories, err := atomic_service.CommonCategoryList(ctx)
	if err != nil {
		return
	}
	resp.Categories = make([]*pbBlog.Category, 0)
	for _, cg := range categories {
		pbCategory := &pbBlog.Category{}
		err = proto_model.ProtoToModel(cg, pbCategory)
		if err != nil {
			return err
		}
		resp.Categories = append(resp.Categories, pbCategory)
	}
	resp.Res = common.SuccessResponse()
	return nil
}

func (c *CategoryHandler) Create(ctx context.Context, req *pbBlog.CreateReq, resp *pbBlog.CreateResp) (err error) {
	m := &category.Category{}
	err = proto_model.ProtoToModel(req, m)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	err = atomic_service.CategoryInsert(ctx, m)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	return nil
}

func (c *CategoryHandler) Update(ctx context.Context, req *pbBlog.UpdateReq, resp *pbBlog.UpdateResp) (err error) {
	m := &category.Category{}
	err = proto_model.ProtoToModel(req, m)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	err = atomic_service.CategoryUpdate(ctx, m)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	return nil
}

func (c *CategoryHandler) Delete(ctx context.Context, req *pbBlog.DeleteReq, resp *pbBlog.DeleteResp) (err error) {
	m := &category.Category{}
	err = proto_model.ProtoToModel(req, m)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	err = atomic_service.CategoryDelete(ctx, m)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	return nil
}
