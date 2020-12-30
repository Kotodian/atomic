package atomic_handler

import (
	pbBlog "atomic/atomic_proto/blog"
	"atomic/atomic_proto/common"
	pbUser "atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_service"
	"atomic/internal/proto_model"
	"context"
)

type BlogHandler struct {
}

func (u *BlogHandler) Delete(ctx context.Context, req *pbBlog.DeleteRequest, resp *pbBlog.DeleteResponse) (err error) {
	blogModel := proto_model.CommonBlog(&pbBlog.CommonBlog{
		Id: req.BlogId,
	})

	err = atomic_service.DeleteBlog(ctx, blogModel)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}

	resp.Res = common.SuccessResponse(err)
	return

}

func (u *BlogHandler) Create(ctx context.Context, req *pbBlog.CreateRequest, resp *pbBlog.CreateResponse) (err error) {
	userModel := proto_model.User(&pbUser.User{
		Username: req.Username,
	})
	blogModel := proto_model.CommonBlog(&pbBlog.CommonBlog{
		Title:   req.Title,
		Content: req.Content,
	})
	err = atomic_service.CreateBlog(ctx, userModel, blogModel)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}

	resp.Res = common.SuccessResponse(err)
	return
}
