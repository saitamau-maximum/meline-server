package redis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_ADDR     = "localhost:6379"
	REDIS_PASSWORD = ""
	REDIS_DB       = 0
)

func NewClient(ctx context.Context, opt *redis.Options) (*redis.Client, error) {
	if opt == nil {
		opt = &redis.Options{
			Addr:     REDIS_ADDR,
			Password: REDIS_PASSWORD,
			DB:       REDIS_DB,
		}
	}

	client := redis.NewClient(opt)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errors.New("[FATAL] : redis connection error")
	}

	return client, nil
}
