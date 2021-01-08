package atomic_model

import (
	"context"
	"gorm.io/gorm"
)

type Category interface {
	// 插入
	Insert(ctx context.Context, db *gorm.DB) error
	// 更新
	Update(ctx context.Context, db *gorm.DB) error
	// 删除
	Delete(ctx context.Context, db *gorm.DB) error
}
