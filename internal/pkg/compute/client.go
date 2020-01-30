// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shurcooL/graphql"
	"github.com/sirupsen/logrus"
	"github.com/sylabs/compute-cli/internal/pkg/schema"
)

const userAgent = "compute-cli/0.1"

type Client struct {
	*graphql.Client
}

func NewClient(serverURL string) *Client {
	endpoint := fmt.Sprintf("%s/graphql", serverURL)
	logrus.Debugf("Creating graphql client for: %s", endpoint)
	return &Client{graphql.NewClient(endpoint, &http.Client{
		Transport: setUserAgent(http.DefaultTransport, userAgent),
	})}
}

func (c *Client) Create(ctx context.Context, spec schema.WorkflowSpec) (Workflow, error) {
	variables := map[string]interface{}{
		"workflowSpec": spec,
	}

	m := struct {
		schema.Workflow `graphql:"createWorkflow(spec: $workflowSpec)"`
	}{}

	err := c.Mutate(ctx, &m, variables)
	if err != nil {
		return Workflow{}, fmt.Errorf("while creating workflow: %w", err)
	}

	return convertWorkflow(m.Workflow), nil
}

func (c *Client) Delete(ctx context.Context, id string) (Workflow, error) {
	variables := map[string]interface{}{
		"id": graphql.ID(id),
	}

	m := struct {
		schema.Workflow `graphql:"deleteWorkflow(id: $id)"`
	}{}

	// TODO: gracefully catch case where the workflow does not exist
	err := c.Mutate(ctx, &m, variables)
	if err != nil {
		return Workflow{}, fmt.Errorf("while deleting workflow: %w", err)
	}

	return convertWorkflow(m.Workflow), nil
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
