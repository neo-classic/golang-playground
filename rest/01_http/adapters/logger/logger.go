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

type Logger struct{}

var lgr zerolog.Logger

func (l Logger) Init(logMode Mode) {
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

	lgr = zerolog.New(logWriter).With().Timestamp().Logger()
}

func (l Logger) Debug(msg string) {
	lgr.Debug().Msg(msg)
}

func (l Logger) Info(msg string) {
	lgr.Info().Msg(msg)
}

func (l Logger) Error(msg string) {
	lgr.Error().Msg(msg)
}
