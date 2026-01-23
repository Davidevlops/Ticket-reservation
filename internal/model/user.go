package model

import (
	"time"
)

type User struct {
	ID          int
	status      string // e.g., 1: active, 2: inactive, 3: banned
	Username    Username
	PhoneNumber PhoneNumber
	Email       Email
	Password    Password
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) Status() string {
	return u.status
}
func (u *User) IDValue() int {
	return u.ID
}
