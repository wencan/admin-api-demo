package admindb

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/model/adminmodel"
)

// GetPermissionIDsByRoles 获取角色的权限ID列表。
func GetPermissionIDsByRoles(ctx context.Context, db dbinterface.Selector, roleIDs []int64) ([]int64, error) {
	query := `
SELECT permission_id
FROM admin_role_permission
WHERE role_id IN (?) AND deleted = 0
`
	query, args, err := sqlx.In(query, roleIDs)
	if err != nil {
		return nil, err
	}

	var ids []int64
	err = db.SelectContext(ctx, &ids, query, args...)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// AddRolePermissions 为角色添加权限。
func AddRolePermissions(ctx context.Context, db dbinterface.Execer, roleID int64, permissionIDs []int64) error {
	query := `
INSERT INTO admin.admin_role_permission
(role_id, permission_id)
VALUES(?, ?)
	`

	var rolePermissions []*adminmodel.RolePermission
	for _, permissionID := range permissionIDs {
		rolePermissions = append(rolePermissions, &adminmodel.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
		})
	}

	_, err := db.NamedExecContext(ctx, query, rolePermissions)
	if err != nil {
		return fmt.Errorf("failed in insert role %d permission %v, error: [%w]", roleID, permissionIDs, err)
	}
	return nil
}

// DeleteRolePermissions 删除角色的权限。
func DeleteRolePermissions(ctx context.Context, db dbinterface.Execer, roleID int64, permissionIDs []int64) error {
	query := `
UPDATE admin_role_permission
SET deleted = 1
WHERE role_id = ? AND permission_id IN (?)
	`

	query, args, err := sqlx.In(query, roleID, permissionIDs)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed in delete role %d permission %v, error: [%w]", roleID, permissionIDs, err)
	}
	return nil
}
