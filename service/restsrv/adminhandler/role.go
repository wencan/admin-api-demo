package adminhandler

import (
	"net/http"

	"github.com/wencan/fastrest/restserver/httpserver"
	"github.com/wencan/go-service-demo/business/admin"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/model/protocolmodel"
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
func (roleBusiness RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req protocolmodel.CreateRoleRequest
	err := httpserver.ReadValidateRequest(ctx, &req, r)
	if err != nil {
		httpserver.WriteResponse(ctx, w, r, nil, err)
		return
	}

	resp, err := roleBusiness.business.CreateRole(ctx, &req)
	httpserver.WriteResponse(ctx, w, r, resp, err)
}

// SearchRoles 搜索角色。
func (roleBusiness RoleHandler) SearchRoles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req protocolmodel.SearchRolesRequest
	err := httpserver.ReadValidateRequest(ctx, &req, r)
	if err != nil {
		httpserver.WriteResponse(ctx, w, r, nil, err)
		return
	}

	resp, err := roleBusiness.business.SearchRoles(ctx, &req)
	httpserver.WriteResponse(ctx, w, r, resp, err)
}
