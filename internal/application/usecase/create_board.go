package usecase

import (
	"context"

	"taskboard/internal/application/dto"
	"taskboard/internal/domain/board"
)

type CreateBoardUseCase interface {
	Execute(ctx context.Context, cmd *dto.CreateBoardRequest) (*dto.CreateBoardResponse, error)
}

// createBoardHandler представляет обработчик команды создания доски
type createBoardUseCase struct {
	boardRepo board.Repository
}

// NewCreateBoardHandler создает новый экземпляр обработчика команды создания доски
func NewCreateBoardUseCase(repo board.Repository) CreateBoardUseCase {
	return &createBoardUseCase{
		boardRepo: repo,
	}
}

// Execute обрабатывает команду создания доски
func (h *createBoardUseCase) Execute(ctx context.Context, cmd *dto.CreateBoardRequest) (*dto.CreateBoardResponse, error) {
	// Преобразование запроса в доменную модель
	newBoard := &board.Board{
		Title:       cmd.Title,
		Description: cmd.Description,
		OwnerID:     cmd.OwnerID,
	}

	// Сохранение доски в репозитории
	err := h.boardRepo.Create(ctx, newBoard)
	if err != nil {
		return nil, err
	}

	// Возвращаем успешный ответ
	response := &dto.CreateBoardResponse{
		ID: newBoard.ID,
	}

	return response, nil
}
