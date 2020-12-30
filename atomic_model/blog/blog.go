package blog

import (
	"atomic/atomic_model"
	"atomic/internal/atomic_error"
	"atomic/internal/log"
	"context"
	"gorm.io/gorm"
)

type CommonBlog struct {
	Id       int64  `gorm:"id"`
	Username string `gorm:"username"`
	Title    string `gorm:"title"`
	Content  string `gorm:"type:longtext;content"`
}

const (
	blog = "blogs"
)

func (c *CommonBlog) Update(ctx context.Context, db *gorm.DB) error {
	panic("implement me")
}

func (c *CommonBlog) Delete(ctx context.Context, db *gorm.DB) error {
	panic("implement me")
}

func (c *CommonBlog) Proposer(ctx context.Context, db *gorm.DB) atomic_model.User {
	panic("implement me")
}

func (c *CommonBlog) ProposeTime(ctx context.Context, db *gorm.DB) int64 {
	panic("implement me")
}

func (c *CommonBlog) HasNode(ctx context.Context, db *gorm.DB) bool {
	panic("implement me")
}

func (c *CommonBlog) CreateNode(ctx context.Context, db *gorm.DB, node atomic_model.Node) error {
	panic("implement me")
}

func (c *CommonBlog) Insert(ctx context.Context, db *gorm.DB, username string) error {
	c.Username = username
	err := db.WithContext(ctx).Create(c).Error
	if err != nil {
		log.Error(err)
		return atomic_error.ErrCreateBlog
	}
	return nil
}
