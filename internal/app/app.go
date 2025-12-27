package app

import (
	"fmt"

	"taskboard/config"
	"taskboard/internal/application/usecase"
	"taskboard/internal/infrastructure/gateways/http"
	v1 "taskboard/internal/infrastructure/gateways/http/v1"
	handler "taskboard/internal/infrastructure/gateways/http/v1/handler/board"
	"taskboard/internal/infrastructure/persistence/postgres"
	client "taskboard/pkg/db/postgres"
	loggerPkg "taskboard/pkg/logger"
	"taskboard/pkg/validator"
)

func Run(cfg *config.Config) {
	log := loggerPkg.NewLogger(cfg)

	db, err := client.NewClient(cfg)
	if err != nil {
		log.Fatal("failed to connect to postgres", err)
	}
	boardRepo := postgres.NewBoardRepository(db.Pool)
	valid := validator.NewValidator()

	boardUC := usecase.NewCreateBoardUseCase(boardRepo)
	boardHandler := handler.NewBoardHandler(boardUC, log, valid)

	serv := http.NewServer(cfg, log)
	v1.NewRouter(serv.App(), boardHandler, log, valid)

	urlApp := fmt.Sprintf("%s:%d", serv.Config().Server.Host, serv.Config().Server.Port)
	if err := serv.App().Listen(urlApp); err != nil {
		log.Fatal("Error with init server", err)
	}
	// TODO: запустить gRPC-сервер приложения
}
