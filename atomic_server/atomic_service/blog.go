package atomic_service

import (
	"atomic/atomic_model"
	"atomic/atomic_model/user"
	"atomic/atomic_store"
	"atomic/internal/etcd"
	"context"
)

func CreateBlog(ctx context.Context, user *user.User, blog atomic_model.Blog) error {
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}
	cli, err := etcd.NewClient(ctx, etcd.DefaultTimeout)
	if err != nil {
		return err
	}
	defer cli.Close()

	res, err := etcd.Get(ctx, cli, user.Username)
	if err != nil {
		return err
	}
	if res != nil {
		user.Status = string(res)
	}

	err = user.CreateBlog(ctx, db, blog)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBlog(ctx context.Context, blog atomic_model.Blog) error {
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}

	err = blog.Delete(ctx, db)
	if err != nil {
		return err
	}
	return nil
}
