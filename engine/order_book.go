package engine

import (
	"fmt"
	"io"
	"matching-engine/engine/binarytree"
	"os"
)

// OrderBook type
type OrderBook struct {
	BuyOrders  *OrderType
	SellOrders *OrderType
}

// NewOrderBook Returns new order book
func NewOrderBook() *OrderBook {
	return &OrderBook{
		BuyOrders:  NewOrderType("buy"),
		SellOrders: NewOrderType("sell"),
	}
}

// Add a buy order to the order book
func (book *OrderBook) addBuyOrder(order Order) {
	node := book.BuyOrders.Tree.Root.SearchSubTree(order.Price)
	if node != nil {
		node.SetAmount(order.Amount + node.Amount)
		return
	}
	book.BuyOrders.Tree.Insert(order.Price, order.Amount)
}

// Add a sell order to the order book
func (book *OrderBook) addSellOrder(order Order) {

}

// Remove a buy order from the order book at a given index
func (book *OrderBook) removeBuyOrder(index int) {
}

// Remove a sell order from the order book at a given index
func (book *OrderBook) removeSellOrder(index int) {

}

func (book *OrderBook) Print() {
	print(os.Stdout, book.BuyOrders.Tree.Root, 0, 'M')
}

func print(w io.Writer, node *binarytree.BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v -> %v\n", ch, node.Price, node.Amount)
	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}
