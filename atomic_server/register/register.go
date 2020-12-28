package register

import (
	"atomic/atomic_server/atomic_handler"
	"atomic/atomic_server/proto/user"
	"atomic/internal/etcd"
	"atomic/internal/log"
	"atomic/internal/trace"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

func UserServiceRegister(port int) {
	log.Debug("注册用户服务")

	// jaeger
	t, io, err := trace.NewTracer(atomic_handler.GetUserServiceName(), trace.JaegerAddr)
	if err != nil {
		log.Error(err)
		return
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)

	reg := etcdv3.NewRegistry(
		registry.Addrs(etcd.ETCD1, etcd.ETCD2, etcd.ETCD3),
	)

	srv := micro.NewService(
		micro.Name(atomic_handler.GetUserServiceName()),
		micro.Address(fmt.Sprintf(":%d", port)),
		micro.WrapHandler(wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Registry(reg),
	)

	srv.Init()
	userService := new(atomic_handler.UserService)
	err = registerUser(srv, userService)
	if err != nil {
		log.Error(err)
		return
	}

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

func registerUser(srv micro.Service, atomicUser *atomic_handler.UserService) error {
	err := user.RegisterUserServiceHandler(srv.Server(), atomicUser)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
