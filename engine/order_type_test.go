package engine

import (
	"fmt"
	"testing"
)

func TestNewOrderType(t *testing.T) {
	t.Log(NewOrderType("sell"))
}

func TestAddOrderInQueue(t *testing.T) {
	var tests = []struct {
		input *Order
		err   bool
	}{
		{NewOrder("b1", Buy, 5.0, 7000.0), true},
		{NewOrder("s2", Sell, 10.0, 7000.0), false},
		{NewOrder("b3", Buy, 11.0, 8000.0), true},
		{NewOrder("s4", Sell, 2.0, 10000.0), false},
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
			t.Fatalf("Volume update failure (have: %f, want: %f)", on.Volume, tt.input.Amount)
		}
		if len(on.Orders) != 1 {
			t.Fatalf("Order length update failure (have: %d, want: 1)", len(on.Orders))
		}
	}

	node := ot.Tree.Root.SearchSubTree(tests[1].input.Price)
	if node == nil {
		t.Fatal("There should exists a node in orderType.Tree")
	}
}
