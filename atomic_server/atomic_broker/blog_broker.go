package atomic_broker

import (
	"atomic/atomic_model/blog"
	"atomic/atomic_model/user"
	"atomic/atomic_server/atomic_service"
	"context"
	"errors"
	"github.com/micro/go-micro/broker"
	"github.com/opentracing/opentracing-go"
	"strconv"
)

func CollectBlog(event broker.Event) error {
	span := opentracing.StartSpan("collect-blog")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.TODO(), span)
	userIdString, ok := event.Message().Header["userId"]
	if !ok {
		return errors.New("no user")
	}
	blogIdString, ok := event.Message().Header["blogId"]
	if !ok {
		return errors.New("no blog")
	}
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	blogId, _ := strconv.ParseInt(blogIdString, 10, 64)
	if string(event.Message().Body) == "add" {
		err := atomic_service.CollectBlog(ctx, &user.User{Id: userId}, &blog.CommonBlog{Id: blogId}, true)
		if err != nil {
			return err
		}
	} else {
		err := atomic_service.CollectBlog(ctx, &user.User{Id: userId}, &blog.CommonBlog{Id: blogId}, false)
		if err != nil {
			return err
		}
	}
	return nil
}
