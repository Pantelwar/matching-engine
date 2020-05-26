package engine

import (
	"math"
	"testing"

	"github.com/Pantelwar/binarytree"
)

func TestNewOrderBook(t *testing.T) {
	t.Log(NewOrderBook())
}

func TestAddOrderInBook(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, 5.0, 7000.0)},
		{NewOrder("s2", Sell, 10.0, 7000.0)},
		{NewOrder("s3", Sell, 10.0, 7000.0)},
		{NewOrder("b4", Buy, 1.0, 7000.0)},
	}
	ob := NewOrderBook()

	for _, tt := range tests {
		if tt.input.Type == Buy {
			ob.addBuyOrder(*tt.input)
		} else {
			ob.addSellOrder(*tt.input)
		}

		if ob.orders[tt.input.ID] == nil {
			t.Fatal("Order should be pushed in orders array")
		}

		startPoint := float64(int(math.Ceil(tt.input.Price)) / ob.orderLimitRange * ob.orderLimitRange)
		endPoint := startPoint + float64(ob.orderLimitRange)
		searchNodePrice := (startPoint + endPoint) / 2

		var node *binarytree.BinaryNode
		if tt.input.Type == Buy {
			node = ob.BuyTree.Root.SearchSubTree(searchNodePrice)
		} else {
			node = ob.SellTree.Root.SearchSubTree(searchNodePrice)
		}

		if node == nil {
			t.Fatal("Order should be present in tree")
		}
	}
}

func TestRemoveOrderNodeFromBook(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, 5.0, 7000.0)},
		{NewOrder("s2", Sell, 10.0, 7000.0)},
		{NewOrder("s3", Sell, 10.0, 7000.0)},
		{NewOrder("b4", Buy, 1.0, 7000.0)},
	}
	ob := NewOrderBook()

	for _, tt := range tests {
		if tt.input.Type == Buy {
			ob.addBuyOrder(*tt.input)
		} else {
			ob.addSellOrder(*tt.input)
		}
	}

	startPoint := float64(int(math.Ceil(tests[0].input.Price)) / ob.orderLimitRange * ob.orderLimitRange)
	endPoint := startPoint + float64(ob.orderLimitRange)
	searchNodePrice := (startPoint + endPoint) / 2

	ob.removeBuyNode(searchNodePrice)

	var node *binarytree.BinaryNode
	node = ob.BuyTree.Root.SearchSubTree(searchNodePrice)

	if node != nil {
		t.Fatal("Buy Mid Price should be get removed from tree")
	}

	startPoint = float64(int(math.Ceil(tests[1].input.Price)) / ob.orderLimitRange * ob.orderLimitRange)
	endPoint = startPoint + float64(ob.orderLimitRange)
	searchNodePrice = (startPoint + endPoint) / 2

	ob.removeSellNode(searchNodePrice)

	node = ob.SellTree.Root.SearchSubTree(searchNodePrice)

	if node != nil {
		t.Fatal("Sell Mid Price should be get removed from tree")
	}
}

func TestRemoveOrderFromBook(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, 5.0, 7000.0)},
		{NewOrder("s2", Sell, 10.0, 7000.0)},
		{NewOrder("s3", Sell, 10.0, 7000.0)},
		{NewOrder("b4", Buy, 1.0, 7000.0)},
	}
	ob := NewOrderBook()

	for _, tt := range tests {
		if tt.input.Type == Buy {
			ob.addBuyOrder(*tt.input)
		} else {
			ob.addSellOrder(*tt.input)
		}
	}

	err := ob.removeOrder(tests[0].input)
	if err != nil {
		t.Fatal("Order is not removed")
	}

	err = ob.removeOrder(tests[0].input)
	if err == nil {
		t.Fatal("The response should be order not found")
	}
}
