package atomic_handler

import (
	"atomic/atomic_proto/common"
	pbUser "atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_service"
	"atomic/internal/proto_model"
	"context"
)

type UserHandler struct {
}

func (u *UserHandler) Login(ctx context.Context, req *pbUser.LoginRequest, resp *pbUser.LoginResponse) (err error) {
	// proto转换
	m := proto_model.User(&pbUser.User{
		Username: req.Username,
		Password: req.Password,
	})

	token, err := atomic_service.Login(ctx, m)

	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}

	// 返回token
	resp.Res = common.SuccessResponse(err)
	resp.Token = token

	return
}

func (u *UserHandler) Register(ctx context.Context, req *pbUser.RegisterRequest, resp *pbUser.RegisterResponse) (err error) {
	m := proto_model.User(&pbUser.User{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
	})

	err = atomic_service.Register(ctx, m)

	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}
	resp.Res = common.SuccessResponse(err)

	return
}

func (u *UserHandler) Update(ctx context.Context, req *pbUser.UpdateRequest, resp *pbUser.UpdateResponse) (err error) {
	m := proto_model.User(&pbUser.User{
		Nickname: req.Nickname,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
	})

	err = atomic_service.Update(ctx, m)

	if err != nil {
		resp.Res = common.ServerErrResponse(err)
		return
	}

	resp.Res = common.SuccessResponse(err)

	return
}

func (u *UserHandler) Logout(ctx context.Context, req *pbUser.LogoutRequest, resp *pbUser.LogoutResponse) (err error) {
	panic("implement me")
}
