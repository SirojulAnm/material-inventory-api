package user

import "time"

type User struct {
	ID           int
	Email        string
	Name         string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
