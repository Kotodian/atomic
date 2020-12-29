package controller

import (
	"atomic/atomic_api/middleware"
	pbUser "atomic/atomic_proto/user"
	"atomic/internal/etcd"
	"atomic/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"net/http"
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
	// 用户登录api接口
	routerUser.POST("/login", func(ctx *gin.Context) {
		req := &pbUser.LoginRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		response, err := cli.Login(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	})
	// 用户注册api接口
	routerUser.POST("/register", func(ctx *gin.Context) {
		req := &pbUser.RegisterRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		response, err := cli.Register(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	})

	routerUser.Use(middleware.JWTAuthMiddleware())

	// 用户更新api接口
	routerUser.POST("/update", func(ctx *gin.Context) {
		req := &pbUser.UpdateRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		response, err := cli.Update(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	})

	err := srv.Init()
	if err != nil {
		panic(err)
	}
	if err = srv.Run(); err != nil {
		panic(err)
	}
}
