package main

import (
	"atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	reg := etcdv3.NewRegistry()

	srv := micro.NewService(
		micro.Name("go.srv.user"),
		//micro.Address(fmt.Sprintf(":%d", port)),
		//micro.WrapHandler(wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Registry(reg),
	)
	srv.Init()

	err := user.RegisterUserServiceHandler(srv.Server(), new(atomic_handler.UserService))

	if err != nil {
		panic(err)
	}

	if err = srv.Run(); err != nil {
		panic(err)
	}
}
