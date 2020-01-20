package controllers

import (
	"github.com/gorilla/mux"
	"github.com/deerling/resource-bridge/internal/durable"
)

func Register(store durable.Datastore, router *mux.Router) {
	registerSystemsEndpoints(store, router)
	registerAuthorizationEndpoints(store, router)
}
