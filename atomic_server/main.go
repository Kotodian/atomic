package main

import (
	"atomic/atomic_proto/user"
	"atomic/atomic_server/atomic_handler"
	"atomic/internal/etcd"
	"atomic/internal/kafka_msg"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	reg := etcdv3.NewRegistry(registry.Addrs(etcd.Etcds...))
	kafkaBroker := kafka.NewBroker(func(o *broker.Options) {
		o.Addrs = kafka_msg.URL
	})

	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
		//micro.Address(fmt.Sprintf(":%d", port)),
		//micro.WrapHandler(wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Registry(reg),
		micro.Broker(kafkaBroker),
	)
	srv.Init()

	err := user.RegisterUserServiceHandler(srv.Server(), new(atomic_handler.UserHandler))

	if err != nil {
		panic(err)
	}

	if err := kafkaBroker.Connect(); err != nil {
		panic(err)
	}
	_, err = kafkaBroker.Subscribe("greeter", func(event broker.Event) error {
		fmt.Println(string(event.Message().Body))
		return nil
	})

	if err != nil {
		panic(err)
	}
	if err = srv.Run(); err != nil {
		panic(err)
	}
}
