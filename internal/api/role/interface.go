package role

import (
	"context"
	"time"
)

type registerRoleRepo interface {
	CreateRole(ctx context.Context, title string, desc string) (int, error)
}

type playgroundRepo interface {
	GetUnixCurrentTimestamp(ctx context.Context) (int64, error)
	GetCurrentTimestamp(ctx context.Context) (time.Time, error)
}
