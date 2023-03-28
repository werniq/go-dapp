package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var (
	logger = logrus.New()
)

func Logger() *log.Logger {
	return log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func InfoLogger(msg string) {
	logger.Out = os.Stdout
	logger.SetLevel(logrus.InfoLevel)

	logger.Info(msg)
}
