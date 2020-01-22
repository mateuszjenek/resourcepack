package durable

import (
	"context"
	"fmt"

	"github.com/deerling/resources.app/internal/models"
	"github.com/sirupsen/logrus"
)

const (
	queryCreateTableUser      = "CREATE TABLE IF NOT EXISTS user (username varchar(50) PRIMARY KEY, passhash varchar(50) NOT NULL, privileges integer NOT NULL, email varchar(255) NOT NULL)"
	queryInsertUser           = "INSERT INTO user (username, passhash, privileges, email) VALUES ($1, $2, $3, $4)"
	querySelectUserByUsername = "SELECT username, passhash, privileges, email from user WHERE username = $1"
)

func (s *store) createUserStore() error {
	_, err := s.db.Exec(queryCreateTableUser)
	if err != nil {
		return fmt.Errorf("error while creating table in database: %w", err)
	}
	logrus.Debug("Successfully created user table")
	return nil
}

// UserStore is responsible for storing users
type UserStore interface {
	GetUser(ctx context.Context, username string) (*models.User, error)
	AddUser(ctx context.Context, user *models.User) error
}

// GetUser returns user with given username from store
func (s *store) GetUser(ctx context.Context, username string) (*models.User, error) {
	row := s.db.QueryRow(querySelectUserByUsername, username)
	user := &models.User{}
	err := row.Scan(&user.Username, &user.PassHash, &user.Role, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("error while scanning row: %w", err)
	}
	return user, nil
}

// AddUser adds given user to store
func (s *store) AddUser(ctx context.Context, user *models.User) error {
	result, err := s.db.Exec(queryInsertUser, user.Username, user.PassHash, user.Role, user.Email)
	if err != nil {
		return fmt.Errorf("error while inserting user to user table: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while getting number of affected rows: %w", err)
	}
	logrus.WithField("rowsAffected", rowsAffected).Info("Succesfully added user")
	return nil
}
