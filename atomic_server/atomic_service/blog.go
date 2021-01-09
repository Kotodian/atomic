package atomic_service

import (
	"atomic/atomic_model"
	blog_model "atomic/atomic_model/blog"
	"atomic/atomic_store"
	"atomic/internal/cache"
	"atomic/internal/etcd"
	"context"
	"encoding/json"
	"strconv"
)

func CreateBlog(ctx context.Context, user atomic_model.User, blog atomic_model.Blog, ca *cache.Cache) error {
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return err
	}
	cli, err := etcd.NewClient(ctx, etcd.DefaultTimeout)
	if err != nil {
		return err
	}
	defer cli.Close()

	res, err := etcd.Get(ctx, cli, strconv.FormatInt(user.GetID(), 10))
	if err != nil {
		return err
	}
	if res != nil {
		user.SetStatus(string(res))
	} else {
		err = user.Get(ctx, db)
		if err != nil {
			return err
		}
	}

	err = user.CreateBlog(ctx, db, blog)
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

	err = etcd.Put(ctx, cli, strconv.FormatInt(blog.GetId(), 10), string(value), -1)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBlog(ctx context.Context, u atomic_model.User, blog atomic_model.Blog, ca *cache.Cache) error {
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

	err = u.DeleteBlog(ctx, db, blog)
	if err != nil {
		return err
	}

	return nil
}

func CollectBlog(ctx context.Context, u atomic_model.User, blog atomic_model.Blog, add bool, ca *cache.Cache) error {
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

	err = u.CollectBlog(ctx, db, blog, add)
	if err != nil {
		return err
	}
	return nil
}
