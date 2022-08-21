package logger

import (
	"github.com/sirupsen/logrus"
	"io"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
}

// InitLogger define standard loggers
func InitLogger(out io.Writer) {
	logrus.SetOutput(out)
}
