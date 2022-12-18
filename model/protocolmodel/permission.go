package protocolmodel

import "github.com/wencan/go-service-demo/model/adminmodel"

// CreatePermissionRequest 创建权限请求。
type CreatePermissionRequest struct {
	// Name 权限名称。
	Name string `json:"name" validate:"required"`

	// Code 权限代码。
	Code string `json:"code" validate:"required"`
}

// CreatePermissionResponse 创建权限响应。
type CreatePermissionResponse struct {
	// Permission 权限。
	Permission adminmodel.Permission `json:"role"`
}

// SearchPermissionsRequest 搜索权限请求。
type SearchPermissionsRequest struct {
	// Page 页码。从1开始。
	Page int `schema:"page"`

	// Size 每页上限。
	Size int `schema:"size"`
}

// SearchPermissionsResponse 搜索权限响应。
type SearchPermissionsResponse struct {
	// Total 总数。
	Total int `json:"total"`

	// Permissions 本页权限列表。
	Permissions []*adminmodel.Permission `json:"roles"`
}
