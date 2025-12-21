package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Priority string

const (
	LOW      Priority = "low"
	MEDIUM   Priority = "medium"
	HIGH     Priority = "high"
	CRITICAL Priority = "critical"
)

type Task struct {
	ID          int64       `json:"id"`
	ColumnID    int64       `json:"column_id"`
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
	Position    int64       `json:"position"`
	Priority    Priority    `json:"priority"`
	CreatedBy   int64       `json:"created_by"`
	AssignedTo  int64       `json:"assigned_to"`
	Deadline    time.Time   `json:"deadline"`
	IsCompleted bool        `json:"is_completed"`
	CompletedAt time.Time   `json:"completed_at"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
