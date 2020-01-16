// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import (
	"context"
	"fmt"

	"github.com/shurcooL/graphql"
	"github.com/sirupsen/logrus"
)

type Client struct {
	*graphql.Client
}

func NewClient(serverURL string) *Client {
	endpoint := fmt.Sprintf("%s/graphql", serverURL)
	logrus.Debugf("Creating graphql client for: %s", endpoint)
	return &Client{graphql.NewClient(endpoint, nil)}
}

func (c *Client) Create(ctx context.Context, name string) (*Workflow, error) {
	variables := map[string]interface{}{
		"name": graphql.String(name),
	}

	cwf := struct {
		Workflow `graphql:"createWorkflow(name: $name)"`
	}{}

	err := c.Mutate(ctx, &cwf, variables)
	if err != nil {
		return nil, fmt.Errorf("while creating workflow: %w", err)
	}

	return &cwf.Workflow, nil
}

func (c *Client) Delete(ctx context.Context, id string) (*Workflow, error) {

	variables := map[string]interface{}{
		"id": graphql.ID(id),
	}

	dwf := struct {
		Workflow `graphql:"deleteWorkflow(id: $id)"`
	}{}

	// TODO: gracefully catch case where the workflow does not exist
	err := c.Mutate(ctx, &dwf, variables)
	if err != nil {
		return nil, fmt.Errorf("while deleting workflow: %w", err)
	}

	return &dwf.Workflow, nil
}

func (c *Client) Info(ctx context.Context, id string) (*Workflow, error) {
	variables := map[string]interface{}{
		"id": graphql.ID(id),
	}

	iwf := struct {
		Workflow `graphql:"workflow(id: $id)"`
	}{}

	// TODO: gracefully catch case where the workflow does not exist
	err := c.Query(ctx, &iwf, variables)
	if err != nil {
		return nil, fmt.Errorf("while getting workflow state: %w", err)
	}

	return &iwf.Workflow, nil
}

func (c *Client) List(ctx context.Context) ([]Workflow, error) {
	// TODO: implement when endpoint is up
	return nil, nil
}
