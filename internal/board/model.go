package board

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Board struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
	OwnerID     int         `json:"owner_id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
