package main

import (
	"atomic/internal/etcd"
	"atomic/internal/kafka_msg"
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
		micro.Registry(reg),
		micro.Broker(kafkaBroker),
	)

	srv.Init()
	if srv.Options().Broker == nil {
		panic("broker cannot be nil")
	}
	err := srv.Options().Broker.Connect()
	if err != nil {
		panic("broker cannot connect")
	}

	err = srv.Options().Broker.Publish("greeter", &broker.Message{
		Header: map[string]string{
			"name": "lqk",
		},
		Body: []byte("Hello World!"),
	})
	if err != nil {
		panic(err)
	}
}
