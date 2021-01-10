package atomic_handler

import (
	"atomic/atomic_model/blog"
	"atomic/atomic_model/user"
	pbBlog "atomic/atomic_proto/blog"
	"atomic/atomic_proto/common"
	pbUser "atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_service"
	"atomic/internal/proto_model"
	"context"
)

type BlogHandler struct {
	BlogService *atomic_service.BlogService `inject:""`
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

	err = u.BlogService.DeleteBlog(ctx, userModel, blogModel)
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

	err = u.BlogService.CreateBlog(ctx, userModel, blogModel)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}

	resp.Res = common.SuccessResponse()
	return
}
