package role

import "context"

type registerRoleRepo interface {
	CreateRole(ctx context.Context, title string, desc string) (int, error)
}
