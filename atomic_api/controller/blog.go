package controller

import (
	pbBlog "atomic/atomic_proto/blog"
	"atomic/internal/etcd"
	"atomic/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"net/http"
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
	).Client()

	cliBlogService := pbBlog.NewBlogService(service.InnerBlog, client)
	cliCategoryService := pbBlog.NewCategoryService(service.InnerBlog, client)
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
	routerCategory := engine.Group("category")
	routerCategory.GET("/list", func(ctx *gin.Context) {
		req := &pbBlog.CategoryListRequest{}
		response, err := cliCategoryService.List(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, response)

	})
	routerCategory.POST("/create", func(ctx *gin.Context) {
		req := &pbBlog.CategoryCreateRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		resp, err := cliCategoryService.Create(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	})
	routerCategory.POST("/update", func(ctx *gin.Context) {
		req := &pbBlog.CategoryUpdateRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		resp, err := cliCategoryService.Update(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	})
	routerCategory.POST("/delete", func(ctx *gin.Context) {
		req := &pbBlog.CategoryDeleteRequest{}
		err := ctx.ShouldBindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		resp, err := cliCategoryService.Delete(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, resp)
	})
	err := srv.Init()
	if err != nil {
		panic(err)
	}
	if err = srv.Run(); err != nil {
		panic(err)
	}
}
