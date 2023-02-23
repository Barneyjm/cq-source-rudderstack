package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/rudderlabs/rudder-api-go/client"
)

type Client struct {
	Logger      zerolog.Logger
	RudderStack *client.Client
}

func (c *Client) ID() string {
	return "rudderstack"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	accessToken := os.Getenv("RUDDERSTACK_ACCESS_TOKEN")
	if accessToken == "" {
		return nil, fmt.Errorf("no access token in specified. Please provide one through the RUDDERSTACK_ACCESS_TOKEN environmental variable")
	}

	baseURL := os.Getenv("RUDDERSTACK_API_URL")
	if baseURL == "" {
		baseURL = client.BASE_URL_V2
	}

	c, err := client.New(accessToken, client.WithBaseURL(baseURL))
	if err != nil {
		return nil, err
	}

	return &Client{
		Logger:      logger,
		RudderStack: c,
	}, nil
}
