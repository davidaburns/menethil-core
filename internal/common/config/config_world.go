package config

type ConfigWorld struct {
	Network NetworkConfig  `ini:"server"`
	Database DatabaseConfig `ini:"database"`
	Logger LoggerConfig `ini:"logger"`
}

func (c ConfigWorld) GetNetworkConfig() NetworkConfig {
	return c.Network
}

func (c ConfigWorld) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c ConfigWorld) GetLoggerConfig() LoggerConfig {
	return c.Logger
}
