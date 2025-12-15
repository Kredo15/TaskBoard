package main

import (
	"fmt"
	"os"

	middleware "github.com/gofiber/contrib/v3/zerolog"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"

	"taskboard/config"
)

func main() {
	// TODO: инициализировать объект конфига
	cfg := config.MustLoad()
	// TODO: инициализировать логгер
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	// TODO: инициализировать приложение (app)
	app := fiber.New()
	app.Use(middleware.New(middleware.Config{
		Logger: &logger,
	}))
	urlApp := fmt.Sprintf("%s:%d", cfg.Serverv.Host, cfg.Serverv.Port)
	if err := app.Listen(urlApp); err != nil {
		logger.Fatal().Err(err).Msg("Fiber app error")
	}
	// TODO: запустить gRPC-сервер приложения
}
