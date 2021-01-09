package atomic_service

import (
	"atomic/atomic_model"
	blog_model "atomic/atomic_model/blog"
	"atomic/atomic_store"
	"atomic/internal/etcd"
	"context"
	"encoding/json"
	"strconv"
)

func CreateBlog(ctx context.Context, u atomic_model.User, b atomic_model.Blog) error {
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}
	cli, err := etcd.NewClient(ctx, etcd.DefaultTimeout)
	if err != nil {
		return err
	}
	defer cli.Close()

	res, err := etcd.Get(ctx, cli, strconv.FormatInt(u.GetID(), 10))
	if err != nil {
		return err
	}
	if res != nil {
		u.SetStatus(string(res))
	} else {
		err = u.Get(ctx, db)
		if err != nil {
			return err
		}
	}

	err = u.CreateBlog(ctx, db, b)
	if err != nil {
		return err
	}
	// 加一层etcd用于获取数据时使用
	blogCache := &blog_model.Cache{
		Browse:     0,
		Kudos:      0,
		Collection: 0,
	}

	value, err := json.Marshal(blogCache)
	if err != nil {
		return err
	}

	err = etcd.Put(ctx, cli, strconv.FormatInt(b.GetId(), 10), string(value), -1)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBlog(ctx context.Context, u atomic_model.User, b atomic_model.Blog) error {
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}

	cli, err := etcd.NewClient(ctx, etcd.DefaultTimeout)
	if err != nil {
		return err
	}
	defer cli.Close()

	res, err := etcd.Get(ctx, cli, u.GetUsername())
	if err != nil {
		return err
	}
	if res != nil {
		u.SetStatus(string(res))
	} else {
		err = u.Get(ctx, db)
		if err != nil {
			return err
		}
	}

	err = u.DeleteBlog(ctx, db, b)
	if err != nil {
		return err
	}

	return nil
}

func CollectBlog(ctx context.Context, u atomic_model.User, b atomic_model.Blog, add bool) error {
	cli, err := etcd.NewClient(ctx, etcd.DefaultTimeout)
	if err != nil {
		return err
	}
	defer cli.Close()

	res, err := etcd.Get(ctx, cli, u.GetUsername())
	if err != nil {
		return err
	}
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}

	if res != nil {
		u.SetStatus(string(res))
	} else {
		err = u.Get(ctx, db)
		if err != nil {
			return err
		}
	}

	err = u.CollectBlog(ctx, db, b, add)
	if err != nil {
		return err
	}
	return nil
}
