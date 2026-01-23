package model

import (
	"time"
)

type Email struct {
	ID        int
	status    string // e.g., 1: active, 2: inactive, 3: banned
	value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *Email) Value() string {
	return e.value
}
func (e *Email) Status() string {
	return e.status
}
