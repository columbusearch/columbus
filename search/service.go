package main

import (
	"fmt"
	"net/http"
	"net/url"

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
	Url url.URL
}

// Search requests from typesense a query and sends back an arrray of Results of the first n results
func (s *SearchService) Search(query string, n int) ([]Result, error) {
	searchResult, err := s.client.Collection("columbus").Documents().Search(
		&api.SearchCollectionParams{
			Q:       query,
			QueryBy: "url,content",
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
		prasedURL, err := url.Parse(document["url"].(string))

		if err != nil {
			s.metricsServer.IncErrorCount(query)
			s.log.Error(err)
		}

		results = append(results, Result{
			Url: *prasedURL,
		})
	}

	if len(results) == 0 {
		s.log.Infof("No results for query \"%s\"", query)
	} else {
		s.log.Tracef("Found %d results for query \"%s\"", len(results), query)
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
