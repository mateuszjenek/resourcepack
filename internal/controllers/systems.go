package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/deerling/resources.app/internal/durable"
	"github.com/deerling/resources.app/internal/views"
)

func registerSystemsEndpoints(store durable.SystemStore, router *mux.Router) {
	endpoints := systemsEndpoints{store}

	router.HandleFunc("/systems", endpoints.get).Methods(http.MethodGet)
}

type systemsEndpoints struct {
	store durable.SystemStore
}

func (endpoint *systemsEndpoints) get(w http.ResponseWriter, r *http.Request) {
	systems, err := endpoint.store.GetSystems(r.Context())
	if err != nil {
		logrus.Error(err)
	}
	views.RenderSystems(w, systems)

}
