package models

import (
	"time"
)

type Column struct {
	ID        int64     `json:"id"`
	BoardID   int64     `json:"board_id"`
	Title     string    `json:"title"`
	Position  int64     `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
