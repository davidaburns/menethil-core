package config

type ConfigAuth struct {
	Network NetworkConfig  `ini:"server"`
	Database DatabaseConfig `ini:"database"`
}

func (c ConfigAuth) GetNetworkConfig() NetworkConfig {
	return c.Network
}

func (c ConfigAuth) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}
