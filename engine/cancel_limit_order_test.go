package engine

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestCancelOrder(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b2", Buy, decimal.NewFromFloat(10.0), decimal.NewFromFloat(6000.0))},
		{NewOrder("b3", Buy, decimal.NewFromFloat(11.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b4", Buy, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("s1", Sell, decimal.NewFromFloat(5.0), decimal.NewFromFloat(8000.0))},
		{NewOrder("s2", Sell, decimal.NewFromFloat(10.0), decimal.NewFromFloat(9000.0))},
		{NewOrder("s3", Sell, decimal.NewFromFloat(11.0), decimal.NewFromFloat(9000.0))},
		{NewOrder("s4", Sell, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7500.0))},
	}
	ob := NewOrderBook()

	for _, tt := range tests {
		orderPrice, _ := tt.input.Price.Float64()
		if tt.input.Type == Buy {
			ob.addBuyOrder(*tt.input, orderPrice)
		} else {
			ob.addSellOrder(*tt.input, orderPrice)
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
