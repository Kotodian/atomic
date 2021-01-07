package atomic_service

import (
	"atomic/atomic_model/user"
	"atomic/atomic_store"
	"atomic/internal/atomic_error"
	"atomic/internal/auth/jwt"
	"atomic/internal/date"
	"atomic/internal/etcd"
	"atomic/internal/log"
	"atomic/internal/redis"
	"context"
	"time"
)

func Login(ctx context.Context, u *user.User) (string, error) {
	// mysql连接 可以自行拓展
	db, err := atomic_store.DefaultDatabase(ctx, &atomic_store.Mysql{})
	if err != nil {
		return "", err
	}

	err = u.Login(ctx, db)
	if err != nil {
		return "", err
	}

	// 生成token

	etcdChan := make(chan bool)
	// 更新状态
	go func(c context.Context) {
		cli, e := etcd.NewClient(c, etcd.DefaultTimeout)
		if e != nil {
			log.Error(e, ctx)
			etcdChan <- false
			return
		}

		defer cli.Close()

		e = etcd.Put(c, cli, u.Username, user.Online, -1)

		if e != nil {
			log.Error(e, ctx)
			etcdChan <- false
			return
		}
		etcdChan <- true
	}(ctx)
	redisChan := make(chan bool)
	// 更新在线人数
	go func(c context.Context) {
		redisConn, e := redis.NewClient(ctx, redis.URL)
		if e != nil {
			log.Error(e, ctx)
			redisChan <- false
			return
		}
		defer redisConn.Close()

		_, e = redisConn.Do("setBit", time.Now().Format(date.Layout), u.Id, 1)
		if e != nil {
			log.Error(e, ctx)
			redisChan <- false
			return
		}

		redisChan <- true
	}(ctx)

	etcdStat := <-etcdChan
	redisStat := <-redisChan

	if etcdStat && redisStat {
		token, _ := jwt.GenToken(u.Username)
		return token, nil
	} else if !redisStat {
		return "", atomic_error.ErrUserLogin
	} else {
		return "", atomic_error.ErrUserLogin
	}
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