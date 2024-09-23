package models

import "time"

type User struct {
	ID        int
	Username  string
	Password  string
	Links     []Link
	CreatedAt time.Time
	UpdatedAt time.Time
}
