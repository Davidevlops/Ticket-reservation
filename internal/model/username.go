package model

import (
	"time"
)

type Username struct {
	ID        int
	status    string // e.g., 1: active, 2: inactive, 3: banned
	value     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
