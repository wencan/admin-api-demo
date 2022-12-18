package admindb

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
	"github.com/wencan/go-service-demo/model/adminmodel"
)

// GetUserForLogin 登录时获取用户信息。如果没找到匹配的，返回false。
func GetUserForLogin(ctx context.Context, db dbinterface.Geter, userName, passSha256 string) (*adminmodel.User, bool, error) {
	query := `
SELECT id, user_name, nice_name, password_sha256, secure_salt, deleted, create_time, update_time
FROM admin_user
WHERE user_name = ? AND password_sha256 = ? AND deleted = 0
`

	var user adminmodel.User
	err := db.GetContext(ctx, &user, query, userName, passSha256)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("failed in get user [%s] for login, error: [%w]", userName, err)
	}

	return &user, true, nil
}

// GetUserByUserName 根据username获取用户。如果没找到匹配的，返回false。
func GetUserByUserName(ctx context.Context, db dbinterface.Geter, userName string) (*adminmodel.User, bool, error) {
	query := `
SELECT id, user_name, nice_name, password_sha256, secure_salt, deleted, create_time, update_time
FROM admin_user
WHERE user_name = ? AND deleted = 0
`

	var user adminmodel.User
	err := db.GetContext(ctx, &user, query, userName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("failed in get user [%s], error: [%w]", userName, err)
	}

	return &user, true, nil
}

// GetUser 获取用户。如果没找到匹配的，返回false。
func GetUser(ctx context.Context, db dbinterface.Geter, userID int64) (*adminmodel.User, bool, error) {
	query := `
SELECT id, user_name, nice_name, password_sha256, secure_salt, deleted, create_time, update_time
FROM admin_user
WHERE id = ? AND deleted = 0
`

	var user adminmodel.User
	err := db.GetContext(ctx, &user, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("failed in get user [%d] for login, error: [%w]", userID, err)
	}

	return &user, true, nil
}

// InsertUser 插入新用户，返回用户id。如果因为username冲突而失败，返回false。
func InsertUser(ctx context.Context, db dbinterface.Execer, user adminmodel.User) (int64, bool, error) {
	query := `
	INSERT INTO admin_user
	(user_name, nice_name, password_sha256, secure_salt)
	VALUES(:user_name, :nice_name, :password_sha256, :secure_salt);	
`
	result, err := db.NamedExecContext(ctx, query, user)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return 0, false, nil
		}
		return 0, false, fmt.Errorf("failed int insert user [%+v], error: [%w]", user, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, false, err
	}
	return id, true, nil
}

// SearchUsers 搜索用户。page从第1页开始。
func SearchUsers(ctx context.Context, db dbinterface.Selector, page, size int) ([]*adminmodel.User, error) {
	query := `
SELECT id, user_name, nice_name, password_sha256, secure_salt, deleted, create_time, update_time
FROM admin_user	
WHERE deleted = 0
ORDER BY id 
LIMIT ? OFFSET ?
`
	var users []*adminmodel.User
	err := db.SelectContext(ctx, &users, query, size, (page-1)*size)
	if err != nil {
		return nil, fmt.Errorf("failed in search user, error: [%w]", err)
	}

	return users, nil
}

// CountUsers 用户总数。
func CountUsers(ctx context.Context, db dbinterface.Geter) (int, error) {
	query := `
SELECT COUNT(*)
FROM admin_user	
WHERE deleted = 0
`
	var count int
	err := db.GetContext(ctx, &count, query)
	if err != nil {
		return 0, fmt.Errorf("failed in get total for users, error: [%w]", err)
	}
	return count, nil
}
