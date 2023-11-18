package redis

import (
	"context"
	"errors"

	goredis "github.com/redis/go-redis/v9"
)

type Redis struct {
	client *goredis.Client
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	if r == nil || r.client == nil {
		return "", errors.New("nil pointer")
	}
	return r.client.Get(ctx, key).Result()
}

func GetRedisKV(c *goredis.Client) *Redis {
	return &Redis{client: c}
}
