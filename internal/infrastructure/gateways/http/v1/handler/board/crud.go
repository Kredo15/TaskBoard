package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v3"

	"taskboard/internal/application/dto"
)

func (b *BoardHandler) CreateBoard(c fiber.Ctx) error {
	var req dto.CreateBoardRequest

	if err := c.Bind().Body(&req); err != nil {
		b.log.Warn("failed to parse request body: %s", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}

	if err := b.validator.Validate(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error validate": err,
		})
	}

	board, err := b.createBoardUseCase.Execute(c.Context(), &req)

	if err != nil {
		b.log.Error("create board use case failed", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create board",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": board,
	})
}

func (b *BoardHandler) UpdateBoard(c fiber.Ctx) error {
	var req dto.UpdateBoardRequest

	if err := c.Bind().Body(&req); err != nil {
		b.log.Warn("failed to parse request body: %s", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}
	return nil
}
