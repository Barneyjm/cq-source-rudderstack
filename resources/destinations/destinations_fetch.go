package destinations

import (
	"context"

	"github.com/barneyjm/cq-source-rudderstack/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	destinationsPage, err := c.RudderStack.Destinations.List(ctx)
	if err != nil {
		return err
	}

	for destinationsPage != nil {
		res <- destinationsPage.Destinations
		destinationsPage, err = c.RudderStack.Destinations.Next(ctx, destinationsPage.Paging)
		if err != nil {
			return err
		}
	}

	return nil
}
