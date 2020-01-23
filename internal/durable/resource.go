package durable

import (
	"context"
	"fmt"

	"github.com/deerling/resourcepack/internal/models"
	"github.com/sirupsen/logrus"
)

const (
	queryCreateTableResource = "CREATE TABLE IF NOT EXISTS resource (id int PRIMARY KEY, name varchar(50), description varchar(255), created_by varchar(50), FOREIGN KEY(created_by) REFERENCES user(username))"
	queryGetAllResources     = "SELECT id, name, description, created_by FROM resource"
	queryResourceInsert      = "INSERT INTO resource (id, name, description, created_by) VALUES ($1, $2, $3 ,$4)"
)

// ResourceStore is responsible for storing resources
type ResourceStore interface {
	GetResources(context.Context) ([]*models.Resource, error)
	AddResource(*models.Resource) error
}

func (s *store) createResourceStore() error {
	_, err := s.db.Exec(queryCreateTableResource)
	if err != nil {
		return fmt.Errorf("error while creating table in database: %w", err)
	}
	logrus.Info("Successfully created resources table")
	return nil
}

// GetResources returns all resources saved in store
func (s *store) GetResources(ctx context.Context) ([]*models.Resource, error) {
	rows, err := s.db.Query(queryGetAllResources)
	if err != nil {
		return nil, fmt.Errorf("error while getting resources from store: %w", err)
	}
	defer rows.Close()
	resources := make([]*models.Resource, 0)
	for rows.Next() {
		resource := &models.Resource{}
		user := &models.User{}
		err = rows.Scan(&resource.ID, &resource.Name, &resource.Description, &user.Username)
		if err != nil {
			return nil, fmt.Errorf("error while scanning row: %w", err)
		}
		user, err = s.GetUser(ctx, user.Username)
		if err != nil {
			return nil, fmt.Errorf("error while getting user from store: %w", err)
		}
		resource.CreatedBy = user
		resources = append(resources, resource)
	}
	return resources, nil
}

func (s *store) AddResource(resource *models.Resource) error {
	result, err := s.db.Exec(queryResourceInsert, resource.ID, resource.Name, resource.Description, resource.CreatedBy.Username)
	if err != nil {
		return fmt.Errorf("error while inserting resource to resource table: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while getting number of affected rows: %w", err)
	}
	logrus.WithField("rowsAffected", rowsAffected).Info("Succesfully added resource")
	return nil
}
