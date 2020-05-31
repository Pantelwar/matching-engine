package engine

import (
	"encoding/json"
	"testing"
)

type testSide struct {
	Type Side `json:"type"`
}

func TestSideUnmarshal(t *testing.T) {
	var tests = []struct {
		input   string
		err     string
		message string
	}{
		{"{\"type\":\"buy\"}", "", "JSON should be approved"},
		{"{\"type\":\"sell\"}", "", "JSON should be approved"},
		{"{}", "err", "Empty JSON should not be passed"},
		{"\"type\":\"random\"}", "err", "type should be either buy or sell"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			side := testSide{}
			err := json.Unmarshal([]byte(tt.input), &side)
			if tt.err == "" && err == nil {
				t.Log("Successfully detecting error")
			} else if tt.err != "" && err != nil {
				t.Log("Successful detection of json")
			} else {
				if tt.err != "" && side.Type == "" {
					t.Log("Successful detecting of empty json")
				} else {
					t.Fatal(tt.message)
				}
			}
		})
	}
}

func TestSideMarshal(t *testing.T) {
	var tests = []struct {
		input  Side
		output string
	}{
		{Buy, "\"buy\""},
		{Sell, "\"sell\""},
	}

	for _, tt := range tests {
		output, _ := json.Marshal(tt.input)
		if string(output) != tt.output {
			t.Fatalf("Marshal error: (have: %s, want: %s\n", string(output), tt.output)
		}
	}
}

func TestSideString(t *testing.T) {
	var tests = []struct {
		input  Side
		output string
	}{
		{Buy, "buy"},
		{Sell, "sell"},
	}

	for _, tt := range tests {
		output := tt.input.String()
		if string(output) != tt.output {
			t.Fatalf("String error: (have: %s, want: %s)\n", string(output), tt.output)
		}
	}
}
