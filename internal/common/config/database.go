package config

type DatabaseConfig struct {
	Driver string `ini:"DatabaseDriver"`
	AuthDSN string `ini:"DatabaseAuthDsn"`
	WorldDSN string `ini:"DatabaseWorldDsn"`
	CharacterDSN string `ini:"DatabaseCharacterDsn"`
	RealmDSN string `ini:"DatabaseRealmDsn"`
}
