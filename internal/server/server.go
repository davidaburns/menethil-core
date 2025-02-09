package server

import (
	"github.com/davidaburns/menethil-core/build"
	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/rs/zerolog"
)

type ServerType int

const (
	ServerAuth ServerType = iota
	ServerWorld
)

func (st *ServerType) String() string {
	switch *st {
	case ServerAuth: return "Authentication"
	case ServerWorld: return "World"
	default: return "Unkown"
	}
}

type Server[SC config.ServerConfig] struct {
	build build.BuildInfo
	log *zerolog.Logger
	config *SC
	serverType ServerType
}

func NewServer[SC config.ServerConfig](serverType ServerType, conf *SC, log *zerolog.Logger) (*Server[SC], error) {
	return &Server[SC]{
		build: *build.NewBuildInfo(),
		log: log,
		config: conf,
		serverType: serverType,
	}, nil
}

func (s *Server[SC]) Start() {
	s.log.Info().Msgf("Menethil Core: %v Server", s.serverType.String())
	s.log.Info().Msgf("Version: %v", s.build.String())
	s.log.Info().Msg("Starting server")
}

func (s *Server[SC]) Stop() {
	s.log.Info().Msg("Shutting down server")
}
