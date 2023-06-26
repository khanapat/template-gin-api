package database

import (
	"context"
	"fmt"
	"template-gin-api/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresConn(cfg *config.Config) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?idle_in_transaction_session_timeout=%d&sslmode=%s&pool_max_conns=%d",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
		cfg.Postgres.Timeout,
		cfg.Postgres.SSLMode,
		cfg.Postgres.PoolMaxConns,
	)
	dbConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
