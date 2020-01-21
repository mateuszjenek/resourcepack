package middlewares

import (
	"net/http"

	"github.com/deerling/resources.app/internal/session"
	"github.com/sirupsen/logrus"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := logrus.WithField("ipSource", r.RemoteAddr).WithField("url", r.URL).WithField("method", r.Method)
		logger.Debug("Recieved HTTP request")
		ctx := session.WithLogger(r.Context(), logger)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
