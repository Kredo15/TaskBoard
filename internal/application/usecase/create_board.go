package usecase

import (
	"context"
	"errors"
	"time"

	"taskboard/internal/application/dto"
	"taskboard/internal/domain/board"
)

type CreateBoardUseCase interface {
	Execute(cmd *dto.CreateBoardRequest) (*dto.CreateBoardResponse, error)
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
func (h *createBoardUseCase) Execute(cmd *dto.CreateBoardRequest) (*dto.CreateBoardResponse, error) {
	// Проверка валидности запроса
	if len(cmd.Title) == 0 || len(cmd.Description.String) > 255 {
		return nil, errors.New("некорректные данные для создания доски")
	}

	// Преобразование запроса в доменную модель
	newBoard := &board.Board{
		Title:       cmd.Title,
		Description: cmd.Description,
		OwnerID:     cmd.OwnerID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

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
