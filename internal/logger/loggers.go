package logger

import (
	"fmt"
	"io"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
}

// InitLogger define standard loggers
func InitLogger(out io.Writer) {
	logrus.SetOutput(out)

	Logger.SetReportCaller(true)
	Logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

}
