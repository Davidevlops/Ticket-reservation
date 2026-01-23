package model

import (
	"time"
)

type PasswordCharacteristics struct {
	Length         int
	UppercaseCount int
	NumericCount   int
	SymbolCount    int
	Score          int
}

type Password struct {
	ID                      int
	value                   string
	status                  string // e.g., 1: active, 2: inactive, 3: banned
	passwordScore           string
	PasswordCharacteristics PasswordCharacteristics
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
