package database

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	dbm "github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
)

type PreparedStatementId string
type QueryParser func (rows *sql.Rows) error
type QueryParseSingle func (row *sql.Row) error

type DatabaseClient struct {
	DB *sql.DB
	DSN string
	Driver DatabaseDriver
	Log *zerolog.Logger
	preparedStatements map[PreparedStatementId]*sql.Stmt
}

func NewDatabaseClient(driver string, dsn string, log *zerolog.Logger) (*DatabaseClient, error) {
	log.Info().Msg("Opening database connection")
	log.Info().Msgf("Database Driver: %v", driver)
	log.Info().Msgf("Database DSN: %v", dsn)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DatabaseClient{
		DB: db,
		DSN: dsn,
		Driver: GetDatabaseDriver(driver),
		Log: log,
		preparedStatements: make(map[PreparedStatementId]*sql.Stmt),
	}, nil
}

func (db *DatabaseClient) Query(query string, args []any, parser QueryParser) error {
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if err := parser(rows); err != nil {
		return err
	}

	return nil
}

func (db *DatabaseClient) QuerySingle(query string, args []any, parser QueryParseSingle) error {
	row := db.DB.QueryRow(query, args...)
	err := parser(row)
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseClient) QueryPrepared(id PreparedStatementId, args []any, parser QueryParser) error {
	stmt := db.preparedStatements[id]
	rows, err := stmt.Query(args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if err := parser(rows); err != nil {
		return err
	}

	return nil
}

func (db *DatabaseClient) QueryPreparedSingle(id PreparedStatementId, args []any, parser QueryParseSingle) error {
	stmt := db.preparedStatements[id]
	row := stmt.QueryRow(args...)
	err := parser(row)
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseClient) PrepareStatement(id PreparedStatementId, query string) error {
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	db.preparedStatements[id]=stmt
	return nil
}

func (db *DatabaseClient) PerformMigrations(src string) error {
	var driver dbm.Driver
	var err error

	db.Log.Info().Msgf("Creating database migration driver for: %v", db.Driver)
	switch db.Driver {
	case DRIVER_POSTGRES:
		driver, err = postgres.WithInstance(db.DB, &postgres.Config{})
	default:
		err = fmt.Errorf("Cannot perform database migrations, driver unkown")
	}

	if err != nil {
		return err
	}

	db.Log.Info().Msgf("Database migration source path: %v", src)
	m, err := migrate.NewWithDatabaseInstance(src, string(db.Driver), driver)
	if err != nil {
		return err
	}

	db.Log.Info().Msg("Performing database migrations")
	m.Log = NewDatabaseLogger(db.Log)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (db *DatabaseClient) Close() error {
	db.Log.Info().Msg("Closing database connection")
	return db.DB.Close()
}

func (db *DatabaseClient) IsAlive() bool {
	return db.DB.Ping() != nil
}
