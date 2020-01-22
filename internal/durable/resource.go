package durable

import (
	"context"

	"github.com/deerling/resourcepack/internal/models"
)

// ResourceStore is responsible for storing resources
type ResourceStore interface {
	GetResources(context.Context) ([]*models.Resource, error)
}

// GetResources returns all resources saved in store
func (s *store) GetResources(ctx context.Context) ([]*models.Resource, error) {
	// TODO: It's a mock

	return []*models.Resource{
		{ID: 1, Name: "Resource 1"},
		{ID: 2, Name: "Resource 2"},
		{ID: 3, Name: "Resource 3"},
	}, nil
}
