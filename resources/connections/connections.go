package connections

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	rudderstackClient "github.com/rudderlabs/rudder-api-go/client"
)

func ConnectionsTable() *schema.Table {
	return &schema.Table{
		Name:      "rudderstack_connections",
		Resolver:  fetchConnections,
		Transform: transformers.TransformWithStruct(&rudderstackClient.Connection{}),
	}
}
