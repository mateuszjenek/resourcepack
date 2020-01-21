package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/deerling/resources.app/internal/durable"
	"github.com/deerling/resources.app/internal/session"
	"github.com/deerling/resources.app/internal/views"
	"github.com/gorilla/mux"
)

func registerAuthenticationEndpoints(store durable.UserStore, router *mux.Router) {
	endpoints := authEndpoints{store}

	router.HandleFunc("/auth/token", endpoints.token).Methods(http.MethodPost)
}

type authEndpoints struct {
	store durable.UserStore
}

func (endpoint *authEndpoints) token(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderError(w, err)
		return
	}

	user, err := endpoint.store.GetUser(r.Context(), body.Username)
	if err != nil {
		views.RenderError(w, err)
		return
	}

	if user.PassHash != body.Password {
		views.RenderError(w, session.AuthenticationError())
		return
	}

	token, err := user.GenerateToken()
	if err != nil {
		views.RenderError(w, session.ServerError(err))
		return
	}

	views.RenderToken(w, token)

}