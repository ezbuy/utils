package echoutils

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var logger *log.Entry

func init() {
	l := &log.Logger{
		Handler: text.New(os.Stderr),
		Level:   log.InfoLevel,
	}

	logger = log.NewEntry(l)
}

func Logger() *log.Entry {
	return logger
}

func SetLogger(l *log.Entry) {
	if l != nil {
		logger = l
	}
}
