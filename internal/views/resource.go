package views

import (
	"net/http"

	"github.com/deerling/resources.app/internal/models"
)

type systemResponse struct {
}

func RenderResources(w http.ResponseWriter, systems []*models.Resource) {
	RenderResponse(w, http.StatusOK, systems)
}
