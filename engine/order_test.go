package engine

import (
	"testing"
)

func TestNewOrder(t *testing.T) {
	t.Log(NewOrder("b1", Sell, "5.0", "7000.0"))
}

func TestToJSON(t *testing.T) {
	order := NewOrder("b1", Buy, "5.0", "7000.0")

	result, _ := order.ToJSON()
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
		{"{\"price\":0,\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for amount key"},
		{"{\"amount\":5,\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for price key"},
		{"{\"amount\":5,\"price\":7000,\"type\":\"buy\"}", "err", "Check for id key"},
		{"{\"amount\":5,\"price\":7000,\"id\":\"b1\"}", "err", "Check for type key"},

		{"{\"amount\":5,\"price\":7000,\"id\":\"b1\",\"type\":\"random\"}", "err", "Check for valid type"},
		{"{\"amount\":0,\"price\":7000,\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for valid amount"},
		{"{\"amount\":5,\"price\":0,\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for valid price"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			order := &Order{}
			err := order.FromJSON([]byte(tt.input))
			if tt.err == "" && err == nil {
				t.Log("Successfully detecting error")
			} else if tt.err != "" && err != nil {
				t.Log("Successful detection of json")
			} else {
				t.Fatal(tt.message)
			}
		})
	}
}

func TestOrderString(t *testing.T) {
	var tests = []struct {
		input  *Order
		output string
	}{
		{
			NewOrder("b1", Buy, "5.0", "7000.0"),
			`"b1":
	side: buy
	quantity: 5
	price: 7000
`},
		{
			NewOrder("s1", Sell, "5.124", "9000.0"),
			`"s1":
	side: sell
	quantity: 5.124
	price: 9000
`},
	}
	for _, tt := range tests {
		o := tt.input.String()
		if tt.output != o {
			t.Fatalf("Book prints incorrect (have: \n%s, \nwant: \n%s\n)", o, tt.output)
		}
	}
}
