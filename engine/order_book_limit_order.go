package engine

import "fmt"

// import "fmt"

// Process an order and return the trades generated before adding the remaining amount to the market
func (book *OrderBook) Process(order Order) []Trade {
	if order.Type == "buy" {
		return book.processLimitBuy(order)
	}
	return book.processLimitSell(order)
}

// Process a limit buy order
func (book *OrderBook) processLimitBuy(order Order) []Trade {
	trades := make([]Trade, 0, 1)
	n := len(book.SellOrders)
	// fmt.Println("len(book.SellOrders)", n)
	// fmt.Println("len(book.BuyOrders)", len(book.BuyOrders))
	// check if we have at least one matching order
	if n != 0 && book.SellOrders[0].Price <= order.Price {
		// traverse all orders that match
		sellOrder := book.SellOrders[0]
		fmt.Printf("buyOrder.Price: %f, order.Price: %f\n", sellOrder.Price, order.Price)

		for sellOrder.Price <= order.Price {
		// for i := n - 1; i >= 0; i-- {
			sellOrder = book.SellOrders[0]
			if sellOrder.Price > order.Price {
				break
			}
			// fill the entire order
			if sellOrder.Amount >= order.Amount {
				trades = append(trades, Trade{order.ID, sellOrder.ID, order.Amount, sellOrder.Price})
				sellOrder.Amount -= order.Amount
				if sellOrder.Amount == 0 {
					book.removeSellOrder(0)
				}
				return trades
			}
			// fill a partial order and continue
			if sellOrder.Amount < order.Amount {
				trades = append(trades, Trade{order.ID, sellOrder.ID, sellOrder.Amount, sellOrder.Price})
				order.Amount -= sellOrder.Amount
				book.removeSellOrder(0)
				continue
			}
		}
	}
	// finally add the remaining order to the list
	book.addBuyOrder(order)
	return trades
}

// Process a limit sell order
func (book *OrderBook) processLimitSell(order Order) []Trade {
	trades := make([]Trade, 0, 1)
	n := len(book.BuyOrders)
	// fmt.Println("len(book.SellOrders)", len(book.SellOrders))
	// fmt.Println("len(book.BuyOrders)", n)
	// check if we have at least one matching order
	fmt.Printf("book.BuyOrders[0].Price: %f, order.Price: %f\n", book.BuyOrders[0].Price, order.Price)

	if n != 0 && book.BuyOrders[0].Price >= order.Price {
		// traverse all orders that match
		buyOrder := book.BuyOrders[0]
		fmt.Printf("buyOrder.Price: %f, order.Price: %f\n", buyOrder.Price, order.Price)

		for buyOrder.Price >= order.Price {
		// for i := 0; i <= n-1; i++ {
			buyOrder = book.BuyOrders[0]
			fmt.Printf("%#v\n", buyOrder)
			if buyOrder.Price < order.Price {
				break
			}
			// fill the entire order
			if buyOrder.Amount >= order.Amount {
				trades = append(trades, Trade{order.ID, buyOrder.ID, order.Amount, buyOrder.Price})
				buyOrder.Amount -= order.Amount
				if buyOrder.Amount == 0 {
					book.removeBuyOrder(0)
				}
				return trades
			}
			// fill a partial order and continue
			if buyOrder.Amount < order.Amount {
				trades = append(trades, Trade{order.ID, buyOrder.ID, buyOrder.Amount, buyOrder.Price})
				order.Amount -= buyOrder.Amount
				book.removeBuyOrder(0)
				continue
			}
		}
	}
	// finally add the remaining order to the list
	book.addSellOrder(order)
	return trades
}
