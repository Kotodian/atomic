package redis

import (
	"atomic/internal/log"
	"context"
	"github.com/gomodule/redigo/redis"
)

func NewClient(ctx context.Context, addr string) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", addr)
	if err != nil {
		log.Error(err, ctx)
		return nil, err
	}
	return conn, nil
}
