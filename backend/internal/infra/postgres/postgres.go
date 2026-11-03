package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"FlashHelper/internal/config"
)

func NewDB(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, func(), error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, func() {}, fmt.Errorf("parse postgres config: %w", err)
	}

	poolConfig.MinConns = 2
	poolConfig.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, func() {}, fmt.Errorf("create postgres pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, func() {}, fmt.Errorf("ping postgres: %w", err)
	}

	cleanup := func() { pool.Close() }

	return pool, cleanup, nil
}