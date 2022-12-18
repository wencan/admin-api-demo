package rds

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// GetNow redis服务器时间。
func GetNow(ctx context.Context, rds *redis.Client) (time.Time, error) {
	now, err := rds.Time(ctx).Result()
	if err != nil {
		return time.Time{}, fmt.Errorf("failed in get redis server time, error: [%w]", err)
	}

	return now, nil
}
