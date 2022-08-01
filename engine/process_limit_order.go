package engine

import (
	"fmt"

	"github.com/Pantelwar/binarytree"
	"github.com/Pantelwar/matching-engine/util"
)

var decimalZero, _ = util.NewDecimalFromString("0.0")

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
		// fmt.Println("adding node pending", order.Price)
		add(order)
		return nil, nil
	}
	// fmt.Println("maxNode", maxNode.Key, maxNode.Data.(*OrderType).Tree.Root.Key)
	count := 0
	noMoreOrders := false
	var allOrdersProcessed []*Order
	var partialOrder *Order
	orderOriginalAmount := order.Amount
	for maxNode == nil || order.Amount.Cmp(decimalZero) == 1 {
		count++
		if order.Type == Sell {
			maxNode = tree.Max()
		} else {
			maxNode = tree.Min()
		}
		if maxNode == nil || noMoreOrders {
			if order.Amount.Cmp(decimalZero) == 1 {
				// fmt.Println("adding sell node pending")
				add(order)
				break
			} else {
				break
			}
		}
		// fmt.Println("maxNode in", maxNode.Key, maxNode.Data.(*OrderType).Tree.Root.Key)

		// var t []Trade
		var ordersProcessed []*Order
		noMoreOrders, ordersProcessed, partialOrder = ob.processLimit(&order, partialOrder, maxNode.Data.(*OrderType).Tree, orderOriginalAmount) //, orderPrice)
		fmt.Printf("\npartialOrder in between: %#v\n", partialOrder)
		fmt.Printf("noMoreOrders: %#v\n\n", noMoreOrders)
		allOrdersProcessed = append(allOrdersProcessed, ordersProcessed...)
		// trades = append(trades, t...)

		if maxNode.Data.(*OrderType).Tree.Root == nil {
			// node := remove(maxNode.Key)
			// // node := ob.removeBuyNode(maxNode.Key)
			// tree.Root = node
			// fmt.Println("removing", maxNode.Key)
			remove(maxNode.Key)
		}
	}

	// return trades, allOrdersProcessed, partialOrder
	// if partialOrder.Amount == nil {
	// 	partialOrder = nil
	// }
	fmt.Printf("partialOrder final: %#v\n", partialOrder)

	return allOrdersProcessed, partialOrder
}

func (ob *OrderBook) processLimit(order, partialOrder *Order, tree *binarytree.BinaryTree, orderOriginalAmount *util.StandardBigDecimal) (bool, []*Order, *Order) {
	orderPrice := order.Price.Float64()
	var maxNode *binarytree.BinaryNode
	if order.Type == Sell {
		maxNode = tree.Max()
	} else {
		maxNode = tree.Min()
	}
	noMoreOrders := false
	var ordersProcessed []*Order
	// var partialOrder *Order
	fmt.Printf("partialOrder start: %#v\n", partialOrder)

	// var partialOrder *Order
	if maxNode == nil {
		// return trades, noMoreOrders, nil, nil
		return noMoreOrders, nil, nil
	}
	// countAdd := 0.0
	for maxNode == nil || order.Amount.Cmp(decimalZero) == 1 {
		if order.Type == Sell {
			maxNode = tree.Max()
		} else {
			maxNode = tree.Min()
		}
		// fmt.Println("maxNode more in", maxNode.Key, maxNode.Data.(*OrderNode).Orders)
		if maxNode == nil || noMoreOrders {
			if order.Amount.Cmp(decimalZero) == 1 {
				partialOrder = NewOrder(order.ID, order.Type, order.Amount, order.Price)
				break
			} else {
				// partialOrder = nil
				break
			}
		}
		if order.Type == Sell {
			if orderPrice > maxNode.Key {
				fmt.Println("adding sellnode directly")
				noMoreOrders = true
				// return trades, noMoreOrders, nil, nil
				return noMoreOrders, ordersProcessed, partialOrder
			}
		} else {
			if orderPrice < maxNode.Key {
				fmt.Println("adding buynode directly")
				noMoreOrders = true
				// return trades, noMoreOrders, nil, nil
				return noMoreOrders, ordersProcessed, partialOrder
			}
		}

		nodeData := maxNode.Data.(*OrderNode) //([]*Order)
		nodeOrders := nodeData.Orders         //([]*Order)
		countMatch := 0
		// fmt.Println("length of nodeOrders", len(nodeOrders))
		for _, ele := range nodeOrders {
			if order.Type == Sell {
				if ele.Price.Cmp(order.Price) == -1 {
					noMoreOrders = true
					break
				}
			} else {
				if ele.Price.Cmp(order.Price) == 1 {
					noMoreOrders = true
					break
				}
			}

			// countAdd += ele.Amount
			// fmt.Println(ele.Price, order.Price, ele.Amount, order.Amount, ele.Amount.Cmp(order.Amount))
			if ele.Amount.Cmp(order.Amount) == 1 {
				nodeData.updateVolume(order.Amount.Neg())
				// trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				amount := ele.Amount.Sub(order.Amount)
				// amount = math.Floor(amount*100000000) / 100000000
				ele.Amount = amount

				partialOrder = NewOrder(ele.ID, ele.Type, ele.Amount, ele.Price)
				ordersProcessed = append(ordersProcessed, NewOrder(order.ID, order.Type, orderOriginalAmount, order.Price))

				maxNode.SetData(nodeData)

				order.Amount, _ = util.NewDecimalFromString("0.0")
				noMoreOrders = true
				break
			} else if ele.Amount.Cmp(order.Amount) == 0 {
				nodeData.updateVolume(order.Amount.Neg())

				ordersProcessed = append(ordersProcessed, NewOrder(ele.ID, ele.Type, ele.Amount, ele.Price))
				ordersProcessed = append(ordersProcessed, NewOrder(order.ID, order.Type, orderOriginalAmount, order.Price))
				partialOrder = nil
				countMatch++
				// trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				order.Amount, _ = util.NewDecimalFromString("0.0")
				// orderComplete = true

				// ele.Amount = 0
				ob.mutex.Lock()
				delete(ob.orders, ele.ID)
				ob.mutex.Unlock()

				break
			} else {
				countMatch++

				amount := order.Amount.Sub(ele.Amount)

				partialOrder = NewOrder(order.ID, order.Type, amount, order.Price)

				ordersProcessed = append(ordersProcessed, NewOrder(ele.ID, ele.Type, ele.Amount, ele.Price))

				nodeData.updateVolume(ele.Amount.Neg())

				// trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: ele.Amount, Price: ele.Price})

				order.Amount = order.Amount.Sub(ele.Amount)
				ob.mutex.Lock()
				delete(ob.orders, ele.ID)
				ob.mutex.Unlock()
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
	// fmt.Printf("partialOrder: %#v\n", partialOrder)
	// fmt.Printf("order.Amount: %#v\n", order.Amount.String())
	return noMoreOrders, ordersProcessed, partialOrder
}
