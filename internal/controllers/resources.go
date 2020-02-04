package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jenusek/resourcepack/internal/durable"
	"github.com/jenusek/resourcepack/internal/models"
	"github.com/jenusek/resourcepack/internal/session"
	"github.com/jenusek/resourcepack/internal/views"
)

func registerResourcesEndpoints(store durable.ResourceStore, router *mux.Router) {
	endpoints := resourcesEndpoints{store}

	router.HandleFunc("/resources", endpoints.get).Methods(http.MethodGet)
	router.HandleFunc("/resources", endpoints.post).Methods(http.MethodPost)
	router.HandleFunc("/resources", endpoints.patch).Methods(http.MethodPatch)
}

type resourcesEndpoints struct {
	store durable.ResourceStore
}

func (endpoint *resourcesEndpoints) get(w http.ResponseWriter, r *http.Request) {
	resources, err := endpoint.store.GetResources(r.Context())
	if err != nil {
		views.RenderError(w, err)
		return
	}
	views.RenderResources(w, resources)
}

func (endpoint *resourcesEndpoints) post(w http.ResponseWriter, r *http.Request) {
	user := session.User(r.Context())
	if user.Privileges < models.UserPrivilegesManager {
		views.RenderError(w, session.AuthenticationError())
		return
	}

	var body struct {
		Name string `json:"name"`
		Desc string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderError(w, err)
		return
	}

	err := endpoint.store.AddResource(r.Context(), &models.Resource{
		Name:        body.Name,
		Description: body.Desc,
		CreatedBy:   user,
	})
	if err != nil {
		views.RenderError(w, err)
		return
	}
}

func (endpoint *resourcesEndpoints) patch(w http.ResponseWriter, r *http.Request) {
	user := session.User(r.Context())

	var body struct {
		Duration  string `json:"duration"`
		Resources []int  `json:"resources"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderError(w, err)
		return
	}

	duration, err := time.ParseDuration(body.Duration)
	if err != nil {
		views.RenderError(w, err)
		return
	}

	resources := make([]*models.Resource, 0, len(body.Resources))
	for _, res := range body.Resources {
		resource, err := endpoint.store.GetResource(r.Context(), res)
		if err != nil {
			err := fmt.Errorf("error while getting resource from store")
			views.RenderError(w, err)
			return
		}
		resources = append(resources, resource)
	}

	reservation := &models.Reservation{
		Author:    user.Username,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(duration),
		Resources: resources,
	}

	views.RenderResponse(w, 200, reservation)
}
