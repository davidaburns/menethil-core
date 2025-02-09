package server

import (
	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/rs/zerolog"
)

type ServerType int

const (
	ServerAuth ServerType = iota
	ServerWorld
)

type Server[SC config.ServerConfig] struct {
	log *zerolog.Logger
	config *SC
}
