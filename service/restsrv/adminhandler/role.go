package adminhandler

import (
	"net/http"

	"github.com/wencan/fastrest/restserver/httpserver"
	"github.com/wencan/go-service-demo/business/admin"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
)

// RoleHandler 角色请求处理器。
type RoleHandler struct {
	business admin.RoleBusiness
}

// NewRoleHandler 创建角色业务处理器。
func NewRoleHandler(mydb dbinterface.Execer) *RoleHandler {
	return &RoleHandler{
		business: admin.RoleBusiness{
			DBx: mydb,
		},
	}
}

// CreateRole 创建角色。
func (roleBusiness RoleHandler) CreateRole() http.HandlerFunc {
	return httpserver.NewGenericsHandler(roleBusiness.business.CreateRole)
}

// SearchRoles 搜索角色。
func (roleBusiness RoleHandler) SearchRoles() http.HandlerFunc {
	return httpserver.NewGenericsHandler(roleBusiness.business.SearchRoles)
}
