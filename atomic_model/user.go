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
	// 创建博客
	CreateBlog(context.Context, *gorm.DB, Blog) error
	// 删除博客
	DeleteBlog(context.Context, *gorm.DB, Blog) error
	// 收藏博客
	CollectBlog(context.Context, *gorm.DB, Blog) error
}
