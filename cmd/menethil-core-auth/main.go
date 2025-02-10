package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/davidaburns/menethil-core/build"
	"github.com/davidaburns/menethil-core/internal/auth"
	"github.com/davidaburns/menethil-core/internal/cli"
	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/davidaburns/menethil-core/internal/logger"
	"github.com/davidaburns/menethil-core/internal/server"
)

func main() {
	serverType := server.ServerAuth
	logger := logger.InitializeLogger()
	build := build.NewBuildInfo()

	logger.Info().Msgf("Menethil Core: %v Server", serverType.String())
	logger.Info().Msgf("Version: %v", build.String())

	args, err := cli.ParseCliArguments(server.ServerAuth)
	if err != nil {
		logger.Fatal().Msg("Failed to parse command-line arguments")
	}

	conf, err := config.LoadConfig(args.ConfigPath)
	if err != nil {
		logger.Fatal().Msgf("Failed to load server config file: %v: %v", args.ConfigPath, err)
	}

	serv := auth.NewAuthServer(logger, conf)
	bootstrap, err := server.NewBootstrapperWithServer(serv, conf, logger)
	if err != nil {
		logger.Fatal().Msgf("Failed to bootstrap server: %v", err)
	}

	go func() {
		logger.Info().Msg("Starting server")
		if err := bootstrap.Start(); err != nil {
			logger.Fatal().Msgf("Failed to start server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	bootstrap.Stop()
}
