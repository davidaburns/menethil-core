package database

type DatabaseDriver string

const (
	DRIVER_POSTGRES DatabaseDriver = "postgres"
	DRIVER_UNKNOWN DatabaseDriver = "unknown"
)

func GetDatabaseDriver(driver string) DatabaseDriver {
	switch driver {
	case "postgres": return DRIVER_POSTGRES
	default: return DRIVER_UNKNOWN
	}
}
