package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/davidaburns/menethil-core/internal/cli"
	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/davidaburns/menethil-core/internal/logger"
	"github.com/davidaburns/menethil-core/internal/server"
)

func main() {
	args, err := cli.ParseCliArguments(server.ServerWorld)
	if err != nil {
		panic("Failed to parse command-line arguments")
	}

	conf, err := config.LoadConfig[config.ConfigWorld](args.ConfigPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load server config file: %v", args.ConfigPath))
	}

	logger := logger.InitializeLogger(&conf.Logger)
	s, err := server.NewServer(server.ServerWorld, conf, logger)
	if err != nil {
		logger.Fatal().Msgf("Failed to create server: %v", err)
	}

	go func() {
		s.Start()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	s.Stop()
}
