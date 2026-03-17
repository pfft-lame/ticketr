package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = 100                      // Maximum number of connections
	cfg.MinConns = 20                       // Minimum number of connections to maintain
	cfg.MaxConnLifetime = time.Hour         // Connection lifetime
	cfg.MaxConnIdleTime = 30 * time.Minute  // Idle connection timeout
	cfg.HealthCheckPeriod = 5 * time.Minute // Health check interval

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return pool, pool.Ping(ctx)
}
