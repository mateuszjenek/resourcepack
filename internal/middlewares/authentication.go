package middlewares

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/deerling/resources.app/internal/durable"
)

var whiteList = []string{
	"/auth/token",
}

// Authenticate is a function
func Authenticate(store durable.Datastore) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, endpoint := range whiteList {
				if endpoint == r.RequestURI {
					handler.ServeHTTP(w, r)
					return
				}
			}

			header := r.Header.Get("Authorization")
			if !strings.HasPrefix(header, "Bearer ") {
				logrus.
					WithField("ipSource", r.RemoteAddr).
					WithField("url", r.URL).
					WithField("method", r.Method).
					Info("Unauthorized access")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			handler.ServeHTTP(w, r)
		})
	}
}
