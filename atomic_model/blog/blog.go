package blog

import (
	"atomic/atomic_model"
	"atomic/atomic_server/proto/user"
	"context"
)

type CommonBlog struct {
	Id      int64     `gorm:"id"`
	User    user.User `gorm:"foreignKey:UserId"`
	UserId  int64     `gorm:"user_id"`
	title   string    `gorm:"title"`
	content string    `gorm:"content"`
}

func (c *CommonBlog) Update(ctx context.Context) error {
	panic("implement me")
}

func (c *CommonBlog) Delete(ctx context.Context) error {
	panic("implement me")
}

func (c *CommonBlog) Proposer(ctx context.Context) atomic_model.User {
	panic("implement me")
}

func (c *CommonBlog) ProposeTime(ctx context.Context) int64 {
	panic("implement me")
}

func (c *CommonBlog) HasNode(ctx context.Context) bool {
	panic("implement me")
}

func (c *CommonBlog) CreateNode(ctx context.Context, node atomic_model.Node) error {
	panic("implement me")
}

func (c *CommonBlog) Insert(ctx context.Context, userid int64) error {
	panic("implement me")
}
