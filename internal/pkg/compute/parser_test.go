// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import (
	"io/ioutil"
	"reflect"
	"testing"
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

	correct := &WorkflowSpec{
		Name: "mvp2",
		Jobs: []JobSpec{
			JobSpec{
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
