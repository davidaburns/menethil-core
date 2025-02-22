package realm

type RealmType uint8

const (
	REALM_TYPE_NORMAL RealmType = 0
	REALM_TYPE_PVP RealmType = 1
	REALM_TYPE_NORMAL2 RealmType = 4
	REALM_TYPE_RP RealmType = 6
	REALM_TYPE_RPPVP RealmType = 8
	REALM_TYPE_MAX_CLIENT_TYPE RealmType = 14
)
