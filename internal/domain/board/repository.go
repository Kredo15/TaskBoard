package board

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BoardRepository struct {
	db *pgxpool.Pool
}

func NewBoardRepository(db *pgxpool.Pool) *BoardRepository {
	return &BoardRepository{db: db}
}

func (r *BoardRepository) Create(ctx context.Context, req CreateBoardDTO) (*Board, error) {
	now := time.Now()
	board := &Board{
		Title:       req.Title,
		Description: req.Description,
		OwnerID:     req.OwnerID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	query := `
        INSERT INTO boards (title, description, owner_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
		RETURNIG ID
    `

	err := r.db.QueryRow(ctx, query,
		board.Title,
		board.Description,
		board.OwnerID,
		now,
		now,
	).Scan(&board.ID)

	if err != nil {
		return nil, fmt.Errorf("failed to create board: %w", err)
	}
	board.CreatedAt = now
	board.UpdatedAt = now
	return board, nil
}

func (r *BoardRepository) FindOne(ctx context.Context, id int) (Board, error) {
	var board Board

	query := `
		SELECT id, title, description, owner_id, created_at, updated_at
		FROM boards
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&board.ID,
		&board.Title,
		&board.Description,
		&board.CreatedAt,
		&board.UpdatedAt,
	)
	if err != nil {
		return Board{}, fmt.Errorf("failed to get board: %w", err)
	}
	return board, nil
}

func (r *BoardRepository) FindAll(ctx context.Context) ([]Board, error) {
	query := `
		SELECT id, title, description, owner_id, created_at, updated_at
		FROM boards
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	boards := make([]Board, 0)

	for rows.Next() {
		var b Board
		err = rows.Scan(&b.ID, &b.Title, &b.Description, &b.OwnerID, &b.CreatedAt, &b.UpdatedAt)
		if err != nil {
			return nil, err
		}
		boards = append(boards, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return boards, nil
}

func (r *BoardRepository) Update(ctx context.Context, req UpdateBoardDTO) error {
	query := `
        UPDATE boards
        SET title = $1, description = $2, updated_at = $3
        WHERE id = $4
    `

	now := time.Now()
	result, err := r.db.Exec(ctx, query,
		req.Title,
		req.Description,
		now,
		req.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update board: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("board not found")
	}

	return nil
}

func (r *BoardRepository) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM boards
		WHERE id = $1
	`
	result, err := r.db.Exec(ctx, query, id)

	if err != nil {
		return fmt.Errorf("failed to delete board: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("board not found")
	}
	return nil
}
