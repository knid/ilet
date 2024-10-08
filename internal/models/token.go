package models

import (
	"database/sql"
	"errors"
	"time"
)

type Token struct {
	ID        int           `json:"-"`
	Token     string        `json:"token"`
	User      User          `json:"-"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *sql.NullTime `json:"updated_at"`
}

func (t Token) Validete() error {
	if t.IsExpired() {
		return errors.New("token is expired")
	}
	if len(t.Token) != 64 {
		return errors.New("invalid token")
	}

	return nil
}

func (t Token) IsExpired() bool {
	// TODO: Implement
	return false
}
