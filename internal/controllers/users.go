package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jenusek/resourcepack/internal/durable"
	"github.com/jenusek/resourcepack/internal/models"
	"github.com/jenusek/resourcepack/internal/session"
	"github.com/jenusek/resourcepack/internal/views"
)

func registerUserEndpoints(store durable.UserStore, router *mux.Router) {
	endpoints := &usersEndpoints{store}

	router.HandleFunc("/users", endpoints.post).Methods(http.MethodPost)
}

type usersEndpoints struct {
	store durable.UserStore
}

func (endpoint *usersEndpoints) post(w http.ResponseWriter, r *http.Request) {
	user := session.User(r.Context())
	if user.Privileges < models.UserPrivilegesManager {
		views.RenderError(w, session.AuthenticationError())
		return
	}

	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderError(w, err)
		return
	}

	createdUser := &models.User{
		Username:   body.Username,
		Email:      body.Email,
		Privileges: models.UserPrivilegesStandard,
	}

	_, err := createdUser.GeneratePassword()
	if err != nil {
		views.RenderError(w, err)
		return
	}

	err = endpoint.store.AddUser(r.Context(), createdUser)
	if err != nil {
		views.RenderError(w, err)
		return
	}

	// TODO: Notify via email
}
