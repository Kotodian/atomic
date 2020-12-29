package atomic_api

import (
	pbUser "atomic/atomic_proto/user"
	"atomic/internal/etcd"
	"atomic/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
)

// web api
func WebUser(engine *gin.Engine, port int) {

	reg := etcdv3.NewRegistry(
		registry.Addrs(etcd.ETCD1, etcd.ETCD2, etcd.ETCD3),
	)

	srv := web.NewService(
		web.Name(service.APIUser),
		web.Address(fmt.Sprintf(":%d", port)),
		web.Handler(engine),
		web.Registry(reg),
	)

	cli := pbUser.NewUserService(service.InnerUser, micro.NewService(micro.Name(service.InnerUser), micro.Registry(reg)).Client())

	routerUser := engine.Group("user")
	{
		routerUser.POST("/login", func(ctx *gin.Context) {
			req := &pbUser.LoginRequest{}
			err := ctx.ShouldBindJSON(req)
			if err != nil {
				ctx.JSON(400, nil)
				return
			}
			response, err := cli.Login(ctx, req)
			if err != nil {
				ctx.JSON(500, err)
				return
			}
			ctx.JSON(200, response)
		})

		routerUser.POST("/register", func(ctx *gin.Context) {
			req := &pbUser.RegisterRequest{}
			err := ctx.ShouldBindJSON(req)
			if err != nil {
				ctx.JSON(400, nil)
				return
			}
			response, err := cli.Register(ctx, req)
			if err != nil {
				ctx.JSON(500, err)
				return
			}
			ctx.JSON(200, response)
		})
	}
	err := srv.Init()
	if err != nil {
		panic(err)
	}
	if err = srv.Run(); err != nil {
		panic(err)
	}
}
