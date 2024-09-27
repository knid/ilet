package models

import (
	"errors"
	"time"
)


type Token struct {
    ID int
    Token string
    User User
    CreatedAt time.Time
    UpdatedAt time.Time
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
