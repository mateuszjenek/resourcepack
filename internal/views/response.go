package views

import "net/http"

import "fmt"

import "github.com/deerling/resource-bridge/internal/session"

import "encoding/json"

func RenderResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		RenderError(w, err)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func RenderError(w http.ResponseWriter, err error) {
	sessionError, ok := err.(session.Error)
	if !ok {
		sessionError = session.ServerError(err)
	}
	w.WriteHeader(sessionError.StatusCode)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("%s", sessionError.Description)))
}
