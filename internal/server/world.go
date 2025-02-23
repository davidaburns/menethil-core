package server

import (
	"net"

	"github.com/davidaburns/menethil-core/internal/client"
	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/rs/zerolog"
)

type WorldServer struct {
	Log     *zerolog.Logger
	Config  *config.Config
	Clients map[string]*client.Client
}

func NewWorldServer(log *zerolog.Logger, conf *config.Config) *WorldServer {
	return &WorldServer{
		Log:     log,
		Config:  conf,
		Clients: make(map[string]*client.Client),
	}
}

func (ws *WorldServer) Start() error {
	return nil
}

func (ws *WorldServer) Stop() {}

func (ws *WorldServer) HandleConnection(conn net.Conn) {
	cl := client.FromConnection(conn)
	ws.Clients[cl.Connection.RemoteAddr().String()] = cl
}
