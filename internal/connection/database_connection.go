package connection

import (
	"fmt"

	"go-backend-app/internal/config"

	"github.com/jackc/pgx"
)

func DatabaseConnection() (*pgx.Conn, error) {
	cfg, err := config.DbConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get database configuration: %w", err)
	}

	// Explicit conversion back to pgx.ConnConfig
	conn, err := pgx.Connect(pgx.ConnConfig(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return conn, nil
}
