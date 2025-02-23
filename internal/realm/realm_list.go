package realm

import "github.com/davidaburns/menethil-core/internal/database"

type RealmList struct {
	Realms map[uint32]Realm
}

func GetRealmListFromDb(db *database.DatabaseClient) (*RealmList, error) {
	return nil, nil
}
