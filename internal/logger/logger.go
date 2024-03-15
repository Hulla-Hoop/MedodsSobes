package logger

import (
	formatter "github.com/fabienm/go-logrus-formatters"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	L *logrus.Logger
}

func New() *Logger {
	logger := logrus.New()
	rr := formatter.NewGelf("medos")
	logger.SetFormatter(rr)
	logger.SetLevel(logrus.DebugLevel)
	return &Logger{
		L: logger,
	}
}
