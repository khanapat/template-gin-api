package role

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type roleRepo struct {
	db *pgxpool.Pool
}

func NewRoleRepo(db *pgxpool.Pool) *roleRepo {
	return &roleRepo{
		db: db,
	}
}

func (r *roleRepo) CreateRole(ctx context.Context, title string, desc string) (int, error) {
	var id int

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := tx.QueryRow(ctx, `
		INSERT INTO public."role"
		(
			title,
			description
		)
		VALUES
		(
			$1,
			$2
		)
		RETURNING "id"
	;`, title, desc).Scan(&id); err != nil {
		return 0, err
	}
	if err := tx.Commit(ctx); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *roleRepo) GetUnixCurrentTimestamp(ctx context.Context) (int64, error) {
	var ts int64

	if err := r.db.QueryRow(ctx, `SELECT extract(epoch from now())::int;`).Scan(&ts); err != nil {
		return 0, err
	}

	return ts, nil
}

func (r *roleRepo) GetCurrentTimestamp(ctx context.Context) (time.Time, error) {
	var ts time.Time

	if err := r.db.QueryRow(ctx, `SELECT current_timestamp;`).Scan(&ts); err != nil {
		return time.Time{}, err
	}

	return ts, nil
}
