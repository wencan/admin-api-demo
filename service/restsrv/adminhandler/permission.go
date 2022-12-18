package adminhandler

import (
	"net/http"

	"github.com/wencan/fastrest/restserver/httpserver"
	"github.com/wencan/go-service-demo/business/admin"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/model/protocolmodel"
)

// PermissionHandler 权限请求处理器。
type PermissionHandler struct {
	business admin.PermissionBusiness
}

// NewPermissionHandler 创建权限业务处理器。
func NewPermissionHandler(mydb dbinterface.Execer) *PermissionHandler {
	return &PermissionHandler{
		business: admin.PermissionBusiness{
			DBx: mydb,
		},
	}
}

// CreatePermission 创建权限。
func (permissionBusiness PermissionHandler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req protocolmodel.CreatePermissionRequest
	err := httpserver.ReadValidateRequest(ctx, &req, r)
	if err != nil {
		httpserver.WriteResponse(ctx, w, r, nil, err)
		return
	}

	resp, err := permissionBusiness.business.CreatePermission(ctx, &req)
	httpserver.WriteResponse(ctx, w, r, resp, err)
}

// SearchPermissions 搜索权限。
func (permissionBusiness PermissionHandler) SearchPermissions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req protocolmodel.SearchPermissionsRequest
	err := httpserver.ReadValidateRequest(ctx, &req, r)
	if err != nil {
		httpserver.WriteResponse(ctx, w, r, nil, err)
		return
	}

	resp, err := permissionBusiness.business.SearchPermissions(ctx, &req)
	httpserver.WriteResponse(ctx, w, r, resp, err)
}
