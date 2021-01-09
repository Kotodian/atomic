package controller

import (
	pbBlog "atomic/atomic_proto/blog"
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
	"strconv"
)

func WebBlog(engine *gin.Engine, port int) {

	reg := etcdv3.NewRegistry(
		registry.Addrs(etcd.Etcds...),
	)

	srv := web.NewService(
		web.Name(service.APIBlog),
		web.Address(fmt.Sprintf(":%d", port)),
		web.Handler(engine),
		web.Registry(reg),
	)

	client := micro.NewService(
		micro.Name(service.InnerBlog),
		micro.Registry(reg),
		micro.WrapClient(hystrix.NewClientWrapper()),
		micro.Broker(kafka.NewBroker(broker.Addrs(kafka_msg.URL...))),
	)
	// 消息队列
	if Broker(client) == nil {
		panic("broker can't be nil")
	}
	err := Broker(client).Connect()
	if err != nil {
		panic("cannot connect broker")
	}

	cliBlogService := pbBlog.NewBlogService(service.InnerBlog, client.Client())
	routerBlog := engine.Group("blog")
	//routerBlog.Use(middleware.JWTAuthMiddleware())

	// 用户创建博客api接口
	routerBlog.POST("/create", func(ctx *gin.Context) {
		req := &pbBlog.BlogCreateRequest{}
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

		response, err := cliBlogService.Create(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	})
	// 用户删除博客api接口
	routerBlog.POST("/delete", func(ctx *gin.Context) {
		req := &pbBlog.BlogDeleteRequest{}
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

		response, err := cliBlogService.Delete(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	})
	// 用户收藏博客api接口
	routerBlog.POST("/collect", func(ctx *gin.Context) {
		req := &struct {
			UserId int64 `json:"user_id"`
			BlogId int64 `json:"blog_id"`
			Add    bool  `json:"add"`
		}{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		var add string
		if req.Add {
			add = "add"
		} else {
			add = "reduce"
		}
		err = client.Options().Broker.Publish("collect", &broker.Message{
			Header: map[string]string{
				"userId": strconv.FormatInt(req.UserId, 10),
				"blogId": strconv.FormatInt(req.BlogId, 10),
			},
			Body: []byte(add),
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, "收藏成功")
	})

	err = srv.Init()
	if err != nil {
		panic(err)
	}
	if err = srv.Run(); err != nil {
		panic(err)
	}
}

func Broker(srv micro.Service) broker.Broker {
	return srv.Options().Broker
}
