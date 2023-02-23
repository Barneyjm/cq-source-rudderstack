package destinations

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client, _ := setupClient()
	destinationsPage, err := client.Destinations.List(context.Background())
	if err != nil {
		return err
	}

	for destinationsPage != nil {
		res <- destinationsPage.Destinations
		destinationsPage, err = client.Destinations.Next(context.Background(), destinationsPage.Paging)
		if err != nil {
			return err
		}
	}

	return nil
}
