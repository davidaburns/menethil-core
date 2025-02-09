package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type LoggerConfig struct {
	LevelFormat string `ini:"LogLevelFormat"`
	MessageFormat string `ini:"LogMessageFormat"`
	FieldNameFormat string `ini:"LogFieldNameFormat"`
	FieldValueFormat string `ini:"LogFieldValueFormat"`
}

func InitializeLogger(
	config *LoggerConfig,
) *zerolog.Logger {
	if config == nil {
		config = &LoggerConfig{
			LevelFormat: "| %-6s|",
			MessageFormat: "%s",
			FieldNameFormat: "%s:",
			FieldValueFormat: "%s",
		}
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i any) string {
		return strings.ToUpper(fmt.Sprintf(config.LevelFormat, i))
	}

	output.FormatMessage = func(i any) string {
		return fmt.Sprintf(config.MessageFormat, i)
	}

	output.FormatFieldName = func(i any) string {
		return fmt.Sprintf(config.FieldNameFormat, i)
	}

	output.FormatFieldValue = func(i any) string {
		return fmt.Sprintf(config.FieldValueFormat, i)
	}

	logger := zerolog.New(output).With().Timestamp().Logger()
	return &logger
}
