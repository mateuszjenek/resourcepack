package main

import (
	"context"
	"net/http"
	"time"

	"github.com/deerling/resourcepack/internal/controllers"
	"github.com/deerling/resourcepack/internal/durable"
	"github.com/deerling/resourcepack/internal/middlewares"
	"github.com/deerling/resourcepack/internal/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	datastore, err := durable.OpenDatastore("./development.db")
	if err != nil {
		logrus.Errorf("Error while opening datastore: %w", err)
		return
	}
	defer datastore.Close()

	err = datastore.AddUser(context.Background(), &models.User{"username", "email", "passhash", models.UserPrivilegesAdmin})
	if err != nil {
		logrus.Errorf("Error while creating user: %w", err)
		return
	}

	user, err := datastore.GetUser(context.Background(), "username")
	if err != nil {
		logrus.Errorf("error while getting user: %w", err)
		return
	}
	logrus.Info("%v", *user)

	router := &mux.Router{}
	router.Use(middlewares.Logger)
	router.Use(middlewares.Authenticate(datastore))
	controllers.Register(datastore, router)

	srv := http.Server{
		Handler:      router,
		Addr:         ":2000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
	logrus.Fatal("XD")
}
