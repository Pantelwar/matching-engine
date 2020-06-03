package engine

import (
	"fmt"
	"testing"
)

func TestNewOrder(t *testing.T) {
	t.Log(NewOrder("b1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")))
}

func TestToJSON(t *testing.T) {
	var tests = []struct {
		input  *Order
		output string
	}{
		{NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")), `{"type":"buy","id":"b1","amount":"5.0","price":"7000.0"}`},
		{NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")), `{"type":"sell","id":"s1","amount":"5.0","price":"7000.0"}`},
	}

	for _, tt := range tests {
		result, _ := tt.input.ToJSON()

		if string(result) != tt.output {
			t.Fatalf("Unable to marshal json (have: %s, want: %s)\n", string(result), tt.output)
		}
	}
}

func TestFromJSON(t *testing.T) {
	var tests = []struct {
		input   string
		err     string
		message string
	}{
		{"{\"amount\":\"5.0\",\"price\":\"7000.0\",\"id\":\"b1\",\"type\":\"buy\"}", "", "JSON should be approved"},

		{"{}", "err", "Empty JSON should not be passed"},
		{"{\"price\":\"0.0\",\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for amount key"},
		{"{\"amount\":\"5.0\",\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for price key"},
		{"{\"amount\":\"5.0\",\"price\":\"7000.0\",\"type\":\"buy\"}", "err", "Check for id key"},
		{"{\"amount\":\"5.0\",\"price\":\"7000.0\",\"id\":\"b1\"}", "err", "Check for type key"},

		{"{\"amount\":\"0.0\",\"price\":\"7000.0\",\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for valid amount"},
		{"{\"amount\":\"5.0\",\"price\":\"0.0\",\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for valid price"},

		{"{\"amount\":\"5.0\",\"price\":\"7000.0\",\"id\":\"b1\",\"type\":\"random\"}", "err", "Check for valid type"},
		{"{\"amount\":\"random\",\"price\":\"0.0\",\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for valid amount"},
		{"{\"amount\":\"0.0\",\"price\":\"random\",\"id\":\"b1\",\"type\":\"buy\"}", "err", "Check for valid price"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			order := &Order{}
			err := order.FromJSON([]byte(tt.input))
			fmt.Println("error:", err)
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
			NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
			`"b1":
	side: buy
	quantity: 5
	price: 7000
`},
		{
			NewOrder("s1", Sell, DecimalBig("5.124"), DecimalBig("9000.0")),
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
