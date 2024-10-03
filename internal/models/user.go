package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int           `json:"-"`
	Username  string        `json:"username"`
	Password  string        `json:"-"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *sql.NullTime `json:"updated_at"`
}
