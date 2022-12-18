package admindb

import (
	"context"

	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
)

// GetRoleIDsByUser 获取用户关联的角色的id列表。
func GetRoleIDsByUser(ctx context.Context, db dbinterface.Selector, userID int64) ([]int64, error) {
	query := `
SELECT role_id
FROM admin_user_role
WHERE user_id = ? AND deleted = 0
`
	var ids []int64
	err := db.SelectContext(ctx, &ids, query, userID)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
