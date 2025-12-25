package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v3"

	"taskboard/internal/application/dto"
	"taskboard/internal/application/usecase"
	"taskboard/pkg/logger"
	"taskboard/pkg/validator"
)

type BoardHandler struct {
	createBoardUseCase usecase.CreateBoardUseCase
	log                logger.Logger
	validator          validator.Validator
}

func NewBoardHandler(boardSvc *usecase.CreateBoardUseCase, log *logger.Logger, validator validator.Validator) *BoardHandler {
	return &BoardHandler{
		createBoardUseCase: *boardSvc,
		log:                *log,
		validator:          validator,
	}
}

func (b *BoardHandler) CreateBoard(c fiber.Ctx) error {
	req := dto.CreateBoardRequest{}

	if err := c.Bind().Body(&req); err != nil {
		b.log.Warn("failed to parse request body: %s", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}
	board, err := b.createBoardUseCase.Execute(&req)

	if err != nil {
		b.log.Error("create board use case failed", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create board",
		})
	}

	if err := b.validator.Validate(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error validate": err,
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": board,
	})
}
