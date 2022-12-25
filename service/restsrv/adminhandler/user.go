package adminhandler

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/wencan/fastrest/restserver/httpserver"
	"github.com/wencan/go-service-demo/business/admin"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/model/protocolmodel"
)

// UserHandler 用户业务逻辑的请求处理器。
type UserHandler struct {
	business admin.UserBusiness
}

// NewUserHandler 创建用户逻辑请求处理器。
func NewUserHandler(mydb dbinterface.Execer, rds *redis.Client) *UserHandler {
	return &UserHandler{
		business: admin.UserBusiness{
			DBx: mydb,
			Rds: rds,
		},
	}
}

// Login 登录验证处理。
func (userHandler UserHandler) LoginHandlerFunc() http.HandlerFunc {
	handling := httpserver.GenericsHandling[protocolmodel.LoginRequest, protocolmodel.LoginResponse](userHandler.business.Login)
	return httpserver.NewHandler(handling)
}

// CreateUserHandlerFunc 创建用户。
func (userHandler UserHandler) CreateUserHandlerFunc() http.HandlerFunc {
	handling := httpserver.GenericsHandling[protocolmodel.CreateUserRequest, protocolmodel.CreateUserResponse](userHandler.business.CreateUser)
	return httpserver.NewHandler(handling)
}

// UserByTokenHandlerFunc 验证用户。
func (userHandler UserHandler) UserByTokenHandlerFunc() http.HandlerFunc {
	handling := httpserver.GenericsHandling[protocolmodel.UserInfoByTokenRequest, protocolmodel.UserInfoByTokenResponse](userHandler.business.UserInfoByToken)
	return httpserver.NewHandler(handling)
}
