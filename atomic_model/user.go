package atomic_model

import (
	"context"
	"gorm.io/gorm"
)

type User interface {
	// 登录
	Login(context.Context, *gorm.DB) error
	// 更新个人信息
	Update(context.Context, *gorm.DB) error
	// 获取别名
	NickName(context.Context, *gorm.DB) (string, error)
	// 评论博客
	CommentBlog(context.Context, *gorm.DB, Blog) error
	// 收藏博客
	CollectBlog(context.Context, *gorm.DB, Blog) error
	// 查看博客
	BrowseBlog(context.Context, *gorm.DB, Blog) error
	// 点赞博客
	KudosBlog(context.Context, *gorm.DB, Blog) error
}
