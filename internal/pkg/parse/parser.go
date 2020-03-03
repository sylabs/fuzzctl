// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package parse

import (
	"fmt"

	"github.com/sylabs/fuzzctl/internal/pkg/schema"
)

const specVersion = "0.1"

type Decoder interface {
	Decode(v interface{}) error
}

type Parser struct {
	d Decoder
}

func New(d Decoder) Parser {
	return Parser{d}
}

// ParseWorkflowSpec decodes a specification for a workflow
// into a structure that can be used by a GraphQL client
// as an input object.
func (p Parser) ParseWorkflowSpec() (schema.WorkflowSpec, error) {
	s := struct {
		Version  string
		Workflow schema.WorkflowSpec
	}{}

	if err := p.d.Decode(&s); err != nil {
		return schema.WorkflowSpec{}, err
	}

	if s.Version != specVersion {
		return schema.WorkflowSpec{}, fmt.Errorf("unknown spec version: %s. Expected: %s", s.Version, specVersion)
	}

	return s.Workflow, nil
}
