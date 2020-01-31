package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jenusek/resourcepack/internal/durable"
	"github.com/jenusek/resourcepack/internal/models"
	"github.com/jenusek/resourcepack/internal/session"
	"github.com/jenusek/resourcepack/internal/views"
	"github.com/sirupsen/logrus"
)

func registerResourcesEndpoints(store durable.ResourceStore, router *mux.Router) {
	endpoints := resourcesEndpoints{store}

	router.HandleFunc("/resources", endpoints.get).Methods(http.MethodGet)
	router.HandleFunc("/resources", endpoints.post).Methods(http.MethodPost)
}

type resourcesEndpoints struct {
	store durable.ResourceStore
}

func (endpoint *resourcesEndpoints) get(w http.ResponseWriter, r *http.Request) {
	resources, err := endpoint.store.GetResources(r.Context())
	if err != nil {
		logrus.Error(err)
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

	err := endpoint.store.AddResource(&models.Resource{
		Name:        body.Name,
		Description: body.Desc,
		CreatedBy:   user,
	})
	if err != nil {
		views.RenderError(w, err)
		return
	}
}
