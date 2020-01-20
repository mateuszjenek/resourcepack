package durable

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Datastore interface {
	SystemStore
	UserStore

	Close()
}

type store struct {
	db *sql.DB
}

func OpenDatastore(dataSourceName string) (Datastore, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error while creating handler to database: %w", err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error while connecting to database: %w", err)
	}

	return &store{db}, nil
}

func (s *store) Close() {
	s.db.Close()
}
