package atomic_store

import (
	"context"
	"gorm.io/gorm"
)

type Data interface {
	// 返回自己所在的表
	TableName() string
	// 返回自身的 model
	Model() interface{}
	// 返回它的类型
	Type() string
}

type Database interface {
	Init(ctx context.Context, dsn string) (*gorm.DB, error)
}

type Connect func(context.Context, string) (*gorm.DB, error)

func (c Connect) conn(ctx context.Context, dsn string) (*gorm.DB, error) {
	return c(ctx, dsn)
}

type Store struct {
	// 配置
	option *Option
	// 数据库类型
	source dataSource
}
