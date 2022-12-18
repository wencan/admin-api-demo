package admindb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/model/adminmodel"
)

// InsertRole 插入新角色，返回角色id。
func InsertRole(ctx context.Context, db dbinterface.Execer, role adminmodel.Role) (int64, error) {
	query := `
	INSERT INTO admin_role
	(title)
	VALUES(:title);	
`
	result, err := db.NamedExecContext(ctx, query, role)
	if err != nil {
		return 0, fmt.Errorf("failed int insert role [%+v], error: [%w]", role, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetRole 根据id获得角色信息。如果没找到，返回false。
func GetRole(ctx context.Context, db dbinterface.Geter, roleID int64) (*adminmodel.Role, bool, error) {
	query := `
SELECT id, title, deleted, create_time, update_time
FROM admin_role	
WHERE id = ? AND deleted = 0
`
	var role adminmodel.Role
	err := db.GetContext(ctx, &role, query, roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("failed in get role %d, error: [%w]", roleID, err)
	}

	return &role, true, nil
}

// SearchRoles 搜索角色。page从第1页开始。
func SearchRoles(ctx context.Context, db dbinterface.Selector, page, size int) ([]*adminmodel.Role, error) {
	query := `
SELECT id, title, deleted, create_time, update_time
FROM admin_role	
WHERE deleted = 0
ORDER BY id 
LIMIT ? OFFSET ?
`
	var roles []*adminmodel.Role
	err := db.SelectContext(ctx, &roles, query, size, (page-1)*size)
	if err != nil {
		return nil, fmt.Errorf("failed in search role, error: [%w]", err)
	}

	return roles, nil
}

// CountRoles 角色总数。
func CountRoles(ctx context.Context, db dbinterface.Geter) (int, error) {
	query := `
SELECT COUNT(*)
FROM admin_role	
WHERE deleted = 0
`
	var count int
	err := db.GetContext(ctx, &count, query)
	if err != nil {
		return 0, fmt.Errorf("failed in get total for roles, error: [%w]", err)
	}
	return count, nil
}
