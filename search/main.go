package main

import (
	"fmt"
	"net/http"

	"github.com/columbusearch/columbus/search/metrics"
)

func main() {
	metricsServer := metrics.NewMetricsServer()
	search := NewSearchService("http://localhost:8108", "xyz", metricsServer)
	port := "8080"

	http.Handle("/metrics", search.metricsServer.GetHandler())
	http.Handle("/health", search.metricsServer.PrometheusMetricsMiddleware(http.HandlerFunc(search.HandleHealth)))
	http.Handle("/search", search.metricsServer.PrometheusMetricsMiddleware(http.HandlerFunc(search.HandleSearch)))

	search.log.Infof("Columbus search server starting on port %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		search.log.Fatalf("Error starting search server: %s", err)
	}
}
