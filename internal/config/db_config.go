package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

// Dbconfig is a thin wrapper around pgx.ConnConfig
type Dbconfig pgx.ConnConfig

// DbPoolConfig wraps pgx.ConnPoolConfig to include pool settings
type DbPoolConfig struct {
	ConnConfig     Dbconfig
	MaxConnections int
}

func DbConfig() (DbPoolConfig, error) {
	portStr := os.Getenv("DB_PORT")
	if portStr == "" {
		return DbPoolConfig{}, fmt.Errorf("DB_PORT is not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return DbPoolConfig{}, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	maxConnStr := os.Getenv("DB_MAX_CONNECTIONS")
	if maxConnStr == "" {
		maxConnStr = "10" // sensible default
	}

	maxConnections, err := strconv.Atoi(maxConnStr)
	if err != nil {
		return DbPoolConfig{}, fmt.Errorf("invalid DB_MAX_CONNECTIONS: %w", err)
	}

	cfg := pgx.ConnConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     uint16(port),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	if cfg.Host == "" || cfg.User == "" || cfg.Database == "" {
		return DbPoolConfig{}, fmt.Errorf("missing required database environment variables")
	}

	return DbPoolConfig{
		ConnConfig:     Dbconfig(cfg),
		MaxConnections: maxConnections,
	}, nil
}
