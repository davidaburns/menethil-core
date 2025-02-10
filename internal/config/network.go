package config

import "fmt"

type NetworkConfig struct {
	ListenHost string `ini:"ListenHost"`
	ListenPort int `ini:"ListenPort"`
}

func NetworkConfigFromConfig(c *Config) (NetworkConfig, error) {
	host, err := GetValue[string](c, SECTION_NETWORK, NETWORK_LISTEN_HOST)
	if err != nil {
		return NetworkConfig{}, err
	}

	port, err := GetValue[int](c, SECTION_NETWORK, NETWORK_LISTEN_PORT)
	if err != nil {
		return NetworkConfig{}, err
	}

	return NetworkConfig{ListenHost: host, ListenPort: port}, err
}

func (c NetworkConfig) GetAddress() string {
	return fmt.Sprintf("%v:%v", c.ListenHost, c.ListenPort)
}
