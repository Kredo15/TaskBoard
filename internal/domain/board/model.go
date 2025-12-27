package board

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Board struct {
	ID          int         `json:"id"`
	Title       string      `json:"title" validate:"required,max=100"`
	Description pgtype.Text `json:"description" validate:"max=500"`
	OwnerID     string      `json:"owner_id" validate:"required,uuid"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
