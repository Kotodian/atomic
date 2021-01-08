package comment

import (
	"atomic/internal/atomic_error"
	"atomic/internal/log"
	"context"
	"gorm.io/gorm"
)

type Comment struct {
	Id      int64  `json:"id" gorm:"id"`
	UserId  int64  `json:"user_id" gorm:"user_id"`
	BlogId  int64  `json:"blog_id" gorm:"blog_id"`
	Content string `json:"content"`
}

const (
	comment = "comments"
)

func (c *Comment) Create(ctx context.Context, db *gorm.DB) error {
	err := db.Table(comment).WithContext(ctx).Create(c).Error
	if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrCommentCreate
	}
	log.Info("评论成功")
	return nil
}

func (c *Comment) Delete(ctx context.Context, db *gorm.DB) error {
	err := db.Table(comment).WithContext(ctx).Where("id = ?", c.Id).Delete(c).Error
	if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrCommentDelete
	}
	log.Info("评论删除成功")
	return nil
}
