package server

import (
	"fmt"
	"net"

	"github.com/davidaburns/menethil-core/internal/client"
	"github.com/davidaburns/menethil-core/internal/config"
	"github.com/davidaburns/menethil-core/internal/database"
	"github.com/davidaburns/menethil-core/internal/realm"
	"github.com/rs/zerolog"
)

type AuthServer struct {
	Log     *zerolog.Logger
	Config  *config.Config
	Clients map[string]*client.Client
	AuthDB  *database.DatabaseClient
	RealmList *realm.RealmList
}

func NewAuthServer(log *zerolog.Logger, conf *config.Config) *AuthServer {
	return &AuthServer{
		Log:     log,
		Config:  conf,
		Clients: make(map[string]*client.Client),
		AuthDB:  nil,
		RealmList: nil,
	}
}

func (as *AuthServer) Start() error {
	as.Log.Info().Msg("Initializing auth server")
	dbConfig, err := config.DatabaseConfigFromConfig(as.Config)
	if err != nil {
		return err
	}

	as.AuthDB, err = database.OpenDatabaseClient(dbConfig.Driver, fmt.Sprintf("%v/menethil_auth?multiStatements=true", dbConfig.DSN), as.Log)
	if err != nil {
		return err
	}

	err = as.AuthDB.PerformMigrations(fmt.Sprintf("%v/auth", dbConfig.MigrationSrc))
	if err != nil {
		return err
	}

	err = as.AuthDB.PrepareEmbeddedQueries(database.AuthQueries)
	if err != nil {
		return err
	}

	as.RealmList, err = realm.NewRealmListFromDb(as.AuthDB, as.Log)
	if err != nil {
		return err
	}

	as.RealmList.ResolveRealms()
	return nil
}

func (as *AuthServer) Stop() {
	as.Log.Info().Msg("Cleaning up auth server")
	for _, client := range as.Clients {
		client.Connection.Close()
	}

	as.AuthDB.Close()
}

func (as *AuthServer) HandleConnection(conn net.Conn) {
	as.Log.Info().Msgf("Client Local Network: %v", conn.LocalAddr().String())
	as.Log.Info().Msgf("Client Remote Network: %v", conn.RemoteAddr().String())

	cl := client.FromConnection(conn)
	as.Clients[cl.Connection.RemoteAddr().String()] = cl
}
