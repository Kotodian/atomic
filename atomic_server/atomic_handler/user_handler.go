package atomic_handler

import (
	"atomic/atomic_proto/common"
	"atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_service"
	"atomic/internal/proto_model"
	"context"
)

type UserService struct {
}

func (u *UserService) Login(ctx context.Context, req *user.LoginRequest, resp *user.LoginResponse) (err error) {
	// proto转换
	m := proto_model.User(&user.User{
		Username: req.Username,
		Password: req.Password,
	})

	token, err := atomic_service.Login(ctx, m)

	if err != nil {
		resp.Res = common.SuccessResponse(err)
		return
	}

	// 返回token
	resp.Res = common.SuccessResponse(err)
	resp.Token = token

	return
}

func (u *UserService) Register(ctx context.Context, req *user.RegisterRequest, resp *user.RegisterResponse) (err error) {
	m := proto_model.User(&user.User{
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
