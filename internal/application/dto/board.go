package dto

import "github.com/jackc/pgx/v5/pgtype"

// CreateBoardRequest - запрос на создание доски
type CreateBoardRequest struct {
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
	OwnerID     string      `json:"owner_id"`
}

type CreateBoardResponse struct {
	ID int `json:"id"`
}

type UpdateBoardRequest struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
}
