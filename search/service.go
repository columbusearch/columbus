package main

import (
	"fmt"
	"net/http"

	"github.com/itamadev/columbus/search/metrics"
	log "github.com/sirupsen/logrus"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

type SearchService struct {
	client        *typesense.Client
	metricsServer *metrics.MetricsServer
	log           *log.Entry
}

type Result struct {
	Content string
	Url     string
}

// Search requests from typesense a query and sends back an arrray of Results of the first n results
func (s *SearchService) Search(query string, n int) ([]Result, error) {
	searchResult, err := s.client.Collection("columbus").Documents().Search(
		&api.SearchCollectionParams{
			Q:       query,
			PerPage: &n,
		},
	)

	if err != nil {
		s.metricsServer.IncErrorCount(query)
		return nil, err
	}

	results := make([]Result, 0)
	for _, hit := range *searchResult.Hits {
		document := *hit.Document
		results = append(results, Result{
			Content: document["content"].(string),
			Url:     document["url"].(string),
		})
	}
	return results, nil
}

// HandleSearch handles the search request and returns the results
func (s *SearchService) HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	n := 10

	results, err := s.Search(query, n)
	if err != nil {
		s.log.Errorf("Error searching for \"%s\": %s", query, err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, results)
}

// Health returns the health of the service
func (s *SearchService) HandleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Service is healthy")
}

// NewSearchService creates a new SearchService
func NewSearchService(typesenseUrl string, typesenseKey string, metricsServer *metrics.MetricsServer) *SearchService {
	client := typesense.NewClient(
		typesense.WithAPIKey(typesenseKey),
		typesense.WithServer(typesenseUrl),
	)
	serviceLog := log.NewEntry(log.StandardLogger())
	return &SearchService{
		client:        client,
		metricsServer: metricsServer,
		log:           serviceLog,
	}
}
