package http

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	fiberLog "github.com/gofiber/fiber/v3/middleware/logger"

	"taskboard/config"
	loggerPkg "taskboard/pkg/logger"
)

type Server struct {
	app    *fiber.App
	cfg    *config.Config
	logger loggerPkg.Logger
}

func NewServer() (*Server, error) {
	// TODO: инициализировать объект конфига
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	// TODO: инициализировать логгер
	log := loggerPkg.NewLogger(cfg)

	// TODO: инициализировать приложение (app)
	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second * cfg.Server.ReadTimeout,
		WriteTimeout: time.Second * cfg.Server.WriteTimeout,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	app.Use(fiberLog.New(fiberLog.Config{
		Next:         nil,
		Done:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
	}))

	server := &Server{
		app:    app,
		cfg:    cfg,
		logger: log,
	}
	return server, nil
}

func (s Server) App() *fiber.App {
	return s.app
}

func (s Server) Config() *config.Config {
	return s.cfg
}

func (s Server) Logger() loggerPkg.Logger {
	return s.logger
}
