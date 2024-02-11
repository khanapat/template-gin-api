package employee

import "context"

type getEmployeeByIdRepo interface {
	InquiryEmployeeById(ctx context.Context, id int) (*Employee, error)
}

type getEmployeeQuantityByProductNameRepo interface {
	InquiryEmployeeQuantityByProduct(ctx context.Context, product string) (string, error)
}

type editEmployeeQuantityByProductNameRepo interface {
	UpdateEmployeeQuantityByProduct(ctx context.Context, product string, quantity int) error
}

type upsertEmployeeRepo interface {
	UpsertEmployee(ctx context.Context, id int, userName string, email string) error
}
