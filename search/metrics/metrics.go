package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsServer struct {
	handler      http.Handler
	queryCount   *prometheus.CounterVec
	errorCount   *prometheus.CounterVec
	responseTime *prometheus.HistogramVec
}

func NewMetricsServer() *MetricsServer {
	registry := prometheus.NewRegistry()
	registry.MustRegister(collectors.NewGoCollector())
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	queryCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "search_query_count",
		Help: "Number of queries received",
	}, []string{"query"})
	registry.MustRegister(queryCount)

	errorCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "search_error_count",
		Help: "Number of errors returned",
	}, []string{"query"})
	registry.MustRegister(errorCount)

	responseTime := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "search_response_time",
		Help: "Time taken to respond to a query",
	}, []string{"query"})
	registry.MustRegister(responseTime)

	return &MetricsServer{
		handler:      promhttp.HandlerFor(registry, promhttp.HandlerOpts{}),
		queryCount:   queryCount,
		errorCount:   errorCount,
		responseTime: responseTime,
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

// ObserveResponseTime records the response time for the given query
func (m *MetricsServer) ObserveResponseTime(query string, duration float64) {
	m.responseTime.WithLabelValues(query).Observe(duration)
}
