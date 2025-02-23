package realm

type RealmFlag uint8

const (
	REALM_FLAG_NON              RealmFlag = 0x00
	REALM_FLAG_VERSION_MISMATCH RealmFlag = 0x01
	REALM_FLAG_OFFLINE          RealmFlag = 0x02
	REALM_FLAG_SPECIFYBUILD     RealmFlag = 0x04
	REALM_FLAG_UNK1             RealmFlag = 0x08
	REALM_FLAG_UNK2             RealmFlag = 0x10
	REALM_FLAG_RECOMMENDED      RealmFlag = 0x20
	REALM_FLAG_NEW              RealmFlag = 0x40
	REALM_FLAG_FULL             RealmFlag = 0x80
)
