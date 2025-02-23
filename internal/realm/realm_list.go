package realm

import "github.com/rs/zerolog"

type RealmList struct {
	Realms map[uint32]Realm
	Log    *zerolog.Logger
}

func NewRealmList(log *zerolog.Logger) *RealmList {
	return &RealmList{Realms: map[uint32]Realm{}, Log: log}
}
