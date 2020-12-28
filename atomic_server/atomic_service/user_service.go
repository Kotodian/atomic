package atomic_service

import (
	"atomic/atomic_model/user"
	"atomic/atomic_store"
	"atomic/internal/log"
	"context"
)

func Login(ctx context.Context, user *user.User) error {
	// mysql连接 可以自行拓展
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})

	if err != nil {
		log.Error(err, ctx)
		return err
	}

	err = user.Login(ctx, db)
	if err != nil {
		log.Error(err, ctx)
		return err
	}
	return nil
}
