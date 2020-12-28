package atomic_service

import (
	"atomic/atomic_model/user"
	"atomic/atomic_store"
	"context"
)

func Login(ctx context.Context, user *user.User) error {
	// mysql连接 可以自行拓展
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})

	if err != nil {
		return err
	}

	err = user.Login(ctx, db)
	if err != nil {
		return err
	}
	return nil
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
