// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package parse

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sylabs/compute-cli/internal/pkg/schema"
	"gopkg.in/yaml.v3"
)

func TestParseYAML(t *testing.T) {
	testDataDir := "./testdata/parser"

	tests := []struct {
		name    string
		input   string
		golden  string
		wantErr bool
	}{
		{
			name:    "Single Job YAML",
			input:   "single.yaml",
			golden:  "single.json",
			wantErr: false,
		},
		{
			name:    "Single Job JSON",
			input:   "single.json",
			golden:  "single.json",
			wantErr: false,
		},
		{
			name:    "Multiple Jobs YAML",
			input:   "multiple.yaml",
			golden:  "multiple.json",
			wantErr: false,
		},
		{
			name:    "Multiple Jobs JSON",
			input:   "multiple.json",
			golden:  "multiple.json",
			wantErr: false,
		},
		{
			name:    "Dependent Jobs YAML",
			input:   "dependent.yaml",
			golden:  "dependent.json",
			wantErr: false,
		},
		{
			name:    "Dependent Jobs JSON",
			input:   "dependent.json",
			golden:  "dependent.json",
			wantErr: false,
		},
		{
			name:    "Volumes YAML",
			input:   "volume.yaml",
			golden:  "volume.json",
			wantErr: false,
		},
		{
			name:    "Volumes JSON",
			input:   "volume.json",
			golden:  "volume.json",
			wantErr: false,
		},
		{
			name:    "Unknown Field YAML",
			input:   "unknownfield.yaml",
			golden:  "",
			wantErr: true,
		},
		{
			name:    "Unknown Field JSON",
			input:   "unknownfield.json",
			golden:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			specFile, err := os.Open(filepath.Join(testDataDir, "input", tt.input))
			if err != nil {
				logrus.Fatal(err)
			}

			d := yaml.NewDecoder(specFile)
			d.KnownFields(true)

			p := New(d)
			spec, err := p.ParseWorkflowSpec()
			if (err != nil) != tt.wantErr {
				t.Fatalf("got err %v, wantErr %v", err, tt.wantErr)
			}

			// stop if we got the error we wanted
			if tt.wantErr {
				return
			}

			gd, err := ioutil.ReadFile(filepath.Join(testDataDir, "golden", tt.golden))
			if err != nil {
				t.Fatal(err)
			}

			var golden schema.WorkflowSpec
			if err := json.Unmarshal(gd, &golden); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(spec, golden) {
				t.Fatalf("Workflow spec incorrectly parsed:\nGolden:\n%v\nParsed:\n%v", golden, spec)
			}
		})
	}
}

func TestParseJSON(t *testing.T) {
	testDataDir := "./testdata/parser"

	tests := []struct {
		name    string
		input   string
		golden  string
		wantErr bool
	}{
		{
			name:    "Single Job JSON",
			input:   "single.json",
			golden:  "single.json",
			wantErr: false,
		},
		{
			name:    "Multiple Jobs JSON",
			input:   "multiple.json",
			golden:  "multiple.json",
			wantErr: false,
		},
		{
			name:    "Dependent Jobs JSON",
			input:   "dependent.json",
			golden:  "dependent.json",
			wantErr: false,
		},
		{
			name:    "Volumes JSON",
			input:   "volume.json",
			golden:  "volume.json",
			wantErr: false,
		},
		{
			name:    "Unknown Field JSON",
			input:   "unknownfield.json",
			golden:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			specFile, err := os.Open(filepath.Join(testDataDir, "input", tt.input))
			if err != nil {
				logrus.Fatal(err)
			}

			d := json.NewDecoder(specFile)
			d.DisallowUnknownFields()

			p := New(d)
			spec, err := p.ParseWorkflowSpec()
			if (err != nil) != tt.wantErr {
				t.Fatalf("got err %v, wantErr %v", err, tt.wantErr)
			}

			// stop if we got the error we wanted
			if tt.wantErr {
				return
			}

			gd, err := ioutil.ReadFile(filepath.Join(testDataDir, "golden", tt.golden))
			if err != nil {
				t.Fatal(err)
			}

			var golden schema.WorkflowSpec
			if err := json.Unmarshal(gd, &golden); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(spec, golden) {
				t.Fatalf("Workflow spec incorrectly parsed:\nGolden:\n%v\nParsed:\n%v", golden, spec)
			}
		})
	}
}
