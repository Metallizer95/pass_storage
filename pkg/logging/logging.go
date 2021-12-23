package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func New(std bool, output io.Writer) {
	log := logrus.New()
	log.SetReportCaller(true)
	if std {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(ioutil.Discard)
	}
	log.AddHook(&writeHook{Writer: []io.Writer{output}, LogLevels: logrus.AllLevels})
	log.Formatter = &logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
	}
	log.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(log)
}

func GetLogger() (*Logger, error) {
	if e == nil {
		return nil, ErrLoggerNotInitialized
	}
	return &Logger{e}, nil
}
