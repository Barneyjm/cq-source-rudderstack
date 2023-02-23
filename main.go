package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/rudderlabs/cq-source-rudderstack/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
