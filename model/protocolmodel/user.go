package protocolmodel

import (
	"time"

	"github.com/wencan/go-service-demo/model/adminmodel"
)

// CreateUserRequest 创建用户请求。
type CreateUserRequest struct {
	// UserName 用户名。
	UserName string `json:"user_name" validate:"required"`

	// NiceName 昵称
	NiceName string `json:"nice_name" validate:"required"`

	// Password 密码。
	Password string `json:"password" validate:"required"`
}

// CreateUserResponse 创建用户响应。
type CreateUserResponse struct {
	// User 用户
	User adminmodel.User `json:"user"`
}

// LoginRequest 登录请求。
type LoginRequest struct {
	// UserName
	UserName string `json:"user_name" validate:"required"`

	// Password 密码。
	Password string `json:"password" validate:"required"`
}

// LoginResponse 登录响应。
type LoginResponse struct {
	// Token 令牌。
	Token string `json:"token"`

	// ExpireTime 令牌过期时间。
	ExpireTime time.Time `json:"expire_time"`
}

// UserInfoByTokenRequest 根据令牌获取用户信息的请求。
type UserInfoByTokenRequest struct {
	// Token 令牌。
	Token string `json:"token" validate:"required"`
}

// UserInfoByTokenResponse 根据令牌获取用户信息的响应。
type UserInfoByTokenResponse struct {
	// User 用户。
	User *adminmodel.User `json:"user"`

	// PermissionCodes 权限代码列表。
	PermissionCodes []string `json:"permission_codes"`
}
