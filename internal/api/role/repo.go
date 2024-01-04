package role

import (
	"context"

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
