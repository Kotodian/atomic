package user

import (
	"atomic/atomic_model"
	"atomic/internal/atomic_error"
	"atomic/internal/encrypt"
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
	Nickname string `gorm:"nickname"`
	Email    string `gorm:"email"`
	Phone    string `gorm:"phone"`
	Status   string `gorm:"-"`
}

const user = "users"

const (
	Offline  = "offline"
	Online   = "online"
	Disabled = "disabled"
)

func (u *User) Login(ctx context.Context, db *gorm.DB) error {
	tmp := &User{}

	err := db.WithContext(ctx).Table(user).Where("username = ?", u.Username).First(&tmp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return atomic_error.ErrUserNotExists
	}

	if encrypt.MD5(tmp.Password) != u.Password {
		log.Error(atomic_error.ErrPasswordWrong, ctx)
		return atomic_error.ErrPasswordWrong
	}

	log.Info("登录成功", ctx)

	return nil
}

func (u *User) CreateBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog) error {
	tmp := &User{}

	err := db.WithContext(ctx).Table(user).Where("username = ?", u.Username).First(&tmp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return atomic_error.ErrUserNotExists
	}

	if u.Status != Online {
		return atomic_error.ErrUserNotOnline
	}

	err = blog.Insert(ctx, db, u.Username)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog) error {
	panic("implement me")
}

func (u *User) CollectBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog) error {
	panic("implement me")
}

func (u *User) Register(ctx context.Context, db *gorm.DB) error {
	u.Password = encrypt.MD5(u.Password)
	err := db.WithContext(ctx).Table(user).Create(u).Error
	if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrUserCreate
	}
	log.Info("注册成功", ctx)
	return nil
}

func (u *User) Update(ctx context.Context, db *gorm.DB) error {
	err := db.WithContext(ctx).Updates(u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(err, ctx)
		return atomic_error.ErrUserNotExists
	} else {
		log.Error(err, ctx)
		return atomic_error.ErrUpdate
	}

	return nil
}
