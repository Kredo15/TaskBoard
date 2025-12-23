package board

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, board *Board) error
	FindAll(ctx context.Context) (b []Board, err error)
	FindOne(ctx context.Context, id int) (Board, error)
	Update(ctx context.Context, board Board) error
	Delete(ctx context.Context, id int) error
}
