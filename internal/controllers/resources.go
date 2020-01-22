package controllers

import (
	"net/http"

	"github.com/deerling/resourcepack/internal/durable"
	"github.com/deerling/resourcepack/internal/views"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func registerResourcesEndpoints(store durable.ResourceStore, router *mux.Router) {
	endpoints := resourcesEndpoints{store}

	router.HandleFunc("/resources", endpoints.get).Methods(http.MethodGet)
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
