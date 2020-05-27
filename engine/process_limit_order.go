package engine

import (
	"math/big"

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
	bigZero := new(big.Float).SetFloat64(0)
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
	for maxNode == nil || order.Amount.Cmp(bigZero) == 1 {
		count++
		if order.Type == Sell {
			maxNode = tree.Max()
		} else {
			maxNode = tree.Min()
		}
		if maxNode == nil || noMoreOrders {
			if order.Amount.Cmp(bigZero) == 1 {
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
	bigZero := new(big.Float).SetFloat64(0)
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
	orderPrice, _ := order.Price.Float64()
	for maxNode == nil || order.Amount.Cmp(bigZero) == 1 {
		if order.Type == Sell {
			maxNode = tree.Max()
		} else {
			maxNode = tree.Min()
		}
		if maxNode == nil || noMoreOrders {
			if order.Amount.Cmp(bigZero) == 1 {
				partialOrder = NewOrder(order.ID, order.Type, order.Amount.String(), order.Price.String())
				break
			} else {
				break
			}
		}
		if order.Type == Sell {
			if orderPrice > maxNode.Key {
				// fmt.Println("adding sellnode directly")
				noMoreOrders = true
				// return trades, noMoreOrders, nil, nil
				return noMoreOrders, nil, nil
			}
		} else {
			if orderPrice < maxNode.Key {
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
			if ele.Amount.Cmp(order.Amount) == 1 {
				negAmount := new(big.Float).Sub(bigZero, order.Amount)
				nodeData.updateVolume(negAmount)
				// trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				amount := new(big.Float).Sub(ele.Amount, order.Amount)
				// amount = math.Floor(amount*100000000) / 100000000
				ele.Amount = amount

				partialOrder = NewOrder(ele.ID, ele.Type, ele.Amount.String(), ele.Price.String())
				ordersProcessed = append(ordersProcessed, NewOrder(order.ID, order.Type, order.Amount.String(), order.Price.String()))

				maxNode.SetData(nodeData)

				order.Amount = new(big.Float).SetFloat64(0)
				noMoreOrders = true
				break
			}
			if ele.Amount.Cmp(order.Amount) == 0 {
				negAmount := new(big.Float).Sub(bigZero, order.Amount)
				nodeData.updateVolume(negAmount)

				ordersProcessed = append(ordersProcessed, NewOrder(ele.ID, ele.Type, ele.Amount.String(), ele.Price.String()))
				ordersProcessed = append(ordersProcessed, NewOrder(order.ID, order.Type, order.Amount.String(), order.Price.String()))

				countMatch++
				// trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				order.Amount = new(big.Float).SetFloat64(0)
				// orderComplete = true

				// ele.Amount = 0
				delete(ob.orders, ele.ID)

				break
			} else {
				countMatch++

				ordersProcessed = append(ordersProcessed, NewOrder(ele.ID, ele.Type, ele.Amount.String(), ele.Price.String()))

				negEleAmount := new(big.Float).Sub(bigZero, ele.Amount)
				nodeData.updateVolume(negEleAmount)

				// trades = append(trades, Trade{BuyOrderID: ele.ID, SellOrderID: order.ID, Amount: ele.Amount, Price: ele.Price})

				order.Amount = new(big.Float).Sub(order.Amount, ele.Amount)

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
