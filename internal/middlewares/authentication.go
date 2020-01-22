package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/deerling/resourcepack/internal/config"
	"github.com/deerling/resourcepack/internal/durable"
	"github.com/deerling/resourcepack/internal/session"
	"github.com/deerling/resourcepack/internal/views"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var whiteList = []string{
	"/auth/token",
}

// Authenticate is a function
func Authenticate(store durable.Datastore) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if isOnWhiteList(r.RequestURI) {
				handler.ServeHTTP(w, r)
				return
			}

			logger := session.Logger(r.Context())

			header := r.Header.Get("Authorization")
			if !strings.HasPrefix(header, "Bearer ") {
				logger.Info("Bearer token is not provided")
				views.RenderError(w, session.AuthenticationError())
				return
			}

			tokenString := header[7:]
			logger = logger.WithField("token", tokenString)

			token, err := parseToken(tokenString)
			if err != nil {
				logger.Error(err)
				views.RenderError(w, err)
				return
			}

			if !token.Valid {
				logger.Error("Given token is invalid")
				views.RenderError(w, session.AuthenticationError())
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				err := errors.New("could not map jwt.Claims to jwt.MapClaims")
				logger.Error(err)
				views.RenderError(w, err)
				return
			}
			logger = logger.WithField("issuer", claims["iss"])

			user, err := store.GetUser(r.Context(), claims["iss"].(string))
			if err != nil {
				logger.Error(err)
				views.RenderError(w, err)
				return
			}

			ctx := session.WithUser(r.Context(), user)
			ctx = session.WithLogger(ctx, logger)

			handler.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}

func isOnWhiteList(uri string) bool {
	for _, endpoint := range whiteList {
		if endpoint == uri {
			return true
		}
	}
	return false
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.SecretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error while parsing token: %w", err)
	}
	return token, nil
}
