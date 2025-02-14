package database

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
)

type QueryParser func(rows *sql.Rows) error
type QueryParseSingle func(row *sql.Row) error

type DatabaseClient struct {
	DB                 *sql.DB
	DSN                string
	Driver             DatabaseDriver
	Log                *zerolog.Logger
	preparedStatements map[QueryStatementPath]*sql.Stmt
}

func OpenDatabaseClient(driver string, dsn string, log *zerolog.Logger) (*DatabaseClient, error) {
	log.Info().Msg("Opening database connection")
	log.Info().Msgf("Database Driver: %v", driver)
	log.Debug().Msgf("Database DSN: %v", dsn)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DatabaseClient{
		DB:                 db,
		DSN:                dsn,
		Driver:             GetDatabaseDriver(driver),
		Log:                log,
		preparedStatements: make(map[QueryStatementPath]*sql.Stmt),
	}, nil
}

func (db *DatabaseClient) QueryRaw(query string, args []any, parser QueryParser) error {
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

func (db *DatabaseClient) QueryRawSingle(query string, args []any, parser QueryParseSingle) error {
	row := db.DB.QueryRow(query, args...)
	err := parser(row)
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseClient) ExecuteRaw(query string, args []any) (int64, error) {
	res, err := db.DB.Exec(query, args...)
	if err != nil {
		return -1, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return -1, nil
	}

	return affected, nil
}

func (db *DatabaseClient) QueryPrepared(id QueryStatementPath, args []any, parser QueryParser) error {
	stmt, exists := db.preparedStatements[id]
	if !exists {
		return fmt.Errorf("The query: %v, has not been prepared", id)
	}

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

func (db *DatabaseClient) QueryPreparedSingle(id QueryStatementPath, args []any, parser QueryParseSingle) error {
	stmt, exists := db.preparedStatements[id]
	if !exists {
		return fmt.Errorf("The query: %v, has not been prepared", id)
	}

	row := stmt.QueryRow(args...)
	err := parser(row)
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseClient) ExecutePrepared(id QueryStatementPath, args []any) (int64, error) {
	stmt, exists := db.preparedStatements[id]
	if !exists {
		return -1, fmt.Errorf("The query: %v, has not been prepared", id)
	}

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return -1, nil
	}

	return affected, nil
}

func (db *DatabaseClient) PrepareStatement(id QueryStatementPath, query string) error {
	db.Log.Debug().Msgf("Preparing query: %v", id)
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	db.preparedStatements[id] = stmt
	return nil
}

func (db *DatabaseClient) PrepareEmbeddedQueries(ids []QueryStatementPath) error {
	db.Log.Info().Msgf("Preparing Queries Count: %v", len(ids))
	for _, id := range ids {
		content, err := db.LoadEmbeddedQuery(id)
		if err != nil {
			return fmt.Errorf("Failed to load embedded query: %v: %v", id, err)
		}

		err = db.PrepareStatement(id, content)
		if err != nil {
			return fmt.Errorf("Failed to prepare query: %v: %v", id, err)
		}
	}

	return nil
}

func (db *DatabaseClient) LoadEmbeddedQuery(id QueryStatementPath) (string, error) {
	data, err := embeddedQueries.ReadFile(string(id))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (db *DatabaseClient) PerformMigrations(src string) error {
	db.Log.Info().Msgf("Creating database migration driver for: %v", db.Driver)
	driver, err := GetDatabaseMigrationDriver(db.DB, db.Driver)
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
