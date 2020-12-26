package register

import (
	"atomic/atomic_server/atomic_service"
	"atomic/atomic_server/proto/user"
	"atomic/internal/log"
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func UserRegister(ctx context.Context) {
	log.Debug("注册用户服务", ctx)

	registry := etcdv3.NewRegistry()
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Registry(registry))
	srv.Init()
	err := user.RegisterUserServiceHandler(srv.Server(), new(atomic_service.UserService))
	if err != nil {
		log.Error(err.Error(), ctx)
		return
	}
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
