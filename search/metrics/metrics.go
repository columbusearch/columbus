package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsServer struct {
	handler             http.Handler
	queryCount          *prometheus.CounterVec
	errorCount          *prometheus.CounterVec
	responseStatusCount *prometheus.CounterVec
	responseTime        *prometheus.HistogramVec
}

func NewMetricsServer() *MetricsServer {
	registry := prometheus.NewRegistry()
	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	queryCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "columbus_search_query_count",
		Help: "Number of queries received",
	}, []string{"query"})
	registry.MustRegister(queryCount)

	errorCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "columbus_search_error_count",
		Help: "Number of errors",
	}, []string{"query"})
	registry.MustRegister(errorCount)

	responseStatusCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "columbus_search_response_status_count",
		Help: "Number of HTTP responses",
	}, []string{"status"})
	registry.MustRegister(responseStatusCount)

	responseTime := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "columbus_search_response_time",
		Help: "Time taken to respond to a query",
	}, []string{"query"})
	registry.MustRegister(responseTime)

	return &MetricsServer{
		handler:             promhttp.HandlerFor(registry, promhttp.HandlerOpts{}),
		queryCount:          queryCount,
		responseStatusCount: responseStatusCount,
		responseTime:        responseTime,
	}
}

func (m *MetricsServer) GetHandler() http.Handler {
	return m.handler
}

// IncQueryCount increments the query count for the given query
func (m *MetricsServer) IncQueryCount(query string) {
	m.queryCount.WithLabelValues(query).Inc()
}

// IncErrorCount increments the error count for the given query
func (m *MetricsServer) IncErrorCount(query string) {
	m.errorCount.WithLabelValues(query).Inc()
}

// IncResponseStatusCount increments the error count for the given query
func (m *MetricsServer) IncResponseStatusCount(status string) {
	m.responseStatusCount.WithLabelValues(status).Inc()
}

// ObserveResponseTime records the response time for the given query
func (m *MetricsServer) ObserveResponseTime(query string, duration float64) {
	m.responseTime.WithLabelValues(query).Observe(duration)
}

func (m *MetricsServer) PrometheusMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")

		timer := prometheus.NewTimer(m.responseTime.WithLabelValues(query))

		next.ServeHTTP(w, r)
		statusCode := w.Header().Get("Status")

		m.IncResponseStatusCount(statusCode)
		m.IncQueryCount(query)

		timer.ObserveDuration()
	})
}
