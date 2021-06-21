package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Mode string

const (
	Prod Mode = "prod"
	Dev  Mode = "dev"
)

type Logger struct {
	log zerolog.Logger
}

func NewLogger(logMode Mode) *Logger {
	var logger Logger
	logger.Init(logMode)
	return &logger
}

func (l *Logger) Init(logMode Mode) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.DurationFieldInteger = true

	var logWriter io.Writer
	switch logMode {
	case Prod:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		logWriter = io.Writer(os.Stdout)
		break
	case Dev:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logWriter = io.Writer(zerolog.ConsoleWriter{Out: os.Stdout})
		break
	}

	l.log = zerolog.New(logWriter).With().Timestamp().Logger()
}

func (l *Logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *Logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *Logger) Error(msg string) {
	l.log.Error().Msg(msg)
}
