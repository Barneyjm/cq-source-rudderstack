package connections

import (
	"context"

	"github.com/barneyjm/cq-source-rudderstack/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	connectionsPage, err := c.RudderStack.Connections.List(ctx)
	if err != nil {
		return err
	}

	for connectionsPage != nil {
		res <- connectionsPage.Connections
		connectionsPage, err = c.RudderStack.Connections.Next(ctx, connectionsPage.Paging)
		if err != nil {
			return err
		}
	}

	return nil
}
