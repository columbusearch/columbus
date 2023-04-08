package main

import (
	"github.com/itamadev/columbus/search/metrics"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

type SearchService struct {
	client        *typesense.Client
	metricsServer *metrics.MetricsServer
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

func NewSearchService(typesenseUrl string, typesenseKey string, metricsServer *metrics.MetricsServer) *SearchService {
	client := typesense.NewClient(
		typesense.WithAPIKey(typesenseKey),
		typesense.WithServer(typesenseUrl),
	)
	return &SearchService{
		client:        client,
		metricsServer: metricsServer,
	}
}
