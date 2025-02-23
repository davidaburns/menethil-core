package realm

import (
	"database/sql"

	"github.com/davidaburns/menethil-core/internal/account"
)

type Realm struct {
	ID                   uint32
	Build                uint32
	ExternalAddress      string
	LocalAddress         string
	LocalSubnetMask      string
	Port                 uint16
	Name                 string
	Type                 RealmType
	Flags                RealmFlag
	Timezone             uint8
	AllowedSecurityLevel account.AccountType
	PopulationLevel      float32
}

func RealmFromSqlRow(row *sql.Row) (*Realm, error) {
	var realm *Realm
	return realm, nil
}
