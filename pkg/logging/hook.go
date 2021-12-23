package logging

import (
	"github.com/sirupsen/logrus"
	"io"
)

type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writeHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return nil
	}
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
	}
	return err
}

func (hook *writeHook) Levels() []logrus.Level {
	return hook.LogLevels
}
