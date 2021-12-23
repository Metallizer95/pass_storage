package logging

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Logger interface {
	Debug(msg interface{}, args ...interface{})
	Warn(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Error(msg error, args ...interface{})
	Fatal(msg error, args ...interface{})
}

type logger struct {
	logger *logrus.Entry
}

var glLogger = logger{}

func New(level Level, output []io.Writer) (Logger, error) {
	log := logrus.New()
	log.SetReportCaller(false)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableColors:    false,
		FullTimestamp:    true,
		CallerPrettyfier: nil,
	})
	log.SetLevel(logrus.InfoLevel)
	glLogger.logger = logrus.NewEntry(log)
	return &glLogger, nil
}

func Get() (Logger, error) {
	if glLogger.logger == nil {
		return nil, ErrLoggerNotInitialized
	}
	return &glLogger, nil
}

func (l *logger) Debug(message interface{}, args ...interface{}) {
	//l.msg(DEBUG, message, args...)
}

func (l *logger) Info(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info(message)
	} else {
		l.logger.Infof(message, args...)
	}
}

func (l *logger) Warn(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Warn(message)
	} else {
		l.logger.Warnf(message, args...)
	}
}

func (l *logger) Error(message error, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Error(message.Error())
	} else {
		l.logger.Errorf(message.Error(), args...)
	}
}

func (l *logger) Fatal(message error, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Fatal(message.Error())
	} else {
		l.logger.Fatalf(message.Error(), args...)
	}
}
