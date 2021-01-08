package blog

import (
	"atomic/atomic_model"
	"atomic/internal/atomic_error"
	"atomic/internal/log"
	"context"
	"gorm.io/gorm"
	"time"
)

type CommonBlog struct {
	// Id 唯一表示
	Id int64 `gorm:"id"`
	// username 用户名
	Username string `gorm:"username"`
	// categoryId 类别
	CategoryId int64 `gorm:"category_id"`
	// title 标题
	Title string `gorm:"title"`
	// 发表时间
	CreateAt int64 `gorm:"create_at"`
	// 上次修改时间
	LastUpdateAt int64 `gorm:"last_update_at"`
	// 文章内容
	Content string `gorm:"type:longtext;content"`
	// 浏览次数
	Browse int `gorm:"browse"`
	// 点赞
	Kudos int `gorm:"kudos"`
	// 收藏
	Collection int `gorm:"collection"`
}

type Cache struct {
	Browse     int `json:"browse"`
	Kudos      int `json:"kudos"`
	Collection int `json:"collection"`
}

func (c *CommonBlog) GetId() int64 {
	return c.Id
}

const (
	blog = "blogs"
)

func (c *CommonBlog) UpdateCollection(ctx context.Context, db *gorm.DB, add bool) error {
	if add {
		c.Collection++
	} else {
		c.Collection--
	}
	err := c.Update(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommonBlog) UpdateKudos(ctx context.Context, db *gorm.DB, add bool) error {
	if add {
		c.Kudos++
	} else {
		c.Kudos--
	}
	err := c.Update(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommonBlog) UpdateBrowse(ctx context.Context, db *gorm.DB) error {
	c.Browse++
	err := c.Update(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

func (c *CommonBlog) Update(ctx context.Context, db *gorm.DB) error {
	panic("implement me")
}

func (c *CommonBlog) Delete(ctx context.Context, db *gorm.DB) error {
	err := db.Table(blog).WithContext(ctx).Delete(&c).Error
	if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrDeleteBlog
	}
	log.Info("博客删除成功", ctx)

	return nil
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
	c.CreateAt = time.Now().Unix()
	c.LastUpdateAt = time.Now().Unix()
	err := db.Table(blog).WithContext(ctx).Create(c).Error
	if err != nil {
		log.Error(err)
		return atomic_error.ErrCreateBlog
	}
	log.Info("博客创建成功", ctx)

	return nil
}
