package pool

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbx *sqlx.DB

// ConnectMySQL 链接到MySQL服务器。
func ConnectMySQL(ctx context.Context, host string, port int, user, password string) error {
	config := mysql.Config{
		Addr:   fmt.Sprintf("%s:%d", host, port),
		User:   user,
		Passwd: password,
	}
	return ConnectMySQLWithDSN(ctx, config.FormatDSN())
}

// ConnectMySQLWithDSN 链接到MySQL服务器。
func ConnectMySQLWithDSN(ctx context.Context, dataSourceName string) error {
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return fmt.Errorf("failed in connect mysql server, error: [%w]", err)
	}
	dbx = db

	return nil
}

// GetMySQLDBx 获取MySQL客户端对象。
func GetMySQLDBx() *sqlx.DB {
	if dbx == nil {
		panic("Not connected to MySQL server")
	}

	return dbx
}

// DisconnectMySQL 断开MySQL连接。
func DisconnectMySQL(ctx context.Context) {
	if dbx == nil {
		return
	}

	db := dbx
	dbx = nil
	db.Close()
}
