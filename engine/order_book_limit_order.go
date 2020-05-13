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

	minNode := book.SellOrders.Tree.Min()
	if minNode == nil {
		book.addBuyOrder(order)
		return trades
	}
	countAdd := 0.0
	noMoreOrders := false
	for minNode == nil || order.Amount > 0 {
		minNode = book.SellOrders.Tree.Min()
		if minNode == nil || noMoreOrders {
			if order.Amount > 0 {
				book.addBuyOrder(order)
				break
			} else {
				break
			}
		}
		nodeData := minNode.Data.(*OrderNode) //([]*Order)
		nodeOrders := nodeData.Orders         //([]*Order)
		countMatch := 0
		for _, ele := range nodeOrders {
			countAdd += ele.Amount
			fmt.Println("ele.Price", ele.Price)
			if ele.Price > order.Price {
				fmt.Println("No more orders")
				noMoreOrders = true
				break
			}
			// fmt.Println(order.Price, "AMOUNT", order.Amount, ele, ele.Amount > order.Amount)
			if ele.Amount > order.Amount {
				nodeData.UpdateVolume(-order.Amount)

				trades = append(trades, Trade{SellOrderID: ele.ID, BuyOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				amount := ele.Amount - order.Amount
				// amount = math.Floor(amount*100000000) / 100000000
				ele.Amount = amount

				minNode.SetData(nodeData)

				order.Amount = 0.0

				break
			}
			if ele.Amount == order.Amount {
				nodeData.UpdateVolume(-order.Amount)

				countMatch++
				ele.Amount = 0
				// node := book.removeBuyOrder(minNode.Key)
				// book.SellOrders.Tree.Root = node
				// book.addSellOrder(order)

				trades = append(trades, Trade{SellOrderID: ele.ID, BuyOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				// break
			} else {
				// fmt.Println("Removing Node and continue")
				countMatch++

				nodeData.UpdateVolume(-ele.Amount)

				order.Amount -= ele.Amount
				ele.Amount = 0.0
				// order.Amount = math.Floor(order.Amount*100000000) / 100000000
				// node := book.removeBuyOrder(minNode.Key)
				// book.SellOrders.Tree.Root = node

				trades = append(trades, Trade{SellOrderID: ele.ID, BuyOrderID: order.ID, Amount: ele.Amount, Price: ele.Price})

			}
		}

		if len(nodeOrders) == countMatch {
			node := book.removeSellOrder(minNode.Key)
			book.SellOrders.Tree.Root = node
		}

		// fmt.Println("countMatch", countMatch, nodeOrders[countMatch:])
		nodeData.Orders = nodeOrders[countMatch:]
		minNode.SetData(nodeData)

		// continue
	}
	return trades
}

// Process a limit sell order
func (book *OrderBook) processLimitSell(order Order) []Trade {
	trades := make([]Trade, 0, 1)

	maxNode := book.BuyOrders.Tree.Max()
	if maxNode == nil {
		book.addSellOrder(order)
		return trades
	}
	countAdd := 0.0
	noMoreOrders := false
	for maxNode == nil || order.Amount > 0 {
		maxNode = book.BuyOrders.Tree.Max()
		if maxNode == nil || noMoreOrders {
			if order.Amount > 0 {
				book.addSellOrder(order)
				break
			} else {
				break
			}
		}
		nodeData := maxNode.Data.(*OrderNode) //([]*Order)
		nodeOrders := nodeData.Orders         //([]*Order)
		countMatch := 0
		for _, ele := range nodeOrders {
			fmt.Println("ele.Price", ele.Price)
			if ele.Price < order.Price {
				fmt.Println("No more orders")
				noMoreOrders = true
				break
			}
			countAdd += ele.Amount
			// fmt.Println(order.Price, "AMOUNT", order.Amount, ele, ele.Amount > order.Amount)
			if ele.Amount > order.Amount {
				nodeData.UpdateVolume(-order.Amount)

				trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				amount := ele.Amount - order.Amount
				// amount = math.Floor(amount*100000000) / 100000000
				ele.Amount = amount

				maxNode.SetData(nodeData)

				order.Amount = 0.0

				break
			}
			if ele.Amount == order.Amount {
				nodeData.UpdateVolume(-order.Amount)

				countMatch++
				trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				ele.Amount = 0
				// node := book.removeBuyOrder(maxNode.Key)
				// book.BuyOrders.Tree.Root = node
				// book.addSellOrder(order)

				// break
			} else {
				// fmt.Println("Removing Node and continue")
				countMatch++

				nodeData.UpdateVolume(-ele.Amount)

				trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: ele.Amount, Price: ele.Price})

				order.Amount -= ele.Amount
				ele.Amount = 0.0
				// order.Amount = math.Floor(order.Amount*100000000) / 100000000
				// node := book.removeBuyOrder(maxNode.Key)
				// book.BuyOrders.Tree.Root = node

			}
		}

		if len(nodeOrders) == countMatch {
			node := book.removeBuyOrder(maxNode.Key)
			book.BuyOrders.Tree.Root = node
		}

		// fmt.Println("countMatch", countMatch, nodeOrders[countMatch:])
		nodeData.Orders = nodeOrders[countMatch:]
		maxNode.SetData(nodeData)

		// continue
	}
	return trades
}
