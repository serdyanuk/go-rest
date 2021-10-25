package logger

import (
	"errors"
	"io"
	"os"

	"github.com/serdyanuk/go-rest/config"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var defaultLogger = newLogger(config.Get())

func newLogger(conifg *config.Config) *Logger {
	var log = logrus.New()

	err := os.Mkdir("logs", 0766)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}
	file, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	if conifg.IsProduction() {
		log.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		log.SetOutput(os.Stdout)
	}

	return &Logger{
		Logger: log,
	}
}

func Get() *Logger {
	return defaultLogger
}
