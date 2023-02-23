package resources

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rudderlabs/rudder-api-go/client"
)

func setupClient() (*client.Client, error) {
	accessToken := os.Getenv("RUDDERSTACK_ACCESS_TOKEN")
	if accessToken == "" {
		return nil, fmt.Errorf("no access token in specified. Please provide one through the RUDDERSTACK_ACCESS_TOKEN environmental variable")
	}

	baseURL := os.Getenv("RUDDERSTACK_API_URL")
	if baseURL == "" {
		baseURL = client.BASE_URL_V2
	}

	return client.New(accessToken, client.WithBaseURL(baseURL))
}

func SourcesTable() *schema.Table {
	return &schema.Table{
		Name:      "rudderstack_sources",
		Resolver:  fetchSources,
		Transform: transformers.TransformWithStruct(&client.Source{}),
	}
}

func fetchSources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client, _ := setupClient()
	sourcesPage, err := client.Sources.List(context.Background())
	if err != nil {
		return err
	}

	for sourcesPage != nil {
		res <- sourcesPage.Sources
		// sources = append(sources, sourcesPage.Sources...)
		sourcesPage, err = client.Sources.Next(context.Background(), sourcesPage.Paging)
		if err != nil {
			return err
		}
	}

	return nil
}

// func fetchSources() ([]client.Source, error) {
// 	var sources []client.Source
// 	sourcesPage, err := cl.Sources.List(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	for sourcesPage != nil {
// 		sources = append(sources, sourcesPage.Sources...)
// 		sourcesPage, err = cl.Sources.Next(context.Background(), sourcesPage.Paging)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return sources, nil
// }

// func getAPIDestinations() ([]client.Destination, error) {
// 	var destinations []client.Destination
// 	destinationsPage, err := cl.Destinations.List(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	for destinationsPage != nil {
// 		destinations = append(destinations, destinationsPage.Destinations...)
// 		destinationsPage, err = cl.Destinations.Next(context.Background(), destinationsPage.Paging)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return destinations, nil
// }

// func getAPIConnections() ([]client.Connection, error) {
// 	var connections []client.Connection
// 	connectionsPage, err := cl.Connections.List(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	for connectionsPage != nil {
// 		connections = append(connections, connectionsPage.Connections...)
// 		connectionsPage, err = cl.Connections.Next(context.Background(), connectionsPage.Paging)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return connections, nil
// }
