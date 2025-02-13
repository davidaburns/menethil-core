package config

type DatabaseConfig struct {
	Driver string `ini:"DatabaseDriver"`
	DSN string `ini:"DatabaseDsn"`
	MigrationSrc string `ini:"DatabaseMigratioSrc"`
}

func DatabaseConfigFromConfig(c *Config) (DatabaseConfig, error) {
	driver, err := GetValue[string](c, SECTION_DATABASE, DATABASE_DRIVER)
	if err != nil {
		return DatabaseConfig{}, err
	}

	dsn, err := GetValue[string](c, SECTION_DATABASE, DATABASE_DSN)
	if err != nil {
		return DatabaseConfig{}, err
	}

	migrationSrc, err := GetValue[string](c, SECTION_DATABASE, DATABASE_MIGRATION_SRC)
	if err != nil {
		return DatabaseConfig{}, err
	}

	return DatabaseConfig{
		Driver: driver,
		DSN: dsn,
		MigrationSrc: migrationSrc,
	}, nil
}
