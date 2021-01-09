package register

import (
	pbBlog "atomic/atomic_proto/blog"
	pbUser "atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_broker"
	"atomic/atomic_server/atomic_handler"
	"atomic/internal/atomic_error"
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

	srv := register(port, service.InnerUser)

	err := pbUser.RegisterUserServiceHandler(srv.Server(), atomic_handler.NewUserHandler())

	if err != nil {
		return
	}

	if srv != nil {
		if err = srv.Run(); err != nil {
			panic(err)
		}
	}
}

func BlogServiceRegistry(port int) {
	log.Debug("注册博客服务")

	srv := register(port, service.InnerBlog)
	blogHandler := atomic_handler.NewBlogHandler()
	err := pbBlog.RegisterBlogServiceHandler(srv.Server(), blogHandler)

	if err != nil {
		return
	}
	categoryHandler := atomic_handler.NewCategoryHandler()
	err = pbBlog.RegisterCategoryServiceHandler(srv.Server(), categoryHandler)
	if err != nil {
		return
	}
	// 订阅收藏博客的消息队列
	_, err = srv.Options().Broker.Subscribe("collect", atomic_broker.CollectBlog)

	if err != nil {
		panic(err)
	}
	if srv != nil {
		if err = srv.Run(); err != nil {
			panic(err)
		}
	}
}

func register(port int, srvName string) micro.Service {
	// jaeger
	t, io, err := trace.NewTracer(srvName, trace.JaegerAddr)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)

	reg := etcdv3.NewRegistry(
		registry.Addrs(etcd.Etcds...),
	)

	srv := micro.NewService(
		micro.Name(srvName),
		micro.Address(fmt.Sprintf(":%d", port)),
		micro.WrapHandler(wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(limmiter.NewHandlerWrapper(100)),
		micro.Registry(reg),
		micro.Broker(kafka.NewBroker(func(opt *broker.Options) {
			opt.Addrs = kafka_msg.URL
		})),
	)

	afterStart := func() error {
		brk := srv.Options().Broker
		if err := brk.Connect(); err != nil {
			log.Fatal(atomic_error.ErrBrokerConnect.Error())
			return err
		}
		return nil
	}

	srv.Init(
		micro.AfterStart(afterStart),
	)
	return srv
}

// 返回消息队列
func Broker(srv micro.Service) broker.Broker {
	return srv.Options().Broker
}
