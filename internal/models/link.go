package models

import "time"

type Link struct {
	ID        int
	User      User
	Short     string
	Long      string
	Active    bool
	Visited   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
