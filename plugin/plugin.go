package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/rudderlabs/cq-source-rudderstack/client"
	"github.com/rudderlabs/cq-source-rudderstack/resources"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"rudderlabs-rudderstack",
		Version,
		schema.Tables{
			resources.SourcesTable(),
		},
		client.New,
	)
}
