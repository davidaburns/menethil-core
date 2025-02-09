package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/davidaburns/menethil-core/internal/common/config"
	"github.com/rs/zerolog"
)


func InitializeLogger(
	conf *config.LoggerConfig,
) *zerolog.Logger {
	if conf == nil {
		conf = &config.LoggerConfig{
			LevelFormat: "| %-6s|",
			MessageFormat: "%s",
			FieldNameFormat: "%s:",
			FieldValueFormat: "%s",
		}
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i any) string {
		return strings.ToUpper(fmt.Sprintf(conf.LevelFormat, i))
	}

	output.FormatMessage = func(i any) string {
		return fmt.Sprintf(conf.MessageFormat, i)
	}

	output.FormatFieldName = func(i any) string {
		return fmt.Sprintf(conf.FieldNameFormat, i)
	}

	output.FormatFieldValue = func(i any) string {
		return fmt.Sprintf(conf.FieldValueFormat, i)
	}

	logger := zerolog.New(output).With().Timestamp().Logger()
	return &logger
}
