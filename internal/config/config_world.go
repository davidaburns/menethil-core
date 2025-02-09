package config

import "github.com/davidaburns/menethil-core/internal/logger"

type ConfigWorld struct {
	Server struct {
		ListenHost string `ini:"ListenHost"`
		ListenPort int `ini:"ListenPort"`
	} `ini:"server"`
	Database struct {
		Driver string `ini:"DatabaseDriver"`
		Host string `ini:"DatabaseHost"`
		Port int `ini:"DatabasePort"`
		Username string `ini:"DatabaseUsername"`
		Password string `ini:"DatabasePassword"`
		WorldDatabase string `ini:"DatabaseWorldDatabase"`
	} `ini:"database"`

	Logger logger.LoggerConfig `ini:"logger"`
}
