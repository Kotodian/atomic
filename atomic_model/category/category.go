package category

import (
	"atomic/internal/atomic_error"
	"atomic/internal/log"
	"context"
	"gorm.io/gorm"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

const (
	category = "category"
)

func (c *Category) Insert(ctx context.Context, db *gorm.DB) error {
	err := db.Table(category).WithContext(ctx).Create(c).Error
	if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrCreateCategory
	}
	return nil
}

func (c *Category) Update(ctx context.Context, db *gorm.DB) error {
	err := db.Table(category).WithContext(ctx).Where("id = ?", c.ID).Updates(c).Error
	if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrUpdateCategory
	}
	return nil
}

func List(ctx context.Context, db *gorm.DB) (categories []*Category, err error) {
	err = db.Table(category).WithContext(ctx).Find(&categories).Error
	if err != nil {
		log.Error(err, ctx)
		return nil, err
	}
	return categories, nil
}

func (c *Category) Delete(ctx context.Context, db *gorm.DB) error {
	err := db.Table(category).WithContext(ctx).Where("id = ?", c.ID).Delete(c).Error
	if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrDeleteCategory
	}
	return nil
}
