package config

import (
	"gopkg.in/ini.v1"
)

type ServerConfig interface {
	ConfigAuth | ConfigWorld
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
