package sources

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/rudderlabs/rudder-api-go/client"
)

func SourcesTable() *schema.Table {
	return &schema.Table{
		Name:      "rudderstack_sources",
		Resolver:  fetchSources,
		Transform: transformers.TransformWithStruct(&client.Source{}),
	}
}
