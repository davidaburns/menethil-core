package realm

import "github.com/davidaburns/menethil-core/internal/account"

type Realm struct {
	ID uint32
	Build uint32
	ExternalAddress string
	LocalAddress string
	LocalSubnetMask string
	Port uint16
	Name string
	Type RealmType
	Flags RealmFlag
	Timezone uint8
	AllowedSecurityLevel account.AccountType
	PopulationLevel float32
}
