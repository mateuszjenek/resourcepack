package durable

import (
	"context"
	"errors"

	"github.com/deerling/resources.app/internal/models"
)

type UserStore interface {
	GetUser(ctx context.Context, username string) (*models.User, error)
}

func (s *store) GetUser(ctx context.Context, username string) (*models.User, error) {
	// TODO: It's a mock

	if username == "admin" {
		return &models.User{Username: "admin", PassHash: "passHash"}, nil
	}
	return nil, errors.New("WRONGUSER")
}
