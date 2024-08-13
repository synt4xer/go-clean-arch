package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"time"
)

func ProvideLog() {
	logLevel := os.Getenv("APP_ENV")
	isDebug := os.Getenv("APP_DEBUG") == "true"

	var logger zerolog.Logger

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	switch logLevel {
	case "production":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	case "development", "local":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

		if isDebug {
			logger = zerolog.New(output).With().Timestamp().Caller().Logger()
		} else {
			logger = zerolog.New(output).With().Timestamp().Logger()
		}
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger = zerolog.New(output).With().Timestamp().Logger()
	}

	log.Logger = logger
}
