package account

import "context"

type registerAccountRepo interface {
	CreateAccount(ctx context.Context, id string, firstName string, lastName string, email string, balance float64, roleId int) error
}

type editAccountRepo interface {
	UpdateAccount(ctx context.Context, id string, balance float64, roleId int) error
}

type getAccountRepo interface {
	InquiryAccount(ctx context.Context, request map[string]interface{}) (*[]Account, error)
}

type getAccountByIdRepo interface {
	InquiryAccountById(ctx context.Context, id string) (*Account, error)
}

type registerAccountPrepareRepo interface {
	CreateAccountWithPrepareState(ctx context.Context, accounts []CreateAccount) error
}

type registerAccountBulkRepo interface {
	CreateAccountWithBulk(ctx context.Context, accounts []CreateAccount) error
}

type registerAccountCopyFromRepo interface {
	CreateAccountWithCopyFrom(ctx context.Context, accounts []CreateAccount) error
}

type upsertAccountRepo interface {
	UpsertAccount(ctx context.Context, id string, firstName string, lastName string, email string, balance float64, roleId int) error
}
