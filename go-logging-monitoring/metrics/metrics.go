package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_requests_total",
			Help: "Total number of requests",
		},
		[]string{"endpoint"},
	)
	errorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "app_errors_total",
			Help: "Total number of errors",
		},
	)
)

func InitializeMetrics() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(errorCounter)
}

func IncrementRequestCounter() {
	requestCounter.WithLabelValues("generic").Inc()
}

func IncrementErrorCounter() {
	errorCounter.Inc()
}
