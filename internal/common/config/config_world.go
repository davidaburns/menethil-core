package config

type ConfigWorld struct {
	Network NetworkConfig  `ini:"server"`
	Database DatabaseConfig `ini:"database"`
}

func (c ConfigWorld) GetNetworkConfig() NetworkConfig {
	return c.Network
}

func (c ConfigWorld) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}
