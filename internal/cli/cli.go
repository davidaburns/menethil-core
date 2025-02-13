package cli

import (
	"flag"
	"fmt"

	"github.com/davidaburns/menethil-core/internal/errors"
	"github.com/davidaburns/menethil-core/internal/server"
)

type CliArguments struct {
	ConfigPath string
}

func ParseCliArguments(serverType server.ServerType) (*CliArguments, error) {
	serverStr := serverType.StringShort()
	if serverStr == "unkown" {
		return nil, fmt.Errorf("Error parsing cli arguments: %v", errors.ServerTypeUnkown)
	}

	defaultConfigPath := fmt.Sprintf("./conf/%v.config", serverStr)
	configPath := flag.String("config", defaultConfigPath, "Path to find server config file")

	flag.Parse()
	return &CliArguments{
		ConfigPath: *configPath,
	}, nil
}
