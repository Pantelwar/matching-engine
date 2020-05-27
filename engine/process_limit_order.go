package engine

import (
	"fmt"

	"github.com/Pantelwar/binarytree"
)

// Process executes limit process
func (ob *OrderBook) Process(order Order) ([]*Order, *Order) {
	if order.Type == Buy {
		// return ob.processOrderB(order)
		return ob.commonProcess(order, ob.SellTree, ob.addBuyOrder, ob.removeSellNode)
	}
	// return ob.processOrderS(order)
	return ob.commonProcess(order, ob.BuyTree, ob.addSellOrder, ob.removeBuyNode)
}

func (ob *OrderBook) commonProcess(order Order, tree *binarytree.BinaryTree, add func(Order), remove func(float64) error) ([]*Order, *Order) {
	var maxNode *binarytree.BinaryNode
	if order.Type == Sell {
		maxNode = tree.Max()
	} else {
		maxNode = tree.Min()
	}
	if maxNode == nil {
		add(order)
		// ob.addSellOrder(order)
		// return trades, nil, nil
		return nil, nil
	}
	count := 0
	noMoreOrders := false
	var allOrdersProcessed []*Order
	var partialOrder *Order
	for maxNode == nil || GreaterThan(order.Amount, "0") { // order.Amount > 0 {
		count++
		if order.Type == Sell {
			maxNode = tree.Max()
		} else {
			maxNode = tree.Min()
		}
		if maxNode == nil || noMoreOrders {
			if GreaterThan(order.Amount, "0") { //order.Amount > 0 {
				// fmt.Println("adding sell node pending")
				add(order)
				//ob.addSellOrder(order)
				break
			} else {
				break
			}
		}

		// var t []Trade
		var ordersProcessed []*Order
		noMoreOrders, ordersProcessed, partialOrder = ob.processLimit(&order, maxNode.Data.(*OrderType).Tree)
		allOrdersProcessed = append(allOrdersProcessed, ordersProcessed...)
		// trades = append(trades, t...)

		if maxNode.Data.(*OrderType).Tree.Root == nil {
			// node := remove(maxNode.Key)
			// // node := ob.removeBuyNode(maxNode.Key)
			// tree.Root = node
			remove(maxNode.Key)
		}
	}

	// return trades, allOrdersProcessed, partialOrder
	return allOrdersProcessed, partialOrder
}

func (ob *OrderBook) processLimit(order *Order, tree *binarytree.BinaryTree) (bool, []*Order, *Order) {
	// trades := make([]Trade, 0, 1)

	var maxNode *binarytree.BinaryNode
	if order.Type == Sell {
		maxNode = tree.Max()
	} else {
		maxNode = tree.Min()
	}
	noMoreOrders := false
	var ordersProcessed []*Order
	var partialOrder *Order
	if maxNode == nil {
		// return trades, noMoreOrders, nil, nil
		return noMoreOrders, nil, nil
	}
	// countAdd := 0.0
	for maxNode == nil || GreaterThan(order.Amount, "0") { //} order.Amount > 0 {
		if order.Type == Sell {
			maxNode = tree.Max()
		} else {
			maxNode = tree.Min()
		}
		if maxNode == nil || noMoreOrders {
			if GreaterThan(order.Amount, "0") { //order.Amount > 0 {
				partialOrder = NewOrder(order.ID, order.Type, order.Amount, order.Price)
				break
			} else {
				break
			}
		}
		if order.Type == Sell {
			if GreaterThan(order.Price, fmt.Sprintf("%f", maxNode.Key)) { //}
				// if order.Price > maxNode.Key {
				// fmt.Println("adding sellnode directly")
				noMoreOrders = true
				// return trades, noMoreOrders, nil, nil
				return noMoreOrders, nil, nil
			}
		} else {
			if LessThan(order.Price, fmt.Sprintf("%f", maxNode.Key)) { //}
				// if order.Price < maxNode.Key {
				// fmt.Println("adding buynode directly")
				noMoreOrders = true
				// return trades, noMoreOrders, nil, nil
				return noMoreOrders, nil, nil
			}
		}

		nodeData := maxNode.Data.(*OrderNode) //([]*Order)
		nodeOrders := nodeData.Orders         //([]*Order)
		countMatch := 0
		for _, ele := range nodeOrders {
			if order.Type == Sell {
				if ele.Price < order.Price {
					noMoreOrders = true
					break
				}
			} else {
				if ele.Price > order.Price {
					noMoreOrders = true
					break
				}
			}

			// countAdd += ele.Amount
			if GreaterThan(ele.Amount, order.Amount) {
				orderAmount := Sub("0", order.Amount)
				nodeData.updateVolume(orderAmount)
				// trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				// amount := ele.Amount - order.Amount
				amount := Sub(ele.Amount, order.Amount)
				// amount = math.Floor(amount*100000000) / 100000000
				ele.Amount = amount

				partialOrder = NewOrder(ele.ID, ele.Type, ele.Amount, ele.Price)
				ordersProcessed = append(ordersProcessed, NewOrder(order.ID, order.Type, order.Amount, order.Price))

				maxNode.SetData(nodeData)

				order.Amount = "0" //0.0
				noMoreOrders = true
				break
			}
			if Equal(ele.Amount, order.Amount) {
				orderAmount := Sub("0", order.Amount)
				nodeData.updateVolume(orderAmount)

				ordersProcessed = append(ordersProcessed, NewOrder(ele.ID, ele.Type, ele.Amount, ele.Price))
				ordersProcessed = append(ordersProcessed, NewOrder(order.ID, order.Type, order.Amount, order.Price))

				countMatch++
				// trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				order.Amount = "0" //0.0
				// orderComplete = true

				// ele.Amount = 0
				delete(ob.orders, ele.ID)

				break
			} else {
				countMatch++

				ordersProcessed = append(ordersProcessed, NewOrder(ele.ID, ele.Type, ele.Amount, ele.Price))

				eleAmount := Sub("0", ele.Amount)
				nodeData.updateVolume(eleAmount)

				// trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: ele.Amount, Price: ele.Price})

				// order.Amount -= ele.Amount
				order.Amount = Sub(order.Amount, ele.Amount)

				delete(ob.orders, ele.ID)

			}
		}

		if len(nodeOrders) == countMatch {
			node := tree.Root.Remove(maxNode.Key) // ob.removeBuyNode(maxNode.Key, buyTree)
			// fmt.Printf("node removed: %#v %#v\n", node, maxNode)
			tree.Root = node
		}

		nodeData.Orders = nodeOrders[countMatch:]
		maxNode.SetData(nodeData)
	}
	// return trades, noMoreOrders, ordersProcessed, partialOrder
	return noMoreOrders, ordersProcessed, partialOrder
}
