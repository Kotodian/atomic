package atomic_service

import (
	"atomic/atomic_model/user"
	"atomic/atomic_store"
	"atomic/internal/auth/jwt"
	"context"
)

func Login(ctx context.Context, user *user.User) (string, error) {
	// mysql连接 可以自行拓展
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})

	if err != nil {
		return "", err
	}

	err = user.Login(ctx, db)
	if err != nil {
		return "", err
	}

	// 生成token
	token, _ := jwt.GenToken(user.Username)
	return token, nil
}

func Register(ctx context.Context, user *user.User) error {

	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}

	err = user.Register(ctx, db)
	if err != nil {
		return err
	}

	return nil
}

func Update(ctx context.Context, user *user.User) error {
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}

	err = user.Update(ctx, db)
	if err != nil {
		return err
	}

	return nil
}
