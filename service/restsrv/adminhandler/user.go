package adminhandler

import (
	"context"
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
	return httpserver.NewHandler(httpserver.NewHandlerFunc(func(ctx context.Context, req *protocolmodel.LoginRequest) (*protocolmodel.LoginResponse, error) {
		return userHandler.business.Login(ctx, *req)
	}, httpserver.ReadValidateRequest))
}

// CreateUserHandlerFunc 创建用户。
func (userHandler UserHandler) CreateUserHandlerFunc() http.HandlerFunc {
	return httpserver.NewHandler(httpserver.NewHandlerFunc(func(ctx context.Context, req *protocolmodel.CreateUserRequest) (*protocolmodel.CreateUserResponse, error) {
		return userHandler.business.CreateUser(ctx, *req)
	}, httpserver.ReadValidateRequest))
}

// UserByTokenHandlerFunc 验证用户。
func (userHandler UserHandler) UserByTokenHandlerFunc() http.HandlerFunc {
	return httpserver.NewHandler(httpserver.NewHandlerFunc(func(ctx context.Context, req *protocolmodel.UserInfoByTokenRequest) (*protocolmodel.UserInfoByTokenResponse, error) {
		return userHandler.business.UserInfoByToken(ctx, *req)
	}, httpserver.ReadValidateRequest))
}
