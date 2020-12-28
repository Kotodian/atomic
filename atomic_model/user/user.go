package user

import (
	"atomic/atomic_model"
	"atomic/internal/atomic_error"
	"atomic/internal/log"
	"context"
	"errors"
	"gorm.io/gorm"
)

// 普通用户 以后添加各种用户
type User struct {
	Id       int64  `gorm:"id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Nickname string `gorm:"nickname" validate:"required"`
	Email    string `gorm:"email" validate:"email"`
	Phone    string `gorm:"phone"`
}

const user = "user"

func (u *User) Login(ctx context.Context, db *gorm.DB) error {
	tmp := &User{}

	err := db.WithContext(ctx).Table(user).Where("username = ?", u.Username).First(&tmp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return atomic_error.UserNotExists
	}

	if tmp.Password != u.Password {
		log.Error(atomic_error.PasswordWrong, ctx)
		return atomic_error.PasswordWrong
	}

	return nil
}

func (u *User) CreateBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog) error {
	panic("implement me")
}

func (u *User) DeleteBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog) error {
	panic("implement me")
}

func (u *User) CollectBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog) error {
	panic("implement me")
}
