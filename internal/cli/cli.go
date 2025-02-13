package cli

import (
	"flag"
	"fmt"

	"github.com/davidaburns/menethil-core/internal/errors"
	"github.com/davidaburns/menethil-core/internal/server"
)

type CliArguments struct {
	ConfigPath string
	SqlPath string
}

func ParseCliArguments(serverType server.ServerType) (*CliArguments, error) {
	serverStr := serverType.StringShort()
	if serverStr == "unkown" {
		return nil, fmt.Errorf("Error parsing cli arguments: %v", errors.ServerTypeUnkown)
	}

	defaultConfigPath := fmt.Sprintf("./conf/%v.config", serverStr)
	defaultSqlPath := fmt.Sprintf("../sql/%v", serverStr)

	configPath := flag.String("config", defaultConfigPath, "Path to find server config file")
	sqlPath := flag.String("sql", defaultSqlPath, "Path to find sql files for database migrations")

	flag.Parse()
	return &CliArguments{
		ConfigPath: *configPath,
		SqlPath: *sqlPath,
	}, nil
}
