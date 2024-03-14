// metrics.go
package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// echoRequestsTotal is a counter that tracks the total number of Echo requests.
	echoRequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "muzz_echo_requests_total",
		Help: "Total number of Echo requests",
	})

	// echoRequestDuration is a histogram that tracks the duration of Echo requests in seconds.
	echoRequestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "muzz_echo_request_duration_seconds",
		Help:    "Duration of Echo requests in seconds",
		Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2, 5},
	})

	// echoRequestsFailedTotal is a counter that tracks the total number of failed Echo requests.
	echoRequestsFailedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "muzz_echo_requests_failed_total",
		Help: "Total number of failed Echo requests",
	})

	// echoRequestDurationSummary is a summary vector that tracks the request latency percentiles for the Echo method.
	echoRequestDurationSummary = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "muzz_echo_request_duration_seconds",
			Help:       "Summary of Echo request durations in seconds",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"method"},
	)

	// echoRequestsCounter is a counter vector that tracks the total number of Echo requests,
	// with labels for the request method and response status (success or error).
	echoRequestsCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "muzz_echo_requests_total",
			Help: "Total number of Echo requests",
		},
		[]string{"method", "status"},
	)
)