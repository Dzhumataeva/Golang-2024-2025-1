package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"go-logging-monitoring/metrics"
	"go-logging-monitoring/middleware"
)


func main() {
	// Initialize logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(middleware.GetLogFile("logs/app.log"))

	// Initialize Prometheus metrics
	metrics.InitializeMetrics()

	// Create Router
	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.RequestLogger)

	// Handlers
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/submit", handlers.DataSubmissionHandler).Methods("POST")
	router.Handle("/metrics", promhttp.Handler()).Methods("GET")

	// Start Server
	logrus.Info("Starting server on port 8080")
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	
}




