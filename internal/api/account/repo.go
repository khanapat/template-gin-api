package account

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepo struct {
	db *pgxpool.Pool
}

func NewAccountRepo(db *pgxpool.Pool) *accountRepo {
	return &accountRepo{
		db: db,
	}
}

func (r *accountRepo) InquiryAccount(ctx context.Context, request map[string]interface{}) (*[]Account, error) {
	accounts := make([]Account, 0)
	param := make([]interface{}, 0)
	query := `
		SELECT	a.id,
				a.first_name,
				a.last_name,
				a.email,
				a.balance,
				a.role_id,
				r.title,
				r.description,
				a.created_date_time,
				a.updated_date_time
		FROM public.accounts a
		LEFT JOIN public.role r ON a.role_id = r.id
		WHERE 1 = 1
	`
	count := 1
	for key, value := range request {
		query = fmt.Sprintf("%s AND a.%s = $%d", query, key, count)
		param = append(param, value)
		count++
	}
	query = fmt.Sprintf(`%s
		ORDER BY a.created_date_time DESC
	;`, query)
	rows, err := r.db.Query(ctx, query, param...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var account Account
		if err := rows.Scan(
			&account.Id,
			&account.FirstName,
			&account.LastName,
			&account.Email,
			&account.Balance,
			&account.RoleId,
			&account.RoleTitle,
			&account.RoleDescription,
			&account.CreatedDateTime,
			&account.UpdatedDateTime,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	defer rows.Close()

	return &accounts, nil
}

func (r *accountRepo) InquiryAccountById(ctx context.Context, id string) (*Account, error) {
	var account Account
	err := r.db.QueryRow(ctx, `
		SELECT	a.id,
				a.first_name,
				a.last_name,
				a.email,
				a.balance,
				a.role_id,
				r.title,
				r.description,
				a.created_date_time,
				a.updated_date_time
		FROM public.accounts a
		LEFT JOIN public.role r ON a.role_id = r.id
		WHERE a.id = $1
	;`, id).Scan(
		&account.Id,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.Balance,
		&account.RoleId,
		&account.RoleTitle,
		&account.RoleDescription,
		&account.CreatedDateTime,
		&account.UpdatedDateTime,
	)
	switch {
	case err == pgx.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &account, nil
	}
}

func (r *accountRepo) CreateAccount(ctx context.Context, id string, firstName string, lastName string, email string, balance float64, roleId int) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	_, err = tx.Exec(ctx, `
		INSERT INTO public.accounts
		(
			id,
			first_name,
			last_name,
			email,
			balance,
			role_id
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	;`, id, firstName, lastName, email, balance, roleId)
	if err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// https://pkg.go.dev/github.com/jackc/pgx/v5#hdr-Prepared_Statements
func (r *accountRepo) CreateAccountWithPrepareState(ctx context.Context, accounts []CreateAccount) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if _, err := tx.Prepare(ctx, "create-account", `
		INSERT INTO public.accounts
		(
			id,
			first_name,
			last_name,
			email,
			balance,
			role_id
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	;`); err != nil {
		return err
	}
	for _, value := range accounts {
		if _, err := tx.Exec(ctx, "create-account", value.Id, value.FirstName, value.LastName, value.Email, value.Balance, value.RoleId); err != nil {
			return err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
func (r *accountRepo) CreateAccountWithBulk(ctx context.Context, accounts []CreateAccount) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	batch := &pgx.Batch{}
	for _, account := range accounts {
		batch.Queue(`
			INSERT INTO public.accounts
			(
				id,
				first_name,
				last_name,
				email,
				balance,
				role_id
			)
			VALUES
			(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6
			)
		;`, account.Id, account.FirstName, account.LastName, account.Email, account.Balance, account.RoleId)
	}

	br := tx.SendBatch(ctx, batch)
	// https://github.com/jackc/pgx/blob/master/batch_test.go
	// exec means getting result from next query
	for i := 0; i < len(accounts); i++ {
		_, err := br.Exec()
		if err != nil {
			return err
		}
		// fmt.Println("rows affected:", result.RowsAffected())
	}
	if err := br.Close(); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// https://pkg.go.dev/github.com/jackc/pgx/v5#hdr-Copy_Protocol
func (r *accountRepo) CreateAccountWithCopyFrom(ctx context.Context, accounts []CreateAccount) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	_, err = tx.CopyFrom(
		ctx,
		pgx.Identifier{"public", "accounts"},
		[]string{"id", "first_name", "last_name", "email", "balance", "role_id"},
		pgx.CopyFromSlice(len(accounts), func(i int) ([]any, error) {
			return []any{accounts[i].Id, accounts[i].FirstName, accounts[i].LastName, accounts[i].Email, accounts[i].Balance, accounts[i].RoleId}, nil
		}),
	)
	if err != nil {
		return err
	}
	// fmt.Println("rows affected:", copyCount)
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (r *accountRepo) UpdateAccount(ctx context.Context, id string, balance float64, roleId int) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	result, err := tx.Exec(ctx, `
		UPDATE public.accounts
		SET balance = $1,
			role_id = $2,
			updated_date_time = CURRENT_TIMESTAMP
		WHERE id = $3
	;`, balance, roleId, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
