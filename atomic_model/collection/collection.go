package collection

import (
	"atomic/internal/atomic_error"
	"atomic/internal/log"
	"context"
	"gorm.io/gorm"
)

type Collection struct {
	Id     int64 `json:"id" gorm:"id"`
	UserId int64 `json:"user_id" gorm:"user_id"`
	BlogId int64 `json:"blog_id" gorm:"blog_id"`
}

const (
	collection = "collections"
)

func (c *Collection) Create(ctx context.Context, db *gorm.DB) error {
	err := db.Table(collection).WithContext(ctx).Create(c).Error
	if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrCollectionCreate
	}
	return nil
}

func (c *Collection) Delete(ctx context.Context, db *gorm.DB) error {
	err := db.Table(collection).WithContext(ctx).Delete(c).Error
	if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrCollectionDelete
	}
	return nil
}
