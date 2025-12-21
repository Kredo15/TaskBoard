package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"taskboard/config"
)

type Logger struct {
	cfg    *config.Config
	logger zerolog.Logger
}

func NewLogger(cfg *config.Config) Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	switch cfg.Logging.Level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return Logger{
		cfg:    cfg,
		logger: logger,
	}
}

func (l Logger) Debug(msg string) {
	log.Debug().Msg(msg)
}

func (l Logger) Info(msg string) {
	log.Info().Msg(msg)
}

func (l Logger) Warn(msg string) {
	log.Warn().Msg(msg)
}

func (l Logger) Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}

func (l Logger) Fatal(msg string, err error) {
	log.Fatal().Err(err).Msg(msg)
}
