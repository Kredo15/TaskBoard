package board

import "github.com/jackc/pgx/v5/pgtype"

// CreateBoardRequest - запрос на создание доски
type CreateBoardDTO struct {
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
	OwnerID     int         `json:"owner_id"`
}

type UpdateBoardDTO struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
}
