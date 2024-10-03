package models

import (
	"database/sql"
	"time"
)

type Link struct {
	ID        int           `json:"id"`
	User      User          `json:"-"`
	Short     string        `json:"short"`
	Long      string        `json:"long"`
	Active    bool          `json:"active"`
	Visited   int           `json:"visited"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *sql.NullTime `json:"updated_at"`
}
