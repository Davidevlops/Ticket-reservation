package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

// Dbconfig is a thin wrapper around pgx.ConnConfig
type Dbconfig pgx.ConnConfig

func DbConfig() (Dbconfig, error) {
	portStr := os.Getenv("DB_PORT")
	if portStr == "" {
		return Dbconfig{}, fmt.Errorf("DB_PORT is not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return Dbconfig{}, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	cfg := pgx.ConnConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     uint16(port),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	if cfg.Host == "" || cfg.User == "" || cfg.Database == "" {
		return Dbconfig{}, fmt.Errorf("missing required database environment variables")
	}

	return Dbconfig(cfg), nil
}
