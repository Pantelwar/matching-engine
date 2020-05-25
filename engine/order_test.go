package engine

import (
	"testing"
)

func TestNewOrder(t *testing.T) {
	t.Log(NewOrder("b1", Sell, 5.0, 7000.0))
}

func TestToJSON(t *testing.T) {
	data := NewOrder("b1", Buy, 5.0, 7000.0)

	result, _ := data.ToJSON()
	if string(result) != "{\"amount\":5,\"price\":7000,\"id\":\"b1\",\"type\":\"buy\"}" {
		t.Fatal("Result should be: {\"amount\":5,\"price\":7000,\"id\":\"b1\",\"type\":\"buy\"}, got: " + string(result))
	}
}

func TestFromJSON(t *testing.T) {
	var tests = []struct {
		input   string
		err     string
		message string
	}{
		{"{\"amount\":5,\"price\":7000,\"id\":\"b1\",\"type\":\"buy\"}", "", "JSON should be approved"},

		{"{}", "err", "Empty JSON should not be passed"},
		{"{\"price\":0,\"id\":\"b1\",\"type\":\"buy\"}", "err", "check for amount key"},
		{"{\"amount\":5,\"id\":\"b1\",\"type\":\"buy\"}", "err", "check for price key"},
		{"{\"amount\":5,\"price\":7000,\"type\":\"buy\"}", "err", "check for id key"},
		{"{\"amount\":5,\"price\":7000,\"id\":\"b1\"}", "err", "check for type key"},

		{"{\"amount\":5,\"price\":7000,\"id\":\"b1\",\"type\":\"random\"}", "err", "check for valid type"},
		{"{\"amount\":0,\"price\":7000,\"id\":\"b1\",\"type\":\"buy\"}", "err", "check for valid amount"},
		{"{\"amount\":5,\"price\":0,\"id\":\"b1\",\"type\":\"buy\"}", "err", "check for valid price"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			data1 := &Order{}
			err := data1.FromJSON([]byte(tt.input))
			if tt.err == "" && err == nil {
				t.Log("Successful Detection")
			} else if tt.err != "" && err != nil {
				t.Log("Successful Detection")
			} else {
				t.Fatal(tt.message)
			}
		})
	}
}
