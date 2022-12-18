package protocolmodel

import (
	"time"

	"github.com/wencan/go-service-demo/model/healthmodel"
)

// HealthStatusResponse 健康状况响应。
type HealthStatusResponse struct {
	// Ok 服务是否正常提供服务。
	Ok bool `json:"ok"`

	// ServerTime 服务器时间。
	ServerTime time.Time `json:"server_time"`

	// UpstreamStatus 上游服务的状况。
	UpstreamStatus healthmodel.UpstreamStatus `json:"upstream_status"`
}
