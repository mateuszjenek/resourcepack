package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/deerling/resource-bridge/internal/controllers"
	"github.com/deerling/resource-bridge/internal/durable"
	"github.com/deerling/resource-bridge/internal/middlewares"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	datastore, err := durable.OpenDatastore("./development.db")
	if err != nil {
		logrus.Errorf("Error while opening datastore: %w", err)
		return
	}
	defer datastore.Close()

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
