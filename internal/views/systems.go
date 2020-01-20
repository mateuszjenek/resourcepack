package views

import "net/http"

import "github.com/deerling/resource-bridge/internal/models"

type systemResponse struct {
}

func RenderSystems(w http.ResponseWriter, systems []*models.System) {
	RenderResponse(w, http.StatusOK, systems)
}
