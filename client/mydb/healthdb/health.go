package healthdb

import (
	"context"
	"fmt"
	"time"

	"github.com/wencan/go-service-demo/client/mydb/dbinterface"
)

// GetNow 获取现在时间。
func GetNow(ctx context.Context, db dbinterface.Geter) (time.Time, error) {
	query := `SELECT NOW()`

	var now time.Time
	err := db.GetContext(ctx, &now, query)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed in select now(), error: [%w]", err)
	}

	return now, nil
}
