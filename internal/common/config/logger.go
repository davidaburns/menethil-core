package config

type LoggerConfig struct {
	LevelFormat string `ini:"LogLevelFormat"`
	MessageFormat string `ini:"LogMessageFormat"`
	FieldNameFormat string `ini:"LogFieldNameFormat"`
	FieldValueFormat string `ini:"LogFieldValueFormat"`
}
