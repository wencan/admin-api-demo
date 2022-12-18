package rds

import (
	"context"
	"fmt"
)

// KeyPrefix redis key的前缀。
const KeyPrefix = "admin"

// tokenKey token的redis key。
func tokenKey(ctx context.Context, token string) string {
	return fmt.Sprintf("%s:token:%s", KeyPrefix, token)
}
