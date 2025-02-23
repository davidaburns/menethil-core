package server

import (
	"fmt"
	"net"

	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/davidaburns/menethil-core/internal/process"
	"github.com/rs/zerolog"
)

type ServerBootstrap struct {
	Log     *zerolog.Logger
	Config  *config.Config
	Server  Server
	Process *process.SystemProcess
}

func NewBootstrapperWithServer(server Server, conf *config.Config, log *zerolog.Logger) (*ServerBootstrap, error) {
	proc, err := process.GetSystemProcess()
	if err != nil {
		return nil, err
	}

	return &ServerBootstrap{
		Log:     log,
		Config:  conf,
		Server:  server,
		Process: proc,
	}, nil
}

func (s *ServerBootstrap) Start() error {
	s.Process.CreatePidFile()
	s.Process.SetProcessPriority(process.PROCESS_PRIORITY_HIGH)
	prio, _ := s.Process.GetProcessPriority()

	s.Log.Info().Msgf("Process PID: %v", s.Process.Pid)
	s.Log.Info().Msgf("Process Priority: %v", prio)

	err := s.Server.Start()
	if err != nil {
		return err
	}

	networkConfig, err := config.NetworkConfigFromConfig(s.Config)
	if err != nil {
		return err
	}

	ln, err := net.Listen("tcp", fmt.Sprintf("%v", networkConfig.GetAddress()))
	if err != nil {
		return err
	}
	defer ln.Close()

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
