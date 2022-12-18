package health

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/client/mydb/healthdb"
	"github.com/wencan/go-service-demo/client/rds"
	"github.com/wencan/go-service-demo/model/healthmodel"
	"github.com/wencan/go-service-demo/model/protocolmodel"
)

// HealthStatusBusiness 健康检查。
type HealthStatusBusiness struct {
	// DBx MySQL服务。
	DBx dbinterface.Geter

	// Rds redis客户端
	Rds *redis.Client
}

// GetHealthStatus 服务健康状况。
func (healthStatus HealthStatusBusiness) GetHealthStatus(ctx context.Context) (*protocolmodel.HealthStatusResponse, error) {
	ok := true
	serverTime := time.Now()

	mySQLTime, err := healthdb.GetNow(ctx, healthStatus.DBx)
	if err != nil {
		ok = false
	}

	rdsTime, err := rds.GetNow(ctx, healthStatus.Rds)
	if err != nil {
		ok = false
	}

	return &protocolmodel.HealthStatusResponse{
		Ok:         ok,
		ServerTime: serverTime,
		UpstreamStatus: healthmodel.UpstreamStatus{
			MySQLTime: mySQLTime,
			RedisTime: rdsTime,
		},
	}, nil
}
