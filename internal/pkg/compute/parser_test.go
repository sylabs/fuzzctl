// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package compute

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/sylabs/compute-cli/internal/pkg/schema"
)

func TestParseSpec(t *testing.T) {
	testDataDir := "./testdata/parser"

	tests := []struct {
		name    string
		input   string
		golden  string
		wantErr bool
	}{
		{
			name:    "Single Job",
			input:   "single.yaml",
			golden:  "single.json",
			wantErr: false,
		},
		{
			name:    "Multiple Jobs",
			input:   "multiple.yaml",
			golden:  "multiple.json",
			wantErr: false,
		},
		{
			name:    "Dependent Jobs",
			input:   "dependent.yaml",
			golden:  "dependent.json",
			wantErr: false,
		},
		{
			name:    "Volumes",
			input:   "volume.yaml",
			golden:  "volume.json",
			wantErr: false,
		},
		{
			name:    "Unknown Field",
			input:   "unknownfield.yaml",
			golden:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := ioutil.ReadFile(filepath.Join(testDataDir, "input", tt.input))
			if err != nil {
				t.Fatal(err)
			}

			spec, err := ParseSpec(d)
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
