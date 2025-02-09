package config

type ConfigAuth struct {
	Network NetworkConfig  `ini:"server"`
	Database DatabaseConfig `ini:"database"`
	Logger LoggerConfig `ini:"logger"`
}

func (c ConfigAuth) GetNetworkConfig() NetworkConfig {
	return c.Network
}

func (c ConfigAuth) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c ConfigAuth) GetLoggerConfig() LoggerConfig {
	return c.Logger
}
