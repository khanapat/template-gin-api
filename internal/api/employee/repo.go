package employee

import (
	"context"

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

// jsonb vs json (https://dev.to/ftisiot/json-vs-jsonb-in-postgresql-5cj3#:~:text=PostgreSQL%C2%AE%20offers%20two%20types,in%20a%20custom%20binary%20format)
func (r *employeeRepo) InquiryEmployeeById(ctx context.Context, id int) (*Employee, error) {
	var employee Employee
	err := r.db.QueryRow(ctx, `
		SELECT	id,
				username,
				email,
				metadata,
				job,
				created_date_time,
				updated_date_time
		FROM public.employee
		WHERE	id = $1
	;`, id).Scan(
		&employee.Id,
		&employee.Username,
		&employee.Email,
		&employee.Metadata,
		&employee.Job,
		&employee.CreatedDateTime,
		&employee.UpdatedDateTime,
	)
	switch {
	case err == pgx.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &employee, nil
	}
}

// ILIKE - case-insensitive matching
func (r *employeeRepo) InquiryEmployeeQuantityByProduct(ctx context.Context, product string) (string, error) {
	var qty string
	err := r.db.QueryRow(ctx, `
		SELECT	metadata -> 'items' ->> 'qty' AS quantity
		FROM public.employee
		WHERE	metadata -> 'items' ->> 'product' ILIKE $1
	;`, product).Scan(
		&qty,
	)
	switch {
	case err == pgx.ErrNoRows:
		return qty, nil
	case err != nil:
		return qty, err
	default:
		return qty, nil
	}
}

func (r *employeeRepo) UpdateEmployeeQuantityByProduct(ctx context.Context, product string, quantity int) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	// https://www.postgresql.org/docs/9.5/functions-json.html
	result, err := tx.Exec(ctx, `
		UPDATE public.employee
		SET	metadata['items']['qty'] = TO_JSONB($1::INT)
		WHERE	metadata['items']['product'] = TO_JSONB($2::TEXT)
	;`, quantity, product) // where case jsonb = jsonb (can't use ILIKE because of jsonb comparison)
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
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
