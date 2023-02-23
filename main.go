package main

import (
	"github.com/barneyjm/cq-source-rudderstack/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
