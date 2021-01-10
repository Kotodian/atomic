package atomic_service

import (
	"atomic/atomic_model"
	"atomic/atomic_model/user"
	"atomic/internal/atomic_error"
	"atomic/internal/auth/jwt"
	"atomic/internal/date"
	"atomic/internal/etcd"
	"atomic/internal/log"
	"atomic/internal/redis"
	"context"
	"strconv"
	"time"
)

func (s *UserService) Login(ctx context.Context, u atomic_model.User) (string, error) {

	err := u.Login(ctx, s.DB)
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

		e = etcd.Put(c, cli, strconv.FormatInt(u.GetID(), 10), user.Online, -1)

		if e != nil {
			log.Error(e, ctx)
			etcdChan <- false
			return
		}
		etcdChan <- true
	}(ctx)
	redisChan := make(chan bool)

	redisConn, err := redis.NewClient(ctx, redis.URL)
	defer redisConn.Close()

	if err != nil {
		log.Error(err, ctx)
		return "", err
	}
	// 更新在线人数
	go func(c context.Context) {
		_, e := redisConn.Do("setBit", time.Now().Format(date.Layout), u.GetID(), 1)
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
		token, _ := jwt.GenToken(u.GetUsername())
		_, err := redisConn.Do("set", u.GetUsername(), token)
		if err != nil {
			return "", err
		}
		return token, nil
	} else if !redisStat {
		return "", atomic_error.ErrUserLogin
	} else {
		return "", atomic_error.ErrUserLogin
	}
}

func (s *UserService) Register(ctx context.Context, user atomic_model.User) error {

	err := user.Register(ctx, s.DB)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, user atomic_model.User) error {

	err := user.Update(ctx, s.DB)
	if err != nil {
		return err
	}

	return nil
}
func (s *UserService) Logout(ctx context.Context, u atomic_model.User) error {

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

		e = etcd.Put(c, cli, strconv.FormatInt(u.GetID(), 10), user.Offline, -1)

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

		_, e = redisConn.Do("setBit", time.Now().Format(date.Layout), u.GetID(), 0)
		if e != nil {
			log.Error(e, ctx)
			redisChan <- false
			return
		}
		_, e = redisConn.Do("del", u.GetUsername())
		redisChan <- true
	}(ctx)

	redisRes := <-redisChan
	etcdRes := <-etcdChan
	if !redisRes || !etcdRes {
		return atomic_error.ErrUserLogout
	}
	return nil
}
