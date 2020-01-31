package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jenusek/resourcepack/internal/controllers"
	"github.com/jenusek/resourcepack/internal/durable"
	"github.com/jenusek/resourcepack/internal/middlewares"
	"github.com/jenusek/resourcepack/internal/models"
	"github.com/jenusek/resourcepack/internal/services"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	config, err := models.LoadConfigurationFromFile("/home/mjenek/Desktop/resourcepack/config/development.yaml")
	if err != nil {
		logrus.Errorf("Error while reading configuration file: %v", err)
		return
	}
	datastore, err := durable.OpenDatastore("./development.db")
	if err != nil {
		logrus.Errorf("Error while opening datastore: %v", err)
		return
	}
	defer datastore.Close()

	createRootUserIfNotExists(datastore, config)

	router := &mux.Router{}
	router = router.PathPrefix("/respack/v1").Subrouter()

	router.Use(middlewares.Logger)
	router.Use(middlewares.Authenticate(datastore, config))
	controllers.Register(datastore, config, router)

	srv := http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", config.Server.IP, config.Server.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logrus.Fatal(srv.ListenAndServe())
}

func createRootUserIfNotExists(store durable.UserStore, config *models.Configuration) {
	hash, err := services.EncryptPassword(config.RootUser.Password)
	if err != nil {
		logrus.Errorf("Error while hashing root user password: %v", err)
		return
	}
	store.AddUser(context.Background(), &models.User{
		Username:   config.RootUser.Username,
		Email:      config.RootUser.Email,
		PassHash:   string(hash),
		Privileges: models.UserPrivilegesAdmin,
	})
}
