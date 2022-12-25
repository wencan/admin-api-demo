package service

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/service/restsrv"
)

// Startup 启动服务。
func Startup(ctx context.Context, httpAddr string, mydb dbinterface.Execer, rds *redis.Client) (listenAddr string, err error) {
	return restsrv.Startup(ctx, httpAddr, mydb, rds)
}

// Shutdown 关闭。
func Shutdown(ctx context.Context) error {
	return restsrv.Shutdown(ctx)
}

// Wait 等待服务结束。
func Wait(ctx context.Context) error {
	return restsrv.Wait(ctx)
}
