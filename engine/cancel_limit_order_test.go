package engine

import (
	"testing"
)

func TestCancelOrder(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, "5.0", "7000.0")},
		{NewOrder("b2", Buy, "10.0", "5000.0")},
		{NewOrder("b3", Buy, "11.0", "7000.0")},
		{NewOrder("b4", Buy, "1.0", "7000.0")},
		{NewOrder("s1", Sell, "5.0", "8000.0")},
		{NewOrder("s2", Sell, "10.0", "9000.0")},
		{NewOrder("s3", Sell, "11.0", "9000.0")},
		{NewOrder("s4", Sell, "1.0", "7500.0")},
	}
	ob := NewOrderBook()

	for _, tt := range tests {
		if tt.input.Type == Buy {
			ob.addBuyOrder(*tt.input)
		} else {
			ob.addSellOrder(*tt.input)
		}
	}

	on := ob.orders[tests[4].input.ID]

	order := ob.CancelOrder("s1")

	for _, o := range on.Orders {
		if o.ID == tests[4].input.ID {
			t.Fatal("Order is not removed from the OrderNode")
		}
	}

	if order == nil {
		t.Fatal("Order is not removed")
	}

	err := ob.removeOrder(order)
	if err == nil {
		t.Fatal("Order is not removed from Tree of Orderbook")
	}

	if ob.orders[order.ID] != nil {
		t.Fatal("Order is not removed from \"orders\" of Orderbook")
	}
}
