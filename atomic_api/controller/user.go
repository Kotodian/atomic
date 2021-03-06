package controller

import (
	pbUser "atomic/atomic_proto/user"
	"atomic/internal/etcd"
	"atomic/internal/kafka_msg"
	"atomic/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"net/http"
)

// web api
func WebUser(engine *gin.Engine, port int) {

	reg := etcdv3.NewRegistry(
		registry.Addrs(etcd.Etcds...),
	)

	srv := web.NewService(
		web.Name(service.APIUser),
		web.Address(fmt.Sprintf(":%d", port)),
		web.Handler(engine),
		web.Registry(reg),
	)

	client := micro.NewService(
		micro.Name(service.InnerUser),
		micro.Registry(reg),
		micro.WrapClient(hystrix.NewClientWrapper()),
		micro.Broker(kafka.NewBroker(broker.Addrs(kafka_msg.URL...))),
	).Client()
	// 消息队列
	if client.Options().Broker == nil {
		panic("broker can't be nil")
	}
	err := client.Options().Broker.Connect()
	if err != nil {
		panic("cannot connect broker")
	}

	cliService := pbUser.NewUserService(service.InnerUser, client)

	routerUser := engine.Group("user")
	// 用户登录api接口
	routerUser.POST("/login", func(ctx *gin.Context) {
		req := &pbUser.LoginRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err = validator.New().Struct(req)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		response, err := cliService.Login(ctx, req)
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

		err = validator.New().Struct(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		response, err := cliService.Register(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	})

	//routerUser.Use(middleware.JWTAuthMiddleware())

	// 用户更新api接口
	routerUser.POST("/update", func(ctx *gin.Context) {
		req := &pbUser.UpdateRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err = validator.New().Struct(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := cliService.Update(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	})
	routerUser.POST("/logout", func(ctx *gin.Context) {
		req := &pbUser.LogoutRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		response, err := cliService.Logout(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	})
	err = srv.Init()
	if err != nil {
		panic(err)
	}
	if err = srv.Run(); err != nil {
		panic(err)
	}
}
