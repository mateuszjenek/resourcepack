package views

import (
	"net/http"

	"github.com/jenusek/resourcepack/internal/models"
)

type resourceResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy"`
}

func RenderResources(w http.ResponseWriter, resources []*models.Resource) {
	response := make([]resourceResponse, 0, len(resources))
	for _, resource := range resources {
		response = append(response, resourceResponse{
			ID:          resource.ID,
			Name:        resource.Name,
			Description: resource.Description,
			CreatedBy:   resource.CreatedBy.Username,
		})
	}
	RenderResponse(w, http.StatusOK, response)
}
