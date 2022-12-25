package restsrv

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/wencan/fastrest/restserver/httpserver"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/service/restsrv/adminhandler"
	"github.com/wencan/go-service-demo/service/restsrv/healthhandler"
)

// NewRestRouer 创建rest服务路由。
func NewRestRouer(mydb dbinterface.Execer, rds *redis.Client) http.HandlerFunc {
	// 配置默认的Handler工厂
	httpserver.DefaultHandlerFactory.ReadRequestFunc = httpserver.ReadValidateRequest

	// 健康检查
	healthHandler := healthhandler.NewHealthStatusHandler(mydb, rds)
	// 用户
	userHandler := adminhandler.NewUserHandler(mydb, rds)
	// 角色
	roleHandler := adminhandler.NewRoleHandler(mydb)
	// 权限
	permissionHandler := adminhandler.NewPermissionHandler(mydb)

	var mux http.ServeMux
	// 健康检查
	mux.HandleFunc("/health/status", healthHandler.GetHealthStatus)

	// 用户登录
	mux.HandleFunc("/user/login", userHandler.LoginHandlerFunc())
	// 创建用户
	mux.HandleFunc("/user/create", userHandler.CreateUserHandlerFunc())
	// 验证用户，返回信息包括权限
	mux.HandleFunc("/user/verify", userHandler.UserByTokenHandlerFunc())
	// 创建角色
	mux.HandleFunc("/role/create", roleHandler.CreateRole())
	// 搜索角色
	mux.HandleFunc("/role/search", roleHandler.SearchRoles())
	// 创建权限
	mux.HandleFunc("/permission/create", permissionHandler.CreatePermission())
	// 搜索权限
	mux.HandleFunc("/permission/search", permissionHandler.SearchPermissions())

	return mux.ServeHTTP
}
