package v1

import (
	"github.com/gofiber/fiber/v3"

	handler "taskboard/internal/infrastructure/gateways/http/v1/handler/board"
	"taskboard/pkg/logger"
	"taskboard/pkg/validator"
)

func NewRouter(app *fiber.App, boardHandler *handler.BoardHandler, log logger.Logger, validator validator.Validator) {
	apiV1Group := app.Group("/api/v1")

	apiV1Group.Post("/board", boardHandler.CreateBoard)
}
