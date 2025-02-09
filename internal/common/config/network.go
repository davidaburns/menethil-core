package config

import "fmt"

type NetworkConfig struct {
	ListenHost string `ini:"ListenHost"`
	ListenPort int `ini:"ListenPort"`
}

func (c NetworkConfig) GetAddress() string {
	return fmt.Sprintf("%v:%v", c.ListenHost, c.ListenPort)
}
