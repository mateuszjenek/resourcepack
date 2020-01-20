package views

import "net/http"

type tokenResponse struct {
	Token string `json:"token"`
}

func RenderToken(w http.ResponseWriter, token string) {
	RenderResponse(w, http.StatusOK, tokenResponse{token})
}
