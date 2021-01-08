package atomic_handler

import (
	"atomic/atomic_model/blog"
	"atomic/atomic_model/user"
	pbBlog "atomic/atomic_proto/blog"
	"atomic/atomic_proto/common"
	pbUser "atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_service"
	"atomic/internal/cache"
	"atomic/internal/proto_model"
	"context"
	"time"
)

type BlogHandler struct {
	Cache *cache.Cache
}

func NewBlogHandler() *BlogHandler {
	handler := &BlogHandler{}
	handler.Cache = cache.New(10*time.Hour, 10*time.Minute)
	return handler
}

func (u *BlogHandler) Delete(ctx context.Context, req *pbBlog.BlogDeleteRequest, resp *pbBlog.BlogDeleteResponse) (err error) {
	blogModel := &blog.CommonBlog{}
	err = proto_model.ProtoToModel(req, blogModel)
	if err != nil {
		return
	}

	userModel := &user.User{}
	err = proto_model.ProtoToModel(req, userModel)
	if err != nil {
		return err
	}

	err = atomic_service.DeleteBlog(ctx, userModel, blogModel, u.Cache)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}

	resp.Res = common.SuccessResponse()
	return

}

func (u *BlogHandler) Create(ctx context.Context, req *pbBlog.BlogCreateRequest, resp *pbBlog.BlogCreateResponse) (err error) {
	userModel := &user.User{}
	blogModel := &blog.CommonBlog{}
	err = proto_model.ProtoToModel(&pbUser.User{
		Username: req.GetUsername(),
	}, userModel)

	if err != nil {
		return
	}

	err = proto_model.ProtoToModel(&pbBlog.CommonBlog{
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}, blogModel)
	if err != nil {
		return
	}

	err = atomic_service.CreateBlog(ctx, userModel, blogModel, u.Cache)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}

	resp.Res = common.SuccessResponse()
	return
}
