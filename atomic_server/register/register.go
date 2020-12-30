package register

import (
	pbUser "atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_handler"
	"atomic/internal/etcd"
	"atomic/internal/kafka_msg"
	"atomic/internal/log"
	"atomic/internal/service"
	"atomic/internal/trace"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-plugins/registry/etcdv3"
	limmiter "github.com/micro/go-plugins/wrapper/ratelimiter/uber"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

func UserServiceRegister(port int) {
	log.Debug("注册用户服务")

	// jaeger
	t, io, err := trace.NewTracer(service.InnerUser, trace.JaegerAddr)
	if err != nil {
		log.Error(err)
		return
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)

	reg := etcdv3.NewRegistry(
		registry.Addrs(etcd.Etcds...),
	)

	srv := micro.NewService(
		micro.Name(service.InnerUser),
		micro.Address(fmt.Sprintf(":%d", port)),
		micro.WrapHandler(wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(limmiter.NewHandlerWrapper(100)),
		micro.Registry(reg),
		micro.Broker(kafka.NewBroker(func(opt *broker.Options) {
			opt.Addrs = kafka_msg.URL
		})),
	)
	srv.Init()

	err = pbUser.RegisterUserServiceHandler(srv.Server(), new(atomic_handler.UserService))

	if err != nil {
		return
	}

	err = broker.Connect()
	if err != nil {
		return
	}

	if err = srv.Run(); err != nil {
		panic(err)
	}
}
