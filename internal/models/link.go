package models

import (
	"database/sql"
	"time"
)

type Link struct {
	ID        int
	User      User
	Short     string
	Long      string
	Active    bool
	Visited   int
	CreatedAt time.Time
	UpdatedAt *sql.NullTime
}
