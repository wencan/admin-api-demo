package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/wencan/go-service-demo/business/common/variables"
)

// UserBusiness 用户业务逻辑接口。
type UserBusiness interface {
	// UserIDByToken 根据token取得userID。
	UserIDByToken(ctx context.Context, token string) (userID int64, err error)
}

// NewLoginMiddleware 新建登录认证中间件。required表示认证是否是必须的。
func NewLoginMiddleware(userBusiness UserBusiness, required bool) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var userID int64
			ctx := r.Context()
			token := r.Header.Get("Token")

			if token != "" {
				var err error
				userID, err = userBusiness.UserIDByToken(ctx, token)
				if err != nil {
					log.Println("failed in verify user token, error:", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}

			if userID <= 0 && required {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// next
			ctx = variables.NewContextWithUserID(ctx, userID)
			r = r.WithContext(ctx)
			next(w, r)
		}
	}
}
