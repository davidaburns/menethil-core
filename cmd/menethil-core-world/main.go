package main

import (
	"fmt"

	"github.com/davidaburns/menethil-core/internal/cli"
	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/davidaburns/menethil-core/internal/logger"
	"github.com/davidaburns/menethil-core/internal/server"
)

func main() {
	args, err := cli.ParseCliArguments(server.ServerAuth)
	if err != nil {
		panic("Failed to parse command-line arguments")
	}

	conf, err := config.LoadConfig[config.ConfigAuth](args.ConfigPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load server config file: %v", args.ConfigPath))
	}

	logger := logger.InitializeLogger(&conf.Logger)
	logger.Info().Msg("Menethil Core: World Server")
}
