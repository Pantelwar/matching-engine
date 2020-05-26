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

func TestString(t *testing.T) {
	var tests = []struct {
		input  []*Order
		output string
	}{
		{
			[]*Order{
				NewOrder("b1", Buy, 5.0, 7000.0),
				NewOrder("b2", Buy, 10.0, 7000.0),
			},
			`------------------------------------------
7000 -> 15
`},
		{
			[]*Order{
				NewOrder("b1", Buy, 5.0, 7000.0),
				NewOrder("b2", Buy, 10.0, 8000.0),
			},
			`------------------------------------------
8000 -> 10
7000 -> 5
`},
		{
			[]*Order{
				NewOrder("s1", Sell, 5.0, 7000.0),
				NewOrder("s2", Sell, 10.0, 7000.0),
			},
			`7000 -> 15
------------------------------------------
`},
		{
			[]*Order{
				NewOrder("s1", Sell, 5.0, 7000.0),
				NewOrder("s2", Sell, 10.0, 8000.0),
			},
			`8000 -> 10
7000 -> 5
------------------------------------------
`},
		{
			[]*Order{
				NewOrder("s1", Sell, 5.0, 7000.0),
				NewOrder("b2", Buy, 10.0, 6000.0),
				NewOrder("s3", Sell, 1.0, 8000.0),
				NewOrder("b4", Buy, 2.0, 6500.0),
			},
			`8000 -> 1
7000 -> 5
------------------------------------------
6500 -> 2
6000 -> 10
`},
		{
			[]*Order{
				NewOrder("s1", Sell, 5.134, 7000.0),
				NewOrder("b2", Buy, 10.134, 6000.0),
				NewOrder("s3", Sell, 1.32, 7000.0),
				NewOrder("b4", Buy, 2.1278, 6000.0),
			},
			`7000 -> 6.454
------------------------------------------
6000 -> 12.2618
`},
	}

	for _, tt := range tests {
		ob := NewOrderBook()
		for _, o := range tt.input {
			ob.Process(*o)
		}

		if tt.output != ob.String() {
			t.Fatalf("Book prints incorrect (have: \n%s, \nwant: \n%s\n)", ob.String(), tt.output)
		}
	}
}
