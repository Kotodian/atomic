package atomic_handler

import (
	"atomic/atomic_model/user"
	"atomic/atomic_proto/common"
	pbUser "atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_service"
	"atomic/internal/proto_model"
	"context"
)

type UserHandler struct {
	UserService *atomic_service.UserService `inject:""`
}

func (u *UserHandler) Login(ctx context.Context, req *pbUser.LoginRequest, resp *pbUser.LoginResponse) (err error) {
	m := &user.User{}
	err = proto_model.ProtoToModel(req, m)
	if err != nil {
		return
	}
	token, err := u.UserService.Login(ctx, m)

	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}

	// 返回token
	resp.Res = common.SuccessResponse()
	resp.Token = token
	resp.UserInfo = new(pbUser.User)
	err = proto_model.ProtoToModel(m, resp.UserInfo)
	if err != nil {
		return err
	}
	return
}

func (u *UserHandler) Register(ctx context.Context, req *pbUser.RegisterRequest, resp *pbUser.RegisterResponse) (err error) {
	m := &user.User{}
	err = proto_model.ProtoToModel(req, m)
	if err != nil {
		return
	}
	err = u.UserService.Register(ctx, m)

	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	resp.Res = common.SuccessResponse()

	return
}

func (u *UserHandler) Update(ctx context.Context, req *pbUser.UpdateRequest, resp *pbUser.UpdateResponse) (err error) {
	m := &user.User{}
	err = proto_model.ProtoToModel(req, m)
	if err != nil {
		return
	}
	err = u.UserService.Update(ctx, m)

	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}

	resp.Res = common.SuccessResponse()

	return
}

func (u *UserHandler) Logout(ctx context.Context, req *pbUser.LogoutRequest, resp *pbUser.LogoutResponse) (err error) {
	m := &user.User{}
	err = proto_model.ProtoToModel(req, m)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	err = u.UserService.Logout(ctx, m)
	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	resp.Res = common.SuccessResponse()
	return nil
}
