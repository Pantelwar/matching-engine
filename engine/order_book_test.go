package engine

import (
	"math"
	"testing"

	"github.com/Pantelwar/binarytree"
	"github.com/Pantelwar/matching-engine/util"
)

func TestNewOrderBook(t *testing.T) {
	t.Log(NewOrderBook())
}

func DecimalBig(val string) *util.StandardBigDecimal {
	a, _ := util.NewDecimalFromString(val)
	return a
}

func TestAddOrderInBook(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0"))},
		{NewOrder("s2", Sell, DecimalBig("10.0"), DecimalBig("7000.0"))},
		{NewOrder("s3", Sell, DecimalBig("10.0"), DecimalBig("7000.0"))},
		{NewOrder("b4", Buy, DecimalBig("1.0"), DecimalBig("7000.0"))},
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

		price := tt.input.Price.Float64()
		startPoint := float64(int(math.Ceil(price)) / ob.orderLimitRange * ob.orderLimitRange)
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
		{NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0"))},
		{NewOrder("s2", Sell, DecimalBig("10.0"), DecimalBig("7000.0"))},
		{NewOrder("s3", Sell, DecimalBig("10.0"), DecimalBig("7000.0"))},
		{NewOrder("b4", Buy, DecimalBig("1.0"), DecimalBig("7000.0"))},
	}
	ob := NewOrderBook()

	for _, tt := range tests {
		if tt.input.Type == Buy {
			ob.addBuyOrder(*tt.input)
		} else {
			ob.addSellOrder(*tt.input)
		}
	}

	price := tests[0].input.Price.Float64()
	startPoint := float64(int(math.Ceil(price)) / ob.orderLimitRange * ob.orderLimitRange)
	endPoint := startPoint + float64(ob.orderLimitRange)
	searchNodePrice := (startPoint + endPoint) / 2

	ob.removeBuyNode(searchNodePrice)

	var node *binarytree.BinaryNode
	node = ob.BuyTree.Root.SearchSubTree(searchNodePrice)

	if node != nil {
		t.Fatal("Buy Mid Price should be get removed from tree")
	}

	price = tests[1].input.Price.Float64()
	startPoint = float64(int(math.Ceil(price)) / ob.orderLimitRange * ob.orderLimitRange)
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
		{NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0"))},
		{NewOrder("s2", Sell, DecimalBig("10.0"), DecimalBig("7000.0"))},
		{NewOrder("s3", Sell, DecimalBig("10.0"), DecimalBig("7000.0"))},
		{NewOrder("b4", Buy, DecimalBig("1.0"), DecimalBig("7000.0"))},
	}
	ob := NewOrderBook()

	for _, tt := range tests {
		if tt.input.Type == Buy {
			ob.addBuyOrder(*tt.input) //, orderPrice)
		} else {
			ob.addSellOrder(*tt.input) //, orderPrice)
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
				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
				NewOrder("b2", Buy, DecimalBig("10.0"), DecimalBig("7000.0")),
			},
			`------------------------------------------
7000 -> 15
`},
		{
			[]*Order{
				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
				NewOrder("b2", Buy, DecimalBig("10.0"), DecimalBig("8000.0")),
			},
			`------------------------------------------
8000 -> 10
7000 -> 5
`},
		{
			[]*Order{
				NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
				NewOrder("s2", Sell, DecimalBig("10.0"), DecimalBig("7000.0")),
			},
			`7000 -> 15
------------------------------------------
`},
		{
			[]*Order{
				NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
				NewOrder("s2", Sell, DecimalBig("10.0"), DecimalBig("8000.0")),
			},
			`8000 -> 10
7000 -> 5
------------------------------------------
`},
		{
			[]*Order{
				NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
				NewOrder("b2", Buy, DecimalBig("10.0"), DecimalBig("6000.0")),
				NewOrder("s3", Sell, DecimalBig("1.0"), DecimalBig("8000.0")),
				NewOrder("b4", Buy, DecimalBig("2.0"), DecimalBig("6500.0")),
			},
			`8000 -> 1
7000 -> 5
------------------------------------------
6500 -> 2
6000 -> 10
`},
		{
			[]*Order{
				NewOrder("s1", Sell, DecimalBig("5.134"), DecimalBig("7000.0")),
				NewOrder("b2", Buy, DecimalBig("10.134"), DecimalBig("6000.0")),
				NewOrder("s3", Sell, DecimalBig("1.32"), DecimalBig("7000.0")),
				NewOrder("b4", Buy, DecimalBig("2.1278"), DecimalBig("6000.0")),
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
