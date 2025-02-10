package server

import (
	"fmt"
	"net"

	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/davidaburns/menethil-core/internal/process"
	"github.com/rs/zerolog"
)

/*
RESPONSIBILITIES:
- Setup Logging
- Setup Process Signals
- Load Configs?: DAVE: This seems like it should be the responsibilities of the actual server not the bootstrapper
- Start Listeners?: In order for this to happen we need to load configs
- Setup Connection Handlers?
- Create PID File
- Set Process Afinitiy/Priority
*/
type ServerBootstrap struct {
	Log *zerolog.Logger
	Config *config.Config
	Server Server
}

func NewBootstrapperWithServer(server Server, conf *config.Config, log *zerolog.Logger) (*ServerBootstrap, error) {
	return &ServerBootstrap{
		Log: log,
		Config: conf,
		Server: server,
	}, nil
}

func (s *ServerBootstrap) Start() error {
	s.Server.Start()
	networkConfig, err := config.NetworkConfigFromConfig(s.Config)
	if err != nil {
		return err
	}

	ln, err := net.Listen("tcp", fmt.Sprintf("%v", networkConfig.GetAddress()))
	if err != nil {
		return err
	}
	defer ln.Close()

	process.CreatePIDFile()
	s.Log.Info().Msgf("Listening on: %v", networkConfig.GetAddress())
	for {
		conn, err := ln.Accept()
		if err != nil {
			s.Log.Error().Msgf("Error accepting client connection: %v", err)
		}

		s.Server.HandleConnection(conn)
	}
}

func (s *ServerBootstrap) Stop() {
	s.Log.Info().Msg("Shutting down server")
	s.Server.Stop()
}
