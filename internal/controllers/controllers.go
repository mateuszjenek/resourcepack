package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jenusek/resourcepack/internal/durable"
)

// Register is responsible for registration all systems endpoints
func Register(store durable.Datastore, router *mux.Router) {
	registerResourcesEndpoints(store, router)
	registerAuthenticationEndpoints(store, router)
	registerUserEndpoints(store, router)
}
