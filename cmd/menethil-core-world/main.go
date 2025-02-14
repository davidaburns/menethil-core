package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/davidaburns/menethil-core/build"
	"github.com/davidaburns/menethil-core/internal/cli"
	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/davidaburns/menethil-core/internal/logger"
	"github.com/davidaburns/menethil-core/internal/server"
	"github.com/rs/zerolog"
)

func main() {
	serverType := server.ServerWorld
	log := logger.InitializeLogger(zerolog.DebugLevel)
	build := build.NewBuildInfo()

	log.Info().Msgf("Menethil Core: %v Server", serverType.String())
	log.Info().Msgf("Version: %v", build.String())

	args, err := cli.ParseCliArguments(server.ServerAuth)
	if err != nil {
		log.Fatal().Msg("Failed to parse command-line arguments")
	}

	conf, err := config.LoadConfig(args.ConfigPath)
	if err != nil {
		log.Fatal().Msgf("Failed to load server config file: %v", args.ConfigPath)
	}

	serv := server.NewWorldServer(log, conf)
	if serv == nil {
		log.Fatal().Msg("Failed to create instance of server")
	}

	bootstrap, err := server.NewBootstrapperWithServer(serv, conf, log)
	if err != nil {
		log.Fatal().Msgf("Failed to create server: %v", err)
	}

	go func() {
		log.Info().Msg("Starting server")
		if err := bootstrap.Start(); err != nil {
			log.Fatal().Msgf("Failed to start server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	bootstrap.Stop()
}
