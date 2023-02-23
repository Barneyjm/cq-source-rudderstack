package plugin

import (
	"github.com/barneyjm/cq-source-rudderstack/client"
	"github.com/barneyjm/cq-source-rudderstack/resources/connections"
	"github.com/barneyjm/cq-source-rudderstack/resources/destinations"
	"github.com/barneyjm/cq-source-rudderstack/resources/sources"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"rudderlabs-rudderstack",
		Version,
		schema.Tables{
			sources.SourcesTable(),
			destinations.DestinationsTable(),
			connections.ConnectionsTable(),
		},
		client.New,
	)
}
