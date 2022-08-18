package internal

import (
	"github.com/sirupsen/logrus"
	"io"
)

var Logger *logrus.Logger

// InitLogger define standard loggers
func InitLogger(out io.Writer) {
	Logger = logrus.New()
	logrus.SetOutput(out)
}
