package sources

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client, _ := setupClient()
	sourcesPage, err := client.Sources.List(context.Background())
	if err != nil {
		return err
	}

	for sourcesPage != nil {
		res <- sourcesPage.Sources
		sourcesPage, err = client.Sources.Next(context.Background(), sourcesPage.Paging)
		if err != nil {
			return err
		}
	}

	return nil
}
