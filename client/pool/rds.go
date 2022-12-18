package pool

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var rds *redis.Client

// ConnectRedis 连接到redis。
func ConnectRedis(ctx context.Context, url string) error {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return fmt.Errorf("failed in parse redis url")
	}

	return ConnectRedisWithOptions(ctx, *opt)
}

// ConnectRedisWithOptions 连接到redis服务器。
func ConnectRedisWithOptions(ctx context.Context, opt redis.Options) error {
	r := redis.NewClient(&opt)
	_, err := r.Ping(ctx).Result()
	if err != nil {
		return err
	}

	rds = r
	return nil
}

// GetRds redis客户端。
func GetRds() *redis.Client {
	return rds
}

// DisconnectRedis 断开redis连接。
func DisconnectRedis(ctx context.Context) {
	if rds == nil {
		return
	}

	r := rds
	rds = nil

	r.Close()
}
