package main

import (
	pbUser "atomic/atomic_proto/user"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	reg := etcdv3.NewRegistry()
	srv := micro.NewService(micro.Name("go.micro.srv.user"), micro.Registry(reg))
	cli := pbUser.NewUserService("go.micro.srv.user", srv.Client())
	srv.Init()
	resp, err := cli.Register(context.Background(), &pbUser.RegisterRequest{
		Username: "1",
		Password: "1",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
