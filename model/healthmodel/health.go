package healthmodel

import "time"

// UpstreamStatus 上游服务的健康状况。
type UpstreamStatus struct {
	// MySQLTime MySQL服务器时间。
	MySQLTime time.Time `json:"mysql_time"`

	// RedisTime redis
	RedisTime time.Time `json:"redis_time"`
}
