package restsrv

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/wencan/fastrest/restserver/httpserver"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
)

var server *httpserver.Server

// Startup 启动服务。
func Startup(ctx context.Context, addr string, mydb dbinterface.Execer, rds *redis.Client) error {
	router := NewRestRouer(mydb, rds)
	server = httpserver.NewServer(ctx, &http.Server{
		Addr:    addr,
		Handler: router,
	})

	_, err := server.Start(ctx) //  启动监听，开始服务。直至收到SIGTERM、SIGINT信号，或Stop被调用。
	if err != nil {
		return err
	}

	return nil
}

// Shutdown 关闭监听，不再接收新连接。
func Shutdown(ctx context.Context) error {
	if server == nil {
		return nil
	}

	server.Stop(ctx)
	return nil
}

// Wait 等待处理完全部已经接受的请求。
func Wait(ctx context.Context) error {
	if server == nil {
		return nil
	}

	return server.Wait(ctx)
}
