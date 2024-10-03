package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt *sql.NullTime
}
