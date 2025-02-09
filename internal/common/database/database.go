package database

import (
	"database/sql"
)

type PreparedStatementId int

type DatabaseClient struct {
	DB *sql.DB
	preparedStatements map[PreparedStatementId]*sql.Stmt
}

func NewDatabaseClient(driver string, dsn string) (*DatabaseClient, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DatabaseClient{DB: db}, nil
}

func (db *DatabaseClient) Query(query string, args []any, parser func(rows *sql.Rows) error) error {
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

func (db *DatabaseClient) QuerySingle(query string, args []any, parser func(row *sql.Row) error) error {
	row := db.DB.QueryRow(query, args...)
	err := parser(row)
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseClient) QueryPrepared(id PreparedStatementId, args []any, parser func(rows *sql.Rows) error) error {
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

func (db *DatabaseClient) QueryPreparedSingle(id PreparedStatementId, args []any, parser func(row *sql.Row) error) error {
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

func (db *DatabaseClient) Close() error {
	return db.DB.Close()
}

func (db *DatabaseClient) IsAlive() bool {
	return db.DB.Ping() != nil
}
