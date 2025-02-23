package config

const (
	// Logger
	SECTION_LOGGER         ConfigSection = "logger"
	LOG_LEVEL_FORMAT       ConfigKey     = "LogLevelFormat"
	LOG_MESSAGE_FORMAT     ConfigKey     = "LogMessageFormat"
	LOG_FIELD_NAME_FORMAT  ConfigKey     = "LogFieldNameFormat"
	LOG_FIELD_VALUE_FORMAT ConfigKey     = "LogFieldValueFormat"

	// Network
	SECTION_NETWORK     ConfigSection = "network"
	NETWORK_LISTEN_HOST ConfigKey     = "ListenHost"
	NETWORK_LISTEN_PORT ConfigKey     = "ListenPort"

	// Database
	SECTION_DATABASE       ConfigSection = "database"
	DATABASE_DRIVER        ConfigKey     = "DatabaseDriver"
	DATABASE_DSN           ConfigKey     = "DatabaseDsn"
	DATABASE_MIGRATION_SRC ConfigKey     = "DatabaseMigrationSrc"
)
