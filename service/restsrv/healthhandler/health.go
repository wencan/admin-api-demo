package healthhandler

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/wencan/fastrest/restserver/httpserver"
	"github.com/wencan/go-service-demo/business/health"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
)

// HealthStatusHandler 健康状况的请求处理器。
type HealthStatusHandler struct {
	business health.HealthStatusBusiness
}

// NewHealthStatusHandler 创建健康检查的Handler。
func NewHealthStatusHandler(mydb dbinterface.Execer, rds *redis.Client) *HealthStatusHandler {
	return &HealthStatusHandler{business: health.HealthStatusBusiness{
		DBx: mydb,
		Rds: rds,
	}}
}

// GetHealthStatus 获取服务健康状况。
func (handler HealthStatusHandler) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := handler.business.GetHealthStatus(ctx)
	httpserver.WriteResponse(ctx, w, r, resp, err)
}
