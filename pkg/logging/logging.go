package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
)

type Logger interface {
	Debug(msg interface{}, args ...interface{})
	Warn(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Error(msg interface{}, args ...interface{})
	Fatal(msg interface{}, args ...interface{})
}

type logger struct {
	logger *zerolog.Logger
}

var glLogger = logger{}

func New(level Level, filepath interface{}) (Logger, error) {
	var output io.Writer
	if filepath == nil {
		output = os.Stdout
	} else {
		var err error
		output, err = os.OpenFile(filepath.(string), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0665)
		if err != nil {
			return nil, err
		}
	}

	log := zerolog.New(output)
	var loglevel zerolog.Level

	switch level {
	case DEBUG:
		loglevel = zerolog.DebugLevel
	case WARNING:
		loglevel = zerolog.WarnLevel
	case ERROR:
		loglevel = zerolog.ErrorLevel
	default:
		loglevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(loglevel)
	glLogger.logger = &log
	return &glLogger, nil
}

func Get() (Logger, error) {
	if glLogger.logger == nil {
		return nil, ErrLoggerNotInitialized
	}
	return &glLogger, nil
}

func (l *logger) Debug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

func (l *logger) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

func (l *logger) Warn(message string, args ...interface{}) {
	l.log(message, args...)
}

func (l *logger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}
	l.msg("error", message, args...)
}

func (l *logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)
}

func (l *logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		l.logger.Info().Msgf(message, args...)
	}
}

func (l *logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
