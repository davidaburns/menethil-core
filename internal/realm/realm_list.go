package realm

import (
	"github.com/davidaburns/menethil-core/internal/database"
	"github.com/rs/zerolog"
)

type RealmList struct {
	Realms map[uint32]Realm
	Log    *zerolog.Logger
}

func NewRealmListFromDb(db *database.DatabaseClient, log *zerolog.Logger) (*RealmList, error) {
	var realms []Realm
	err := db.QueryPrepared(database.AUTH_SELECT_REALMLIST, nil, &realms)
	if err != nil {
		return nil, err
	}


	realmsMap := make(map[uint32]Realm)
	for _, realm := range realms {
		realmsMap[realm.ID] = realm
	}

	return &RealmList{Realms: realmsMap, Log: log}, nil
}

func (rl *RealmList) ResolveRealms() {
	rl.Log.Info().Msg("Resolving realm ips")
	for _, realm := range rl.Realms {
		if !realm.ExternalAddress.IsResovable() {
			rl.Log.Warn().Msgf("Realm: %v external address: %v is not reachable", realm.ID, realm.ExternalAddress.String())
			continue
		}
		if !realm.LocalAddress.IsResovable() {
			rl.Log.Warn().Msgf("Realm: %v local address: %v is not reachable", realm.ID, realm.LocalAddress.String())
			continue
		}

		rl.Log.Info().Msgf("Realm: %v external address: %v is reachable", realm.ID, realm.ExternalAddress.String())
		rl.Log.Info().Msgf("Realm: %v local address: %v is reachable", realm.ID, realm.LocalAddress.String())
	}
}
