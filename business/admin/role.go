package admin

import (
	"context"
	"fmt"

	"github.com/wencan/go-service-demo/client/mydb/admindb"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/model/adminmodel"
	"github.com/wencan/go-service-demo/model/protocolmodel"
)

// RoleBusiness 角色业务逻辑。
type RoleBusiness struct {
	// DBx MySQL对象。
	DBx dbinterface.Execer

	// Rds redis客户端。
	// Rds *redis.Client
}

// CreateRole 创建角色。
func (roleBusiness RoleBusiness) CreateRole(ctx context.Context, req *protocolmodel.CreateRoleRequest) (*protocolmodel.CreateRoleResponse, error) {
	id, err := admindb.InsertRole(ctx, roleBusiness.DBx, adminmodel.Role{Name: req.Name})
	if err != nil {
		return nil, err
	}

	role, _, err := admindb.GetRole(ctx, roleBusiness.DBx, id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, fmt.Errorf("failed in get role by id %d", id)
	}

	return &protocolmodel.CreateRoleResponse{
		Role: *role,
	}, nil
}

// SearchRoles 搜索角色。
func (roleBusiness RoleBusiness) SearchRoles(ctx context.Context, req *protocolmodel.SearchRolesRequest) (*protocolmodel.SearchRolesResponse, error) {
	total, err := admindb.CountRoles(ctx, roleBusiness.DBx)
	if err != nil {
		return nil, err
	}

	roles, err := admindb.SearchRoles(ctx, roleBusiness.DBx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	return &protocolmodel.SearchRolesResponse{
		Total: total,
		Roles: roles,
	}, nil
}
