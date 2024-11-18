package handlers

import (
	"encoding/json"
	"loggining-monitoring/metrics"
	"net/http"

	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		logrus.WithError(err).Error("Failed to decode login request")
		metrics.IncrementErrorCounter()
		return
	}

	logrus.WithFields(logrus.Fields{
		"username": req.Username,
		"event":    "user_login",
	}).Info("User login attempt")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Login successful"})
	metrics.IncrementRequestCounter()
}

func DataSubmissionHandler(w http.ResponseWriter, r *http.Request) {
	// Example of structured logging
	logrus.WithField("event", "data_submission").Info("Data submission received")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Data submitted successfully"})
	metrics.IncrementRequestCounter()
}


