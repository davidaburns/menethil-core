package server

import (
	"fmt"
	"net"

	"github.com/davidaburns/menethil-core/build"
	"github.com/davidaburns/menethil-core/internal/common/client"
	"github.com/davidaburns/menethil-core/internal/common/config"
	"github.com/rs/zerolog"
)

type Server struct {
	Build build.BuildInfo
	Log *zerolog.Logger
	Config config.ServerConfig
	ServerType ServerType
	Clients map[string]*client.Client
}

func NewServer(serverType ServerType, conf config.ServerConfig, log *zerolog.Logger) (*Server, error) {
	return &Server{
		Build: *build.NewBuildInfo(),
		Log: log,
		Config: conf,
		ServerType: serverType,
		Clients: make(map[string]*client.Client),
	}, nil
}

func (s *Server) Start() error {
	s.Log.Info().Msgf("Menethil Core: %v Server", s.ServerType.String())
	s.Log.Info().Msgf("Version: %v", s.Build.String())
	s.Log.Info().Msg("Starting server")

	conf := s.Config.GetNetworkConfig()
	ln, err := net.Listen("tcp", fmt.Sprintf("%v", conf.GetAddress()))
	if err != nil {
		return err
	}
	defer ln.Close()

	s.Log.Info().Msgf("Listening on: %v", conf.GetAddress())
	for {
		conn, err := ln.Accept()
		if err != nil {
			s.Log.Error().Msgf("Error accepting client connection: %v", err)
		}

		s.Log.Info().Msgf("Client Local Network: %v", conn.LocalAddr().String())
		s.Log.Info().Msgf("Client Remote Network: %v", conn.RemoteAddr().String())

		client := client.FromConnection(conn)
		s.Clients[client.Connection.RemoteAddr().String()] = client
	}
}

func (s *Server) Stop() {
	s.Log.Info().Msg("Shutting down server")
	for _, client := range s.Clients {
		client.Connection.Close()
	}
}
