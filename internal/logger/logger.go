package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func InitializeLogger(level zerolog.Level) *zerolog.Logger {
	zerolog.SetGlobalLevel(level)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i any) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	output.FormatMessage = func(i any) string {
		return fmt.Sprintf("%s", i)
	}

	output.FormatFieldName = func(i any) string {
		return fmt.Sprintf("%s:", i)
	}

	output.FormatFieldValue = func(i any) string {
		return fmt.Sprintf("%s", i)
	}

	logger := zerolog.New(output).With().Caller().Timestamp().Logger()
	return &logger
}
