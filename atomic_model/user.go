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
	CommentBlog(context.Context, *gorm.DB, Blog, string) error
	// 收藏博客
	CollectBlog(context.Context, *gorm.DB, Blog, bool) error
	// 查看博客
	BrowseBlog(context.Context, *gorm.DB, Blog) error
	// 点赞博客
	KudosBlog(context.Context, *gorm.DB, Blog, bool) error
	// 创建博客
	CreateBlog(context.Context, *gorm.DB, Blog) error
	// 删除博客
	DeleteBlog(context.Context, *gorm.DB, Blog) error
	// 获取用户名
	GetUsername() string
	// 获取id
	GetID() int64
	// 获取状态
	GetStatus() string
	// 修改状态
	SetStatus(string)
	// 注册用户
	Register(context.Context, *gorm.DB) error
	//
	Get(context.Context, *gorm.DB) error
}
