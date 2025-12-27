package handler

import (
	"taskboard/internal/application/usecase"
	"taskboard/pkg/logger"
	"taskboard/pkg/validator"
)

type BoardHandler struct {
	createBoardUseCase usecase.CreateBoardUseCase
	log                logger.Logger
	validator          validator.Validator
}

func NewBoardHandler(createBoardUC usecase.CreateBoardUseCase, log logger.Logger, validator validator.Validator) *BoardHandler {
	return &BoardHandler{
		createBoardUseCase: createBoardUC,
		log:                log,
		validator:          validator,
	}
}
