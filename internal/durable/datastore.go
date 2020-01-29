package durable

import (
	"database/sql"
	"fmt"

	// SQLite3 database driver
	_ "github.com/mattn/go-sqlite3"
)

// Datastore aggregates all Store interfaces
type Datastore interface {
	ResourceStore
	UserStore

	Close()
}

type store struct {
	db *sql.DB
}

// OpenDatastore initalize datastore object
func OpenDatastore(dataSourceName string) (Datastore, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error while creating handler to database: %v", err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error while connecting to database: %v", err)
	}

	datastore := &store{db}
	err = datastore.createUserStore()
	if err != nil {
		return nil, fmt.Errorf("error while creating user store: %v", err)
	}
	err = datastore.createResourceStore()
	if err != nil {
		return nil, fmt.Errorf("error while creating resource store: %v", err)
	}

	return datastore, nil
}

func (s *store) Close() {
	s.db.Close()
}
