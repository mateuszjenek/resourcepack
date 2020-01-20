package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/deerling/resource-bridge/internal/durable"
	"github.com/deerling/resource-bridge/internal/session"
	"github.com/deerling/resource-bridge/internal/views"
)

func registerAuthorizationEndpoints(store durable.UserStore, router *mux.Router) {
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
		views.RenderError(w, session.AuthorizationError())
		return
	}

	views.RenderToken(w, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

}
