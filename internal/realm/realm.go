package realm

import (
	"github.com/davidaburns/menethil-core/internal/account"
	"github.com/davidaburns/menethil-core/internal/network"
)

type Realm struct {
	ID                   uint32              `db:"id"`
	Build                uint32              `db:"gamebuild"`
	ExternalAddress      network.IP          `db:"address"`
	LocalAddress         network.IP          `db:"localAddress"`
	LocalSubnetMask      string              `db:"localSubnetMask"`
	Port                 uint16              `db:"port"`
	Name                 string              `db:"name"`
	Type                 RealmType           `db:"icon"`
	Flags                RealmFlag           `db:"flag"`
	Timezone             uint8               `db:"timezone"`
	AllowedSecurityLevel account.AccountType `db:"allowedSecurityLevel"`
	PopulationLevel      float32             `db:"population"`
}
