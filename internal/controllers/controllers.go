package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jenusek/resourcepack/internal/durable"
	"github.com/jenusek/resourcepack/internal/models"
)

// Register is responsible for registration all systems endpoints
func Register(store durable.Datastore, config *models.Configuration, router *mux.Router) {
	registerResourcesEndpoints(store, router)
	registerAuthenticationEndpoints(store, config, router)
	registerUserEndpoints(store, config, router)
}
