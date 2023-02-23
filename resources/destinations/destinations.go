package destinations

import (
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

func DestinationsTable() *schema.Table {
	return &schema.Table{
		Name:      "rudderstack_destinations",
		Resolver:  fetchDestinations,
		Transform: transformers.TransformWithStruct(&client.Destination{}),
	}
}

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
