package dbinterface

import (
	"context"
	"database/sql"
)

// Getter 支持GetContext的接口。
type Geter interface {
	// GetContext 查询单个数据。如果没查到，返回错误sql.ErrNoRow。
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// Selector 支持SelectContext的接口。
type Selector interface {
	Geter

	// SelectContext 查询一批。如果没查到，不返回错误。
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// Execer 支持ExecContext的接口。
type Execer interface {
	Selector

	// ExecContext 执行写操作。
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)

	// NamedExecContext 执行写操作。支持结构体对象。
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}
