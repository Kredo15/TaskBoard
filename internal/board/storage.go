package board

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, req CreateBoardDTO) (*Board, error)
	FindAll(ctx context.Context) (b []Board, err error)
	FindOne(ctx context.Context, id int) (Board, error)
	Update(ctx context.Context, req UpdateBoardDTO) error
	Delete(ctx context.Context, id int) error
}
