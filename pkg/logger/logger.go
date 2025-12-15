package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger() Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return Logger{
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
