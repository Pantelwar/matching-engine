package engine

import "fmt"

// OrderBook type
type OrderBook struct {
	BuyOrders  []Order
	SellOrders []Order
}

// Add a buy order to the order book
func (book *OrderBook) addBuyOrder(order Order) {
	n := len(book.BuyOrders)
	var i int
	found := false
	for i = 0; i <= n-1; i++ {
		buyOrder := book.BuyOrders[i]
		fmt.Println("i:", i)
		if buyOrder.Price < order.Price {
			found = true
			fmt.Println("Found at ", i)
			break
		}
	}
	fmt.Printf("i: %d, n: %d\n", i, n)
	if n == 0 || (i == 0 && !found) {
		book.BuyOrders = append(book.BuyOrders, order)
	} else {
		book.BuyOrders = append(book.BuyOrders, order)
		fmt.Println("book.BuyOrders[i+1:]: ", book.BuyOrders[i+1:])
		fmt.Println("book.BuyOrders[i:]: ", book.BuyOrders[i:])
		copy(book.BuyOrders[i+1:], book.BuyOrders[i:])
		fmt.Println("Copy: ", book.BuyOrders)

		book.BuyOrders[i] = order
		fmt.Println("Done: ", book.BuyOrders)

	}
}

// Add a sell order to the order book
func (book *OrderBook) addSellOrder(order Order) {
	n := len(book.SellOrders)
	var i int
	found := false
	for i = 0; i <= n-1; i++ {
		sellOrder := book.SellOrders[i]
		fmt.Println("i:", i)
		if sellOrder.Price > order.Price {
			found = true
			fmt.Println("Found at ", i)
			break
		}
	}
	if n == 0 || (i == 0 && !found) {
		book.SellOrders = append(book.SellOrders, order)
	} else {
		book.SellOrders = append(book.SellOrders, order)
		copy(book.SellOrders[i+1:], book.SellOrders[i:])
		book.SellOrders[i] = order
	}
}

// Remove a buy order from the order book at a given index
func (book *OrderBook) removeBuyOrder(index int) {
	book.BuyOrders = append(book.BuyOrders[:index], book.BuyOrders[index+1:]...)
}

// Remove a sell order from the order book at a given index
func (book *OrderBook) removeSellOrder(index int) {
	book.SellOrders = append(book.SellOrders[:index], book.SellOrders[index+1:]...)
}
