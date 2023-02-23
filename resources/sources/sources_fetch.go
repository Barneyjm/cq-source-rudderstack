package sources

import (
	"context"

	"github.com/barneyjm/cq-source-rudderstack/client"

	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	sourcesPage, err := c.RudderStack.Sources.List(ctx)
	if err != nil {
		return err
	}

	for sourcesPage != nil {
		res <- sourcesPage.Sources
		sourcesPage, err = c.RudderStack.Sources.Next(ctx, sourcesPage.Paging)
		if err != nil {
			return err
		}
	}

	return nil
}
