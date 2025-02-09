package config

import (
	"gopkg.in/ini.v1"
)

type ServerConfig interface {
	GetNetworkConfig() NetworkConfig
	GetDatabaseConfig() DatabaseConfig
	GetLoggerConfig() LoggerConfig
}

func LoadConfig[T ServerConfig](path string) (*T, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	var config T
	err = cfg.MapTo(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
