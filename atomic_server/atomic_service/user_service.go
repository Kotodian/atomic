package atomic_service

import (
	"atomic/atomic_server/proto/user"
	"context"
)

type UserService struct {
}

func (u *UserService) Login(ctx context.Context, req *user.LoginRequest, resp *user.LoginResponse) (err error) {
	return nil
}
