package middlewares

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.
			WithField("ipSource", r.RemoteAddr).
			WithField("url", r.URL).
			WithField("method", r.Method).
			Debug("Recieved HTTP request")

		handler.ServeHTTP(w, r)
	})
}
