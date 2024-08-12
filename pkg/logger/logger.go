package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"time"
)

func ProvideLog() {
	logLevel := os.Getenv("LOG_LEVEL")

	var logger zerolog.Logger

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	switch logLevel {
	case "production":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	case "development":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	}

	log.Logger = logger
}
