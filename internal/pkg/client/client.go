// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shurcooL/graphql"
	"github.com/sylabs/fuzzctl/internal/pkg/schema"
)

// Client is a Fuzzball GraphQL client.
type Client struct {
	*graphql.Client

	httpClient *http.Client
	userAgent  string
}

// OptHTTPClient overrides the default HTTP client.
func OptHTTPClient(hc *http.Client) func(c *Client) error {
	return func(c *Client) error {
		c.httpClient = hc
		return nil
	}
}

// OptUserAgent overrides the default HTTP user agent.
func OptUserAgent(s string) func(c *Client) error {
	return func(c *Client) error {
		c.userAgent = s
		return nil
	}
}

// NewClient creates a GraphQL client targeting the specified GraphQL endpoint, using options opts.
func NewClient(url string, opts ...func(c *Client) error) (*Client, error) {
	// Set up default client.
	c := Client{
		httpClient: http.DefaultClient,
	}

	// Apply options to client.
	for _, o := range opts {
		if err := o(&c); err != nil {
			return nil, err
		}
	}

	// Apply user agent, if set.
	if ua := c.userAgent; ua != "" {
		c.httpClient.Transport = setUserAgent(c.httpClient.Transport, ua)
	}

	// Set GraphQL client.
	c.Client = graphql.NewClient(url, c.httpClient)

	return &c, nil
}

func (c *Client) Create(ctx context.Context, spec schema.WorkflowSpec) (Workflow, error) {
	variables := map[string]interface{}{
		"workflowSpec": spec,
	}

	m := struct {
		Workflow struct {
			ID   string
			Name string
		} `graphql:"createWorkflow(spec: $workflowSpec)"`
	}{}

	err := c.Mutate(ctx, &m, variables)
	if err != nil {
		return Workflow{}, fmt.Errorf("while creating workflow: %w", err)
	}

	return Workflow{ID: m.Workflow.ID, Name: m.Workflow.Name}, nil
}

func (c *Client) Delete(ctx context.Context, id string) (Workflow, error) {
	variables := map[string]interface{}{
		"id": graphql.ID(id),
	}

	m := struct {
		Workflow struct {
			ID   string
			Name string
		} `graphql:"deleteWorkflow(id: $id)"`
	}{}

	// TODO: gracefully catch case where the workflow does not exist
	err := c.Mutate(ctx, &m, variables)
	if err != nil {
		return Workflow{}, fmt.Errorf("while deleting workflow: %w", err)
	}

	return Workflow{ID: m.Workflow.ID, Name: m.Workflow.Name}, nil
}

func (c *Client) Info(ctx context.Context, id string) (Workflow, error) {
	variables := map[string]interface{}{
		"id": graphql.ID(id),
	}

	q := struct {
		schema.Workflow `graphql:"workflow(id: $id)"`
	}{}

	// TODO: gracefully catch case where the workflow does not exist
	err := c.Query(ctx, &q, variables)
	if err != nil {
		return Workflow{}, fmt.Errorf("while getting workflow state: %w", err)
	}

	return convertWorkflow(q.Workflow), nil
}

func (c *Client) List(ctx context.Context) ([]Workflow, error) {
	q := struct {
		schema.Viewer `graphql:"viewer"`
	}{}

	err := c.Query(ctx, &q, nil)
	if err != nil {
		return nil, fmt.Errorf("while getting workflow state: %w", err)
	}

	var wfs []Workflow
	for _, w := range q.Viewer.Workflows.Edges {
		wfs = append(wfs, convertWorkflow(w.Node))
	}

	return wfs, nil
}

// ServerBuildInfo retrieves build information about the server.
func (c *Client) ServerBuildInfo(ctx context.Context) (BuildInfo, error) {
	q := struct {
		schema.BuildInfo `graphql:"serverBuildInfo"`
	}{}

	if err := c.Query(ctx, &q, nil); err != nil {
		return BuildInfo{}, err
	}
	return convertBuildInfo(q.BuildInfo), nil
}
