// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import (
	"context"
	"fmt"

	"github.com/shurcooL/graphql"
	"github.com/sirupsen/logrus"
	"github.com/sylabs/compute-cli/internal/pkg/schema"
	"golang.org/x/oauth2"
)

const (
	userAgent = "compute-cli/0.1"
)

type Client struct {
	*graphql.Client
}

func NewClient(ctx context.Context, ts oauth2.TokenSource, serverURL string) *Client {
	endpoint := fmt.Sprintf("%s/graphql", serverURL)
	logrus.Debugf("Creating graphql client for: %s", endpoint)
	c := oauth2.NewClient(ctx, ts)
	c.Transport = setUserAgent(c.Transport, userAgent)
	return &Client{graphql.NewClient(endpoint, c)}
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

func (c *Client) ServerInfo(ctx context.Context) (ServerInfo, error) {
	q := struct {
		ServerInfo `graphql:"systemInfo"`
	}{}

	err := c.Query(ctx, &q, nil)
	if err != nil {
		return ServerInfo{}, fmt.Errorf("while getting server information: %w", err)
	}

	return q.ServerInfo, nil
}
