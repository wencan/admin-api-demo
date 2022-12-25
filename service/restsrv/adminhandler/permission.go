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
func (permissionBusiness PermissionHandler) CreatePermission() http.HandlerFunc {
	handling := httpserver.GenericsHandling[protocolmodel.CreatePermissionRequest, protocolmodel.CreatePermissionResponse](permissionBusiness.business.CreatePermission)
	return httpserver.NewHandler(handling)
}

// SearchPermissions 搜索权限。
func (permissionBusiness PermissionHandler) SearchPermissions() http.HandlerFunc {
	handling := httpserver.GenericsHandling[protocolmodel.SearchPermissionsRequest, protocolmodel.SearchPermissionsResponse](permissionBusiness.business.SearchPermissions)
	return httpserver.NewHandler(handling)
}
