package engine

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
	// n := len(book.SellOrders)

	// // check if we have at least one matching order
	// if n != 0 && book.SellOrders[0].Price <= order.Price {
	// 	// traverse all orders that match
	// 	sellOrder := book.SellOrders[0]
	// 	// fmt.Printf("sellOrder.Price: %f, order.Price: %f\n", sellOrder.Price, order.Price)

	// 	for sellOrder.Price <= order.Price {
	// 		if (len(book.SellOrders) == 0) {
	// 			break
	// 		}
	// 		sellOrder = book.SellOrders[0]
	// 		// fmt.Printf("%#v\n", sellOrder)
	// 		// fmt.Printf("sellOrder.Price: %f, order.Price: %f\n", sellOrder.Price, order.Price)
	// 		if sellOrder.Price > order.Price {
	// 			break
	// 		}
	// 		// fill the entire order
	// 		// fmt.Printf("sellOrder.Amount: %f, order.Amount: %f\n", sellOrder.Amount, order.Amount)
	// 		if sellOrder.Amount >= order.Amount {
	// 			trades = append(trades, Trade{order.ID, sellOrder.ID, order.Amount, sellOrder.Price})
	// 			sellOrder.Amount -= order.Amount
	// 			if sellOrder.Amount == 0 {
	// 				book.removeSellOrder(0)
	// 			}
	// 			return trades
	// 		}
	// 		// fill a partial order and continue
	// 		if sellOrder.Amount < order.Amount {
	// 			trades = append(trades, Trade{order.ID, sellOrder.ID, sellOrder.Amount, sellOrder.Price})
	// 			order.Amount -= sellOrder.Amount
	// 			book.removeSellOrder(0)
	// 			continue
	// 		}
	// 	}
	// }
	// finally add the remaining order to the list
	book.addBuyOrder(order)
	return trades
}

// Process a limit sell order
func (book *OrderBook) processLimitSell(order Order) []Trade {
	trades := make([]Trade, 0, 1)

	// n := book.BuyOrders.Tree.SearchSubTree(book.BuyOrders.Tree.Root, order.Price)
	// fmt.Println("found", n, order.Price)
	// n.Print(os.Stdout, 0, 'M')
	// m := book.BuyOrders.Tree.GreatThan(n, order.Price)
	// fmt.Println("found", m, order.Price)
	// m.Print(os.Stdout, 0, 'M')
	// n := len(book.BuyOrders)
	// book.BuyOrders.Tree.Root.Remove(7340.0)
	// book.BuyOrders.Tree.Root.Remove(7333.0)
	// node := book.BuyOrders.Tree.Root.Remove(7303.0)
	// book.BuyOrders.Tree.Root.Remove(7333.0)
	// node := book.BuyOrders.Tree.Root.Remove(7340.0)
	// fmt.Println("node: ", node)

	maxNode := book.BuyOrders.Tree.Max()
	countAdd := 0.0
	for maxNode == nil || order.Amount > 0 {
		maxNode = book.BuyOrders.Tree.Max()
		if maxNode == nil {
			book.addBuyOrder(order)
			break
		}
		countAdd += maxNode.Amount
		// fmt.Println(order.Price, "AMOUNT", order.Amount, maxNode, maxNode.Amount > order.Amount)
		if maxNode.Amount > order.Amount {
			amount := maxNode.Amount - order.Amount
			// amount = math.Floor(amount*100000000) / 100000000
			maxNode.SetAmount(amount)
			break
		}
		if maxNode.Amount == order.Amount {
			// fmt.Println("Removing Node")
			node := book.BuyOrders.Tree.Root.Remove(maxNode.Price)
			book.BuyOrders.Tree.Root = node
			book.addBuyOrder(order)
			break
		} else {
			// fmt.Println("Removing Node and continue")
			order.Amount -= maxNode.Amount
			// order.Amount = math.Floor(order.Amount*100000000) / 100000000
			node := book.BuyOrders.Tree.Root.Remove(maxNode.Price)
			book.BuyOrders.Tree.Root = node
		}
		// continue
	}
	// fmt.Println("FinalCount: ", countAdd)
	// // check if we have at least one matching order
	// if n != 0 && book.BuyOrders[0].Price >= order.Price {
	// 	// traverse all orders that match
	// 	buyOrder := book.BuyOrders[0]
	// 	// fmt.Printf("buyOrder.Price: %f, order.Price: %f\n", buyOrder.Price, order.Price)

	// 	for buyOrder.Price >= order.Price {
	// 		if (len(book.BuyOrders) == 0) {
	// 			break
	// 		}
	// 		buyOrder = book.BuyOrders[0]
	// 		// fmt.Printf("%#v\n", buyOrder)
	// 		if buyOrder.Price < order.Price {
	// 			break
	// 		}
	// 		// fill the entire order
	// 		if buyOrder.Amount >= order.Amount {
	// 			trades = append(trades, Trade{order.ID, buyOrder.ID, order.Amount, buyOrder.Price})
	// 			buyOrder.Amount -= order.Amount
	// 			if buyOrder.Amount == 0 {
	// 				book.removeBuyOrder(0)
	// 			}
	// 			return trades
	// 		}
	// 		// fill a partial order and continue
	// 		if buyOrder.Amount < order.Amount {
	// 			trades = append(trades, Trade{order.ID, buyOrder.ID, buyOrder.Amount, buyOrder.Price})
	// 			order.Amount -= buyOrder.Amount
	// 			book.removeBuyOrder(0)
	// 			continue
	// 		}
	// 	}
	// }
	// finally add the remaining order to the list
	// book.addSellOrder(order)
	return trades
}
