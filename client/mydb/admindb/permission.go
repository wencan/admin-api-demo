package admindb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/model/adminmodel"
)

// GetPermissionCodes 获取一批权限的代码。
func GetPermissionCodes(ctx context.Context, db dbinterface.Selector, ids []int64) ([]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	query := `
SELECT permission_code
FROM admin.admin_permission
WHERE id IN (?) AND deleted = 0
	`

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		return nil, err
	}

	var codes []string
	err = db.SelectContext(ctx, &codes, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed in get permission codes by id [%v], error: [%w]", ids, err)
	}

	return codes, nil
}

// InsertPermission 插入新权限，返回权限id。
func InsertPermission(ctx context.Context, db dbinterface.Execer, permission adminmodel.Permission) (int64, error) {
	query := `
	INSERT INTO admin_permission
	(title, permission_code)
	VALUES(:title, :permission_code);	
`
	result, err := db.NamedExecContext(ctx, query, permission)
	if err != nil {
		return 0, fmt.Errorf("failed int insert permission [%+v], error: [%w]", permission, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetPermission 根据id获得权限信息。如果没找到，返回false。
func GetPermission(ctx context.Context, db dbinterface.Geter, permissionID int64) (*adminmodel.Permission, bool, error) {
	query := `
SELECT id, title, permission_code, deleted, create_time, update_time
FROM admin_permission	
WHERE id = ? AND deleted = 0
`
	var permission adminmodel.Permission
	err := db.GetContext(ctx, &permission, query, permissionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("failed in get permission %d, error: [%w]", permissionID, err)
	}

	return &permission, true, nil
}

// SearchPermissions 搜索权限。page从第1页开始。
func SearchPermissions(ctx context.Context, db dbinterface.Selector, page, size int) ([]*adminmodel.Permission, error) {
	query := `
SELECT id, title, permission_code, deleted, create_time, update_time
FROM admin_permission	
WHERE deleted = 0
ORDER BY id 
LIMIT ? OFFSET ?
`
	var permissions []*adminmodel.Permission
	err := db.SelectContext(ctx, &permissions, query, size, (page-1)*size)
	if err != nil {
		return nil, fmt.Errorf("failed in search permission, error: [%w]", err)
	}

	return permissions, nil
}

// CountPermissions 权限总数。
func CountPermissions(ctx context.Context, db dbinterface.Geter) (int, error) {
	query := `
SELECT COUNT(*)
FROM admin_permission	
WHERE deleted = 0
`
	var count int
	err := db.GetContext(ctx, &count, query)
	if err != nil {
		return 0, fmt.Errorf("failed in get total for permissions, error: [%w]", err)
	}
	return count, nil
}
