package atomic_model

import (
	"atomic/atomic_server/service"
	"atomic/internal/atomic_error/user_err"
	"atomic/internal/log"
	"context"
	"gorm.io/gorm"
)

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

	db.Table(user).Where("username = ?", u.Username).First(&tmp)
	if tmp.Password != u.Password {
		log.Error(user_err.PasswordWrong.Error(), ctx)
		return user_err.PasswordWrong
	}

	return nil
}

func (u *User) CreateBlog(ctx context.Context, db *gorm.DB, blog service.Blog) error {
	panic("implement me")
}

func (u *User) DeleteBlog(ctx context.Context, db *gorm.DB, blog service.Blog) error {
	panic("implement me")
}

func (u *User) CollectBlog(ctx context.Context, db *gorm.DB, blog service.Blog) error {
	panic("implement me")
}
