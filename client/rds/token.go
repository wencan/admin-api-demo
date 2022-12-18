package rds

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// SaveToken 保存token
func SaveToken(ctx context.Context, client *redis.Client, token string, userID int64, ttl time.Duration) error {
	key := tokenKey(ctx, token)
	err := client.Set(ctx, key, userID, ttl).Err()
	if err != nil {
		return fmt.Errorf("failed in save token, error: [%w]", err)
	}
	return nil
}

// UserByToken 根据token获得userID。
func UserByToken(ctx context.Context, client *redis.Client, token string) (int64, bool, error) {
	key := tokenKey(ctx, token)
	userID, err := client.Get(ctx, key).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, false, nil
		}
		return 0, false, fmt.Errorf("failed in get user by token, error: [%w]", err)
	}

	return userID, true, nil
}
