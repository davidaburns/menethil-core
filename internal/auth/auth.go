package auth

import (
	"net"

	"github.com/davidaburns/menethil-core/internal/client"
	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/rs/zerolog"
)

type AuthServer struct {
	Log *zerolog.Logger
	Config *config.Config
	Clients map[string]*client.Client
}

func NewAuthServer(log *zerolog.Logger, conf *config.Config) *AuthServer {
	return &AuthServer{
		Log: log,
		Config: conf,
		Clients: make(map[string]*client.Client),
	}
}

func (as *AuthServer) Start() {
	as.Log.Info().Msg("Initializing auth server")
}

func (as *AuthServer) Stop() {
	as.Log.Info().Msg("Cleaning up auth server")
	for _, client := range as.Clients {
		client.Connection.Close()
	}
}

func (as *AuthServer) HandleConnection(conn net.Conn) {
	as.Log.Info().Msgf("Client Local Network: %v", conn.LocalAddr().String())
	as.Log.Info().Msgf("Client Remote Network: %v", conn.RemoteAddr().String())

	cl := client.FromConnection(conn)
	as.Clients[cl.Connection.RemoteAddr().String()] = cl
}
