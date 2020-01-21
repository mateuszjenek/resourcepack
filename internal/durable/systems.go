package durable

import (
	"context"

	"github.com/deerling/resources.app/internal/models"
)

const (
	sqlQuerySystems = "SELECT * FROM systems"
)

type SystemStore interface {
	GetSystems(context.Context) ([]*models.System, error)
}

func (s *store) GetSystems(ctx context.Context) ([]*models.System, error) {
	// rows, err := s.db.QueryContext(ctx, sqlQuerySystems)
	// if err != nil {
	// 	return nil, fmt.Errorf("error occurred while quering to database: %w", err)
	// }
	// var systems []*models.System
	// for rows.Next() {
	// 	system := &models.System{}
	// 	rows.Scan(system.ID, system.Type, system.UUID, system.Address, system.Username, system.Password)
	// 	systems = append(systems, system)
	// }
	// return systems, nil

	return []*models.System{
		{ID: 1, Type: models.SystemTypeHardware, UUID: "uuid", Address: "123.123.123.123:234", Username: "admin", Password: "admin"},
		{ID: 2, Type: models.SystemTypeSimulator, UUID: "uuid", Address: "123.123.123.123:234", Username: "admin", Password: "admin"},
		{ID: 3, Type: models.SystemTypeHardware, UUID: "uuid", Address: "123.123.123.123:234", Username: "admin", Password: "admin"},
	}, nil
}
