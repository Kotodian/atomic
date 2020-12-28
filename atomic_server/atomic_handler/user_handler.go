package atomic_handler

import (
	"atomic/atomic_server/atomic_service"
	"atomic/atomic_server/proto/common"
	"atomic/atomic_server/proto/user"
	"atomic/internal/proto_model"
	"context"
)

type UserService struct {
}

func GetUserServiceName() string {
	return "go.atomic.srv.user"
}

func (u *UserService) Login(ctx context.Context, req *user.LoginRequest, resp *user.LoginResponse) (err error) {
	// proto转换
	m := proto_model.User(&user.User{
		Username: req.Username,
		Password: req.Password,
	})

	err = atomic_service.Login(ctx, m)

	if err != nil {
		resp.Res = common.ServerErrResponse(err)
	}

	resp.Res = common.SuccessResponse(err)

	return
}
