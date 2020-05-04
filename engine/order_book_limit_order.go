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
	book.addBuyOrder(order)
	return trades
}

// Process a limit sell order
func (book *OrderBook) processLimitSell(order Order) []Trade {
	trades := make([]Trade, 0, 1)

	maxNode := book.BuyOrders.Tree.Max()
	// nodeData := maxNode.Data.([]*Order)
	countAdd := 0.0
	for maxNode == nil || order.Amount > 0 {
		maxNode = book.BuyOrders.Tree.Max()
		if maxNode == nil && order.Amount > 0 {
			book.addSellOrder(order)
			break
		}
		nodeData := maxNode.Data.([]*Order)
		countMatch := 0
		for _, ele := range nodeData {
			countAdd += ele.Amount
			// fmt.Println(order.Price, "AMOUNT", order.Amount, ele, ele.Amount > order.Amount)
			if ele.Amount > order.Amount {
				amount := ele.Amount - order.Amount
				// amount = math.Floor(amount*100000000) / 100000000
				ele.Amount = amount
				maxNode.SetData(nodeData)

				trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				order.Amount = 0.0

				break
			}
			if ele.Amount == order.Amount {
				// fmt.Println("Removing Node")
				countMatch++
				ele.Amount = 0
				// node := book.removeBuyOrder(maxNode.Key)
				// book.BuyOrders.Tree.Root = node
				// book.addSellOrder(order)

				trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				// break
			} else {
				// fmt.Println("Removing Node and continue")
				countMatch++
				order.Amount -= ele.Amount
				ele.Amount = 0.0
				// order.Amount = math.Floor(order.Amount*100000000) / 100000000
				// node := book.removeBuyOrder(maxNode.Key)
				// book.BuyOrders.Tree.Root = node

				trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: ele.Amount, Price: ele.Price})

			}
		}

		if len(nodeData) == countMatch {
			node := book.removeBuyOrder(maxNode.Key)
			book.BuyOrders.Tree.Root = node
		}

		// fmt.Println("countMatch", countMatch, nodeData[countMatch:])
		nodeData = nodeData[countMatch:]
		maxNode.SetData(nodeData)

		// continue
	}
	return trades
}
