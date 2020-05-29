package engine

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewOrderType(t *testing.T) {
	t.Log(NewOrderType("sell"))
}

func TestAddOrderInQueue(t *testing.T) {
	var tests = []struct {
		input *Order
		err   bool
	}{
		{NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)), true},
		{NewOrder("s2", Sell, decimal.NewFromFloat(10.0), decimal.NewFromFloat(7000.0)), false},
		{NewOrder("b3", Buy, decimal.NewFromFloat(11.0), decimal.NewFromFloat(8000.0)), true},
		{NewOrder("s4", Sell, decimal.NewFromFloat(2.0), decimal.NewFromFloat(10000.0)), false},
	}
	ot := NewOrderType("sell")
	for _, tt := range tests {
		on, err := ot.AddOrderInQueue(*tt.input)
		if tt.err {
			if err == nil {
				t.Fatalf("Cannot append %s order under %s order type", tt.input.Type, ot.Type)
			}
			continue
		}
		fmt.Println("on", on, err)
		if on.Volume != tt.input.Amount {
			t.Fatalf("Volume update failure (have: %s, want: %s)", on.Volume.String(), tt.input.Amount.String())
		}
		if len(on.Orders) != 1 {
			t.Fatalf("Order length update failure (have: %d, want: 1)", len(on.Orders))
		}
	}

	price, _ := tests[1].input.Price.Float64()
	node := ot.Tree.Root.SearchSubTree(price)
	if node == nil {
		t.Fatal("There should exists a node in orderType.Tree")
	}
}
