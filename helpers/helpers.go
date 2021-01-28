package helpers

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"io"
	"os"
)

func InitLog() {
	logrus.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "02.01.2006 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(io.MultiWriter(os.Stdout))
}
