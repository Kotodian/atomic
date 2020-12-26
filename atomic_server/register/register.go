package register

import (
	"atomic/atomic_server/atomic_service"
	"atomic/atomic_server/proto/user"
	"atomic/internal/log"
	"atomic/internal/trace"
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

func UserRegister(ctx context.Context) {
	log.Debug("注册用户服务", ctx)

	// jaeger
	// todo: 部署完jaeger使用
	t, io, err := trace.NewTracer("go.atomic.srv.user", "")
	if err != nil {
		log.Error(err.Error(), ctx)
		return
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)

	reg := etcdv3.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	srv := micro.NewService(
		micro.Name("go.atomic.srv.user"),
		micro.Address(":10001"),
		micro.WrapHandler(wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Registry(reg),
	)
	srv.Init()
	err = user.RegisterUserServiceHandler(srv.Server(), new(atomic_service.UserService))

	if err != nil {
		log.Error(err.Error(), ctx)
		return
	}

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
