package connection

import (
	"fmt"

	"go-backend-app/internal/config"

	"github.com/jackc/pgx"
)

func DatabasePoolConnection() (*pgx.ConnPool, error) {
	cfg, err := config.DbConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get database configuration: %w", err)
	}

	poolCfg := pgx.ConnPoolConfig{
		ConnConfig:     pgx.ConnConfig(cfg.ConnConfig), // explicit conversion
		MaxConnections: cfg.MaxConnections,
	}

	pool, err := pgx.NewConnPool(poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	return pool, nil
}
