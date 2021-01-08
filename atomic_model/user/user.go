package user

import (
	"atomic/atomic_model"
	"atomic/atomic_model/collection"
	"atomic/atomic_model/comment"
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

func (u *User) CommentBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog, comm string) error {
	if u.Status != Online {
		return atomic_error.ErrUserNotOnline
	}

	modelComment := &comment.Comment{}
	modelComment.UserId = u.Id
	modelComment.BlogId = blog.GetId()
	modelComment.Content = comm
	err := modelComment.Create(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BrowseBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog) error {
	if u.Status != Online {
		return atomic_error.ErrUserNotOnline
	}

	err := blog.UpdateBrowse(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) KudosBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog, add bool) error {
	if u.Status != Online {
		return atomic_error.ErrUserNotOnline
	}

	err := blog.UpdateKudos(ctx, db, add)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) NickName(ctx context.Context, db *gorm.DB) (string, error) {
	tmp := &User{}

	err := db.WithContext(ctx).Table(user).Where("username = ?", u.Username).First(&tmp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", atomic_error.ErrUserNotExists
	}
	return tmp.Nickname, nil
}

func (u *User) Login(ctx context.Context, db *gorm.DB) error {
	tmp := &User{}

	err := db.WithContext(ctx).Table(user).Where("username = ?", u.Username).First(&tmp).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return atomic_error.ErrUserNotExists
	}

	if tmp.Password != encrypt.MD5(u.Password) {
		log.Error(atomic_error.ErrPasswordWrong, ctx)
		return atomic_error.ErrPasswordWrong
	}
	u.Id = tmp.Id
	u.Nickname = tmp.Nickname
	u.Phone = tmp.Phone
	u.Email = tmp.Email

	log.Info("登录成功", ctx)

	return nil
}

func (u *User) CreateBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog) error {

	if u.Status != Online {
		return atomic_error.ErrUserNotOnline
	}

	err := blog.Insert(ctx, db, u.Username)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog) error {

	if u.Status != Online {
		return atomic_error.ErrUserNotOnline
	}

	err := blog.Delete(ctx, db)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) CollectBlog(ctx context.Context, db *gorm.DB, blog atomic_model.Blog, add bool) error {
	collect := &collection.Collection{}
	collect.BlogId = blog.GetId()
	collect.UserId = u.Id

	if u.Status != Online {
		return atomic_error.ErrUserNotOnline
	}

	if add {
		err := collect.Create(ctx, db)
		if err != nil {
			return err
		}
	} else {
		err := collect.Delete(ctx, db)
		if err != nil {
			return err
		}
	}

	err := blog.UpdateCollection(ctx, db, add)
	if err != nil {
		return err
	}

	return nil
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
	if u.Status != Online {
		return atomic_error.ErrUserNotOnline
	}
	var err error
	if u.Id > 0 {
		err = db.WithContext(ctx).Table(user).Where("id = ?", u.Id).Updates(u).Error
	} else if u.Username != "" {
		err = db.WithContext(ctx).Table(user).Where("username = ?", u.Username).Updates(u).Error
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error(err, ctx)
		return atomic_error.ErrUserNotExists
	} else if err != nil {
		log.Error(err, ctx)
		return atomic_error.ErrUpdate
	}

	return nil
}

func (u *User) Get(ctx context.Context, db *gorm.DB) error {
	var err error
	if u.Id > 0 {
		err = db.WithContext(ctx).Table(user).Where("id = ?", u.Id).Updates(u).Error
	} else if u.Username != "" {
		err = db.WithContext(ctx).Table(user).Where("username = ?", u.Username).Updates(u).Error
	}
	if err != nil {
		return err
	}
	return nil
}
