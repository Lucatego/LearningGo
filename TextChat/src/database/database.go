package database

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbms = "sqlite3"
)

var (
	DBManager = Database{connection: nil, location: ""}
)

type Database struct {
	connection *sql.DB
	location   string
}

func (DBManager *Database) Initialize(location string) error {
	// Open the file
	conn, err := sql.Open(dbms, location)
	if err != nil {
		return err
	}
	// Test connection
	err = conn.Ping()
	if err != nil {
		return err
	}
	// Save data
	DBManager.connection = conn
	DBManager.location = location
	return nil
}

func (DBManager *Database) GetInstance() (*sql.DB, error) {
	if DBManager.connection == nil {
		return nil, errors.New("error: the database connection is not initialized")
	}
	return DBManager.connection, nil
}

func (DBManager *Database) CloseDB() error {
	return DBManager.connection.Close()
}
