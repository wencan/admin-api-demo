package main

import (
	"context"
	"flag"
	"log"

	"github.com/wencan/go-service-demo/client/pool"
	"github.com/wencan/go-service-demo/conf"
	"github.com/wencan/go-service-demo/service"
)

var configFile = flag.String("config", "", "配置文件路径")

func main() {
	flag.Parse()

	// 加载配置
	err := conf.LoadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// 连接到数据库
	err = pool.ConnectMySQLWithDSN(ctx, conf.C.MySQLDataSourceName)
	if err != nil {
		panic(err)
	}
	err = pool.ConnectRedis(ctx, conf.C.RedisURL)
	if err != nil {
		panic(err)
	}

	// 启动服务
	addr, err := service.Startup(ctx, conf.C.HTTPAddr, pool.GetMySQLDBx(), pool.GetRds())
	if err != nil {
		panic(err)
	}
	log.Println("Listen running at:", addr)
	// 等待服务结束
	// 收到SIGTERM、SIGINT信号，并且处理完全部已经接受的请求。
	err = service.Wait(ctx)
	if err != nil {
		panic(err)
	}

	// 断开数据库
	pool.DisconnectMySQL(ctx)
	pool.DisconnectRedis(ctx)
}
