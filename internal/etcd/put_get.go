/*
	将etcd与其他数据库分开,store这部分主要是存储关系型数据库
*/

package etcd

import (
	"atomic/internal/log"
	"context"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func NewClient(ctx context.Context, timeout time.Duration) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   Etcds,
		DialTimeout: timeout,
	})

	if err != nil {
		log.Error(err, ctx)
		return nil, err
	}

	return cli, nil
}

func Get(ctx context.Context, client *clientv3.Client, key string) ([]byte, error) {
	result, err := client.Get(ctx, key)
	if err != nil {
		log.Error(err, ctx)
		return nil, err
	}
	return result.Kvs[0].Value, nil
}

func Put(ctx context.Context, client *clientv3.Client, key string, value string, timeout int64) error {
	if timeout > 0 {
		lease, err := client.Grant(ctx, timeout)
		if err != nil {
			log.Error(err, ctx)
			return err
		}

		_, err = client.Put(ctx, key, value, clientv3.WithLease(lease.ID))
		if err != nil {
			log.Error(err, ctx)
			return err
		}
	} else {
		_, err := client.Put(ctx, key, value)
		if err != nil {
			log.Error(err, ctx)
			return err
		}
	}
	return nil
}

func Del(ctx context.Context, client clientv3.Client, key string) error {
	_, err := client.Delete(ctx, key)
	if err != nil {
		log.Error(err, ctx)
		return err
	}
	return nil
}
