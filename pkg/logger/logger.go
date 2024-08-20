package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"go.uber.org/fx/fxevent"
	"os"
	"strings"
	"time"
)

type logger struct {
	Logger zerolog.Logger
}

func (l *logger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Info().
			Str("func", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStart Hook executing")
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.Warn().Err(e.Err).
				Str("func", e.FunctionName).
				Str("caller", e.CallerName).
				Msg("OnStart hook failed")
		} else {
			l.Logger.Info().
				Str("func", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStart hook executed")
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Info().Str("func", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStop hook executing")
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.Warn().Err(e.Err).
				Str("func", e.FunctionName).
				Str("caller", e.CallerName).
				Msg("OnStop hook failed")
		} else {
			l.Logger.Info().Str("func", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStop hook executed")
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.Logger.Warn().Err(e.Err).
				Str("type", e.TypeName).
				Msg("supplied")
		} else {
			l.Logger.Info().
				Str("type", e.TypeName).
				Msg("supplied")
		}
	case *fxevent.Provided:
		for _, rangeType := range e.OutputTypeNames {
			l.Logger.Info().Str("type", rangeType).
				Str("constructor", e.ConstructorName).
				Msg("provided")
		}
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).
				Msg("error encountered while applying options")
		}
	case *fxevent.Invoking:
		// Do nothing. Will log on Invoked.
	case *fxevent.Invoked:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Str("stacktrace", e.Trace).
				Str("function", e.FunctionName).Msg("invoke failed")
		} else {
			l.Logger.Info().Str("function", e.FunctionName).Msg("invoked")
		}
	case *fxevent.Stopping:
		l.Logger.Info().Str("signal", strings.ToUpper(e.Signal.String())).Msg("received signal")
	case *fxevent.Stopped:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("stop failed")
		}
	case *fxevent.RollingBack:
		l.Logger.Error().Err(e.StartErr).Msg("start failed, rolling back")
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("rollback failed")
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("start failed")
		} else {
			l.Logger.Info().Msg("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.Logger.Error().Err(e.Err).Msg("custom logger initialization failed")
		} else {
			l.Logger.Info().Str("function", e.ConstructorName).Msg("initialized custom fxevent.Logger")
		}
	}
}

func Default() func() fxevent.Logger {
	var zeroLogger zerolog.Logger
	logLevel := os.Getenv("APP_ENV")
	isDebug := os.Getenv("APP_DEBUG") == "true"

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	switch logLevel {
	case "production":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zeroLogger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	case "development", "local":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

		if isDebug {
			zeroLogger = zerolog.New(output).With().Timestamp().Caller().Logger()
		} else {
			zeroLogger = zerolog.New(output).With().Timestamp().Logger()
		}
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		zeroLogger = zerolog.New(output).With().Timestamp().Logger()
	}

	log.Logger = zeroLogger

	return WithZerolog(zeroLogger)
}

// WithZerolog customize zerolog instance for fxevent.
func WithZerolog(l zerolog.Logger) func() fxevent.Logger {
	return func() fxevent.Logger {
		return &logger{Logger: l}
	}
}
