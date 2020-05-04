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
		// order.Amount += node.Data.(Order).Amount
		arr := append(node.Data.([]*Order), &order)
		node.SetData(arr)
		return
	}
	book.BuyOrders.AddOrderInQueue(order)
	// book.BuyOrders.Tree.Insert(order.Price, order)
}

// Add a sell order to the order book
func (book *OrderBook) addSellOrder(order Order) {
	node := book.SellOrders.Tree.Root.SearchSubTree(order.Price)
	if node != nil {
		arr := append(node.Data.([]*Order), &order)
		node.SetData(arr)
		return
	}
	book.SellOrders.AddOrderInQueue(order)
	// book.SellOrders.Tree.Insert(order.Price, order)
}

// Remove a buy order from the order book at a given index
func (book *OrderBook) removeBuyOrder(key float64) *binarytree.BinaryNode {
	return book.BuyOrders.Tree.Root.Remove(key)
}

// Remove a sell order from the order book at a given index
func (book *OrderBook) removeSellOrder(key float64) *binarytree.BinaryNode {
	return book.SellOrders.Tree.Root.Remove(key)
}

func (book *OrderBook) Print() {
	fmt.Println("BuyOrders: ")
	print(os.Stdout, book.BuyOrders.Tree.Root, 0, 'M')

	fmt.Println()

	fmt.Println("SellOrders: ")
	print(os.Stdout, book.SellOrders.Tree.Root, 0, 'M')
}

func print(w io.Writer, node *binarytree.BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v -> %#v\n", ch, node.Key, node.Data)
	for _, val := range node.Data.([]*Order) {
		for i := 0; i < ns; i++ {
			fmt.Fprint(w, " ")
		}
		fmt.Fprintf(w, "%c:     -> %#v\n", ch, val)
	}
	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}
