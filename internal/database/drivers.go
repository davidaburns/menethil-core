package database

import (
	"fmt"

	dbm "github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

type DatabaseDriver string

const (
	DRIVER_POSTGRES DatabaseDriver = "postgres"
	DRIVER_MYSQL    DatabaseDriver = "mysql"
	DRIVER_UNKNOWN  DatabaseDriver = "unknown"
)

func GetDatabaseDriver(driver string) DatabaseDriver {
	switch driver {
	case "postgres":
		return DRIVER_POSTGRES
	case "mysql":
		return DRIVER_MYSQL
	default:
		return DRIVER_UNKNOWN
	}
}

func GetDatabaseMigrationDriver(db *sqlx.DB, driver DatabaseDriver) (dbm.Driver, error) {
	var migrationDriver dbm.Driver
	var err error

	switch driver {
	case DRIVER_POSTGRES:
		migrationDriver, err = postgres.WithInstance(db.DB, &postgres.Config{})
	case DRIVER_MYSQL:
		migrationDriver, err = mysql.WithInstance(db.DB, &mysql.Config{})
	default:
		err = fmt.Errorf("Cannot create database migration driver, driver unkown")
	}

	if err != nil {
		return nil, err
	}

	return migrationDriver, nil
}
