package employee

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type employeeRepo struct {
	db *pgxpool.Pool
}

func NewEmployeeRepo(db *pgxpool.Pool) *employeeRepo {
	return &employeeRepo{
		db: db,
	}
}

func (r *employeeRepo) UpsertEmployee(ctx context.Context, id int, userName string, email string) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	result, err := tx.Exec(ctx, `
		INSERT INTO public.employee
		(
			id,
			username,
			email
		)
		VALUES
		(
			$1,
			$2,
			$3
		)
		ON CONFLICT (id)
		DO UPDATE SET
			username = EXCLUDED.username,
			email = EXCLUDED.email,
			updated_date_time = CURRENT_TIMESTAMP
	;`, id, userName, email)
	if err != nil {
		return err
	}
	fmt.Println(result)
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
