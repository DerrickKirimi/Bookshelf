package logger

import (
	"os"
	"io"
	"github.com/rs/zerolog"
	"context"
	"fmt"
)

type Logger struct {
	logger *zerolog.Logger
}

func New(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.TraceLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{logger : &logger}
}

func NewConsole(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.TraceLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &Logger{logger : &logger}
}

func (l *Logger) Output (w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}

 func (l *Logger) With() zerolog.Context {
	return l.logger.With()
 }

 func (l *Logger) Sample (s zerolog.Sampler) zerolog.Logger {
	return l.logger.Sample(s)
 }

 func (l *Logger) Hook(h zerolog.Hook) zerolog.Logger {
	return l.logger.Hook(h)
 }

//Start a new message with error level with err as a fielf if not nil or
//with info level if err is nil.
//To send the event call message on the returned event

func (l *Logger) Err(err error) *zerolog.Event {
	return l.logger.Err(err)
}

//Start a new message with trace level
//To send event call Msg on returned event

func(l *Logger) Trace() *zerolog.Event {
	return l.logger.Trace()
}


func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}


func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}

func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}

func (l *Logger) WithLevel(level zerolog.Level) *zerolog.Event {
	return l.logger.WithLevel(level)	
}

func (l *Logger) Print (v ... interface{}) {
	l.logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

func (l *Logger) Printf (format string, v ...interface{}) {
	l.logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

func (l *Logger) Ctx (ctx context.Context) *Logger {
	return &Logger{logger: zerolog.Ctx(ctx)}
}