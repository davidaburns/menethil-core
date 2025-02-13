package database

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

type DatabaseLogger struct {
	Log *zerolog.Logger
}

func NewDatabaseLogger(log *zerolog.Logger) *DatabaseLogger {
	return &DatabaseLogger{Log: log}
}

func (dbl *DatabaseLogger) Printf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	msg = strings.TrimSuffix(msg, "\n")
	dbl.Log.Print(msg)
}

func (dbl *DatabaseLogger) Verbose() bool {
	return true
}
