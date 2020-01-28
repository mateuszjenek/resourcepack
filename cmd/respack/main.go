package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jenusek/resourcepack/internal/controllers"
	"github.com/jenusek/resourcepack/internal/durable"
	"github.com/jenusek/resourcepack/internal/middlewares"
	"github.com/jenusek/resourcepack/internal/models"
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

	err = datastore.AddResource(&models.Resource{2, "name", "description", user, nil})
	if err != nil {
		logrus.Errorf("Error while creating resource: %w", err)
		return
	}

	router := &mux.Router{}
	router.Use(middlewares.Logger)
	router.Use(middlewares.Authenticate(datastore))
	controllers.Register(datastore, router)

	srv := http.Server{
		Handler:      router,
		Addr:         ":2002",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logrus.Fatal(srv.ListenAndServe())
}
