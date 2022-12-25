package variables

import "context"

var contextKeyUserID struct{}

// NewContextWithUserID 新建一个携带userID的ctx。
func NewContextWithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, contextKeyUserID, userID)
}

// UserIDFromContext 取到认证中间件存下的userID。如果没userID，表示没通过认证。
func UserIDFromContext(ctx context.Context) int64 {
	val := ctx.Value(contextKeyUserID)
	userID, _ := val.(int64)
	return userID
}
