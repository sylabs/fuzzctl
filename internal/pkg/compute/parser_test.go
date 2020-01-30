// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/sylabs/compute-cli/internal/pkg/schema"
)

func TestParseSpec(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/parserdata/singlejob.yaml")
	if err != nil {
		t.Error(err)
	}

	t.Logf(string(data))

	spec, err := ParseSpec(data)
	if err != nil {
		t.Error(err)
	}

	correct := &schema.WorkflowSpec{
		Name: "mvp2",
		Jobs: []schema.JobSpec{
			schema.JobSpec{
				Name:    "date",
				Image:   "library://alpine:latest",
				Command: []string{"date"},
			},
		},
	}

	if !reflect.DeepEqual(spec, correct) {
		t.Errorf("Workflow spec incorrectly parsed:\nCorrect:\n%v\nParsed:\n%v", correct, spec)
	}
}
