package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Wrap response writer to capture status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)

		duration := time.Since(startTime)
		logrus.WithFields(logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   lrw.statusCode,
			"duration": duration.String(),
		}).Info("HTTP Request")
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func GetLogFile(logFilePath string) *os.File {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}
	return file
}


func InitLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    10, // MB
		MaxBackups: 5,  // Number of old files to retain
		MaxAge:     30, // Days to retain old logs
		Compress:   true,
	})
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}
