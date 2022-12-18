package admin

import (
	"context"
	"fmt"

	"github.com/wencan/go-service-demo/client/mydb/admindb"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/model/adminmodel"
	"github.com/wencan/go-service-demo/model/protocolmodel"
)

// PermissionBusiness 权限业务逻辑。
type PermissionBusiness struct {
	// DBx MySQL对象。
	DBx dbinterface.Execer

	// Rds redis客户端。
	// Rds *redis.Client
}

// CreatePermission 创建权限。
func (roleBusiness PermissionBusiness) CreatePermission(ctx context.Context, req *protocolmodel.CreatePermissionRequest) (*protocolmodel.CreatePermissionResponse, error) {
	id, err := admindb.InsertPermission(ctx, roleBusiness.DBx, adminmodel.Permission{Name: req.Name, PermissionCode: req.Code})
	if err != nil {
		return nil, err
	}

	permission, _, err := admindb.GetPermission(ctx, roleBusiness.DBx, id)
	if err != nil {
		return nil, err
	}
	if permission == nil {
		return nil, fmt.Errorf("failed in get permission by id %d", id)
	}

	return &protocolmodel.CreatePermissionResponse{
		Permission: *permission,
	}, nil
}

// SearchPermissions 搜索权限。
func (roleBusiness PermissionBusiness) SearchPermissions(ctx context.Context, req *protocolmodel.SearchPermissionsRequest) (*protocolmodel.SearchPermissionsResponse, error) {
	total, err := admindb.CountPermissions(ctx, roleBusiness.DBx)
	if err != nil {
		return nil, err
	}

	roles, err := admindb.SearchPermissions(ctx, roleBusiness.DBx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	return &protocolmodel.SearchPermissionsResponse{
		Total:       total,
		Permissions: roles,
	}, nil
}
