package atomic_store

import (
	"atomic/internal/atomic_error"
	"atomic/internal/log"
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
	Init(ctx context.Context) (*gorm.DB, error)
	DatabaseDefaultOption()
}

type Connect func(context.Context, string) (*gorm.DB, error)

func (c Connect) Conn(ctx context.Context, dsn string) (*gorm.DB, error) {
	return c(ctx, dsn)
}

type Store struct {
	// 配置
	option *Option
	// 数据库类型
	source dataSource
}

func DefaultDatabase(ctx context.Context, database Database) (*gorm.DB, error) {
	database.DatabaseDefaultOption()
	db, err := database.Init(ctx)
	if err != nil {
		log.Error(err)
		return nil, atomic_error.ErrConnect
	}
	return db, nil
}
