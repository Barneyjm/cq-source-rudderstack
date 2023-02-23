package destinations

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	rudderstackClient "github.com/rudderlabs/rudder-api-go/client"
)

func DestinationsTable() *schema.Table {
	return &schema.Table{
		Name:      "rudderstack_destinations",
		Resolver:  fetchDestinations,
		Transform: transformers.TransformWithStruct(&rudderstackClient.Destination{}),
	}
}
