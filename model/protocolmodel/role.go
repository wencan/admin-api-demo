package protocolmodel

import "github.com/wencan/go-service-demo/model/adminmodel"

// CreateRoleRequest 创建角色请求。
type CreateRoleRequest struct {
	// Name 角色名称。
	Name string `json:"name" validate:"required"`
}

// CreateRoleResponse 创建角色响应。
type CreateRoleResponse struct {
	// Role 角色。
	Role adminmodel.Role `json:"role"`
}

// SearchRolesRequest 搜索角色请求。
type SearchRolesRequest struct {
	// Page 页码。从1开始。
	Page int `schema:"page"`

	// Size 每页上限。
	Size int `schema:"size"`
}

// SearchRolesResponse 搜索角色响应。
type SearchRolesResponse struct {
	// Total 总数。
	Total int `json:"total"`

	// Roles 本页角色列表。
	Roles []*adminmodel.Role `json:"roles"`
}
