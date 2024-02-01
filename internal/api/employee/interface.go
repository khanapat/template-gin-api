package employee

import "context"

type upsertEmployeeRepo interface {
	UpsertEmployee(ctx context.Context, id int, userName string, email string) error
}
