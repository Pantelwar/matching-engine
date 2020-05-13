package engine

import (
	"fmt"
	"math"

	"github.com/Pantelwar/binarytree"
)

// OrderBook type
type OrderBook struct {
	BuyTree         *binarytree.BinaryTree
	SellTree        *binarytree.BinaryTree
	OrderLimitRange int
}

// NewOrderBook Returns new order book
func NewOrderBook2() *OrderBook {
	return &OrderBook{
		BuyTree:         binarytree.NewBinaryTree(),
		SellTree:        binarytree.NewBinaryTree(),
		OrderLimitRange: 100,
	}
}

// Process executes limit process
func (ob *OrderBook) Process(order Order, orderside string) []Trade {
	if orderside == "buy" {
		return ob.processOrderB(order)
	}
	return ob.processOrderS(order)
}

// addBuyOrder a buy order to the order book
func (ob *OrderBook) addBuyOrder(order Order) {
	startPoint := float64(int(math.Ceil(order.Price)) / ob.OrderLimitRange * ob.OrderLimitRange)
	endPoint := startPoint + float64(ob.OrderLimitRange)
	searchNodePrice := (startPoint + endPoint) / 2
	// fmt.Println("search node", startPoint, searchNodePrice, endPoint)
	node := ob.BuyTree.Root.SearchSubTree(searchNodePrice)
	if node != nil {
		// fmt.Println("slab found", order.Price)
		subTree := node.Data.(*OrderType)
		subTreeNode := subTree.Tree.Root.SearchSubTree(order.Price)
		if subTreeNode != nil {
			// order.Amount += node.Datorder.Priceorder.Pricea.(Order).Amount
			// fmt.Println("node found", order.Price)
			subTreeNode.Data.(*OrderNode).UpdateVolume(order.Amount)
			subTreeNode.Data.(*OrderNode).AddOrder(order)
		} else {
			// fmt.Println("not found", order.Price)
			subTree.AddOrderInQueue(order)
		}
		// return
	} else {
		// fmt.Println("adding new slab", order.Price)
		orderTypeObj := NewOrderType("buy")
		orderTypeObj.AddOrderInQueue(order)
		ob.BuyTree.Insert(searchNodePrice, orderTypeObj)
	}
}

// addSellOrder a buy order to the order book
func (ob *OrderBook) addSellOrder(order Order) {
	startPoint := float64(int(math.Ceil(order.Price)) / ob.OrderLimitRange * ob.OrderLimitRange)
	endPoint := startPoint + float64(ob.OrderLimitRange)
	searchNodePrice := (startPoint + endPoint) / 2
	// fmt.Println("search node", startPoint, searchNodePrice, endPoint)
	node := ob.SellTree.Root.SearchSubTree(searchNodePrice)
	if node != nil {
		// fmt.Println("slab found", order.Price)
		subTree := node.Data.(*OrderType)
		subTreeNode := subTree.Tree.Root.SearchSubTree(order.Price)
		if subTreeNode != nil {
			// order.Amount += node.Datorder.Priceorder.Pricea.(Order).Amount
			// fmt.Println("node found", order.Price)
			subTreeNode.Data.(*OrderNode).UpdateVolume(order.Amount)
			subTreeNode.Data.(*OrderNode).AddOrder(order)
		} else {
			// fmt.Println("not found", order.Price)
			subTree.AddOrderInQueue(order)
		}
		// return
	} else {
		// fmt.Println("adding new slab", order.Price)
		orderTypeObj := NewOrderType("sell")
		orderTypeObj.AddOrderInQueue(order)
		ob.SellTree.Insert(searchNodePrice, orderTypeObj)
	}
}

// processOrderS a limit sell order
func (ob *OrderBook) processOrderS(order Order) []Trade {
	trades := make([]Trade, 0, 1)
	maxNode := ob.BuyTree.Max()
	if maxNode == nil {
		// fmt.Println("adding sellnode")
		ob.addSellOrder(order)
		return trades
	}
	// fmt.Println("processing sell node", maxNode.Key)
	// fmt.Printf("processing sell node: %#v\n", maxNode)
	count := 0
	noMoreOrders := false
	for maxNode == nil || order.Amount > 0 {
		count++
		// if count == 10 {
		// 	fmt.Println("breaking abnormal buy")
		// 	break
		// }
		maxNode = ob.BuyTree.Max()
		// fmt.Println()
		// fmt.Println("order amount", order.Amount, maxNode)
		if maxNode == nil || noMoreOrders {
			if order.Amount > 0 {
				ob.addSellOrder(order)
				// fmt.Println("breaking")
				break
			} else {
				break
			}
		}
		// if maxNode == nil && order.Amount > 0 {
		// 	ob.addSellOrder(order)
		// 	fmt.Println("breaking")
		// 	break
		// }

		if order.Price > maxNode.Key {
			// fmt.Println("adding sellnode directly")
			ob.addSellOrder(order)
			return trades
		}
		// fmt.Println()
		// fmt.Println("Execution processLimitSell")
		var t []Trade
		t, noMoreOrders = ob.processLimitSell(&order, maxNode.Data.(*OrderType).Tree)
		trades = append(trades, t...)
		// fmt.Println("Complete processLimitSell")
		// fmt.Println()

		if maxNode.Data.(*OrderType).Tree.Root == nil {
			// fmt.Println("removing main node")
			node := ob.removeBuyOrder(maxNode.Key)
			ob.BuyTree.Root = node
			// fmt.Printf("removing main node #1: %#v\n", val)

		}
		// fmt.Println("TREEEEEEEEEEEEEEE", maxNode.Data.(*OrderType).Tree.Root, maxNode.Key)
		// print(os.Stdout, ob.BuyTree.Root, 0, 'M')

	}
	return trades
}

func (ob *OrderBook) processLimitSell(order *Order, buyTree *binarytree.BinaryTree) ([]Trade, bool) {
	trades := make([]Trade, 0, 1)

	maxNode := buyTree.Max()
	// fmt.Printf("main check smaxNode: %#v\n", maxNode) //.Data) //.(*engine.OrderNode))
	noMoreOrders := false
	if maxNode == nil {
		// addSellOrder(*order, sellTree)
		return trades, noMoreOrders
	}
	countAdd := 0.0
	for maxNode == nil || order.Amount > 0 {
		// fmt.Println()
		// fmt.Println("Loooping again")
		maxNode = buyTree.Max()
		if maxNode == nil || noMoreOrders {
			if order.Amount > 0 {
				break
			} else {
				break
			}
		}
		// fmt.Printf("maxNode: %#v\n", maxNode) //.Data) //.(*OrderNode))
		nodeData := maxNode.Data.(*OrderNode) //([]*Order)
		nodeOrders := nodeData.Orders         //([]*Order)
		countMatch := 0
		for _, ele := range nodeOrders {
			// fmt.Printf("orders %#v\n", ele)
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

		// fmt.Println("countMatch", countMatch, len(nodeOrders))
		if len(nodeOrders) == countMatch {
			node := buyTree.Root.Remove(maxNode.Key) // ob.removeBuyOrder(maxNode.Key, buyTree)
			// fmt.Printf("node removed: %#v %#v\n", node, maxNode)
			buyTree.Root = node
		}

		// fmt.Println("countMatch", countMatch, nodeOrders[countMatch:])
		nodeData.Orders = nodeOrders[countMatch:]
		maxNode.SetData(nodeData)

		// continue
	}
	return trades, noMoreOrders
}

// Process a limit sell order
func (ob *OrderBook) processOrderB(order Order) []Trade {
	trades := make([]Trade, 0, 1)
	maxNode := ob.SellTree.Min()
	if maxNode == nil {
		// fmt.Println("adding buynode")
		ob.addBuyOrder(order)
		return trades
	}
	// fmt.Println("processing sell node", maxNode.Key)
	// fmt.Printf("processing sell node: %#v\n", maxNode)
	count := 0
	noMoreOrders := false
	for maxNode == nil || order.Amount > 0 {
		count++
		// if count == 10 {
		// 	fmt.Println("breaking abnormal sell")
		// 	break
		// }
		maxNode = ob.SellTree.Min()
		// fmt.Println()
		// fmt.Println("order amount", order.Amount, maxNode)
		if maxNode == nil || noMoreOrders {
			if order.Amount > 0 {
				ob.addBuyOrder(order)
				// fmt.Println("breaking")
				break
			} else {
				break
			}
			// ob.addBuyOrder(order)
			// fmt.Println("breaking")
			// break
		}

		if order.Price < maxNode.Key {
			// fmt.Println("adding sellnode directly")
			ob.addBuyOrder(order)
			return trades
		}
		// fmt.Println()
		// fmt.Println("Execution processLimitBuy")
		var t []Trade
		t, noMoreOrders = ob.processLimitBuy(&order, maxNode.Data.(*OrderType).Tree)
		trades = append(trades, t...)
		// fmt.Println("Complete processLimitBuy")
		// fmt.Println()

		if maxNode.Data.(*OrderType).Tree.Root == nil {
			// fmt.Println("removing main node")
			node := ob.removeSellOrder(maxNode.Key)
			ob.SellTree.Root = node
			// fmt.Printf("removing main node #1: %#v\n", val)

		}
		// fmt.Println("TREEEEEEEEEEEEEEE", maxNode.Data.(*OrderType).Tree.Root, maxNode.Key)
		// print(os.Stdout, ob.BuyTree.Root, 0, 'M')

	}
	return trades
}

func (ob *OrderBook) processLimitBuy(order *Order, buyTree *binarytree.BinaryTree) ([]Trade, bool) {
	trades := make([]Trade, 0, 1)

	maxNode := buyTree.Min()
	// fmt.Printf("main check smaxNode: %#v\n", maxNode) //.Data) //.(*engine.OrderNode))
	noMoreOrders := false
	if maxNode == nil {
		// addSellOrder(*order, sellTree)
		return trades, noMoreOrders
	}
	countAdd := 0.0
	for maxNode == nil || order.Amount > 0 {
		// fmt.Println()
		// fmt.Println("Loooping again")
		maxNode = buyTree.Min()
		if maxNode == nil || noMoreOrders {
			// if order.Amount > 0 {
			// 	break
			// } else {
			// 	break
			// }
			// addSellOrder(*order, sellTree)
			break
		}

		// fmt.Printf("maxNode: %#v\n", maxNode) //.Data) //.(*OrderNode))
		nodeData := maxNode.Data.(*OrderNode) //([]*Order)
		nodeOrders := nodeData.Orders         //([]*Order)
		countMatch := 0
		for _, ele := range nodeOrders {
			// fmt.Printf("orders %#v\n", ele)
			// fmt.Println("ele.Price", ele.Price)
			if ele.Price > order.Price {
				// fmt.Println("No more orders")
				noMoreOrders = true
				break
			}
			countAdd += ele.Amount
			// fmt.Println(order.Price, "AMOUNT", order.Amount, ele, ele.Amount > order.Amount)
			if ele.Amount > order.Amount {
				nodeData.UpdateVolume(-order.Amount)

				trades = append(trades, Trade{SellOrderID: ele.ID, BuyOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

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
				trades = append(trades, Trade{SellOrderID: ele.ID, BuyOrderID: order.ID, Amount: order.Amount, Price: ele.Price})

				ele.Amount = 0
				// node := book.removeBuyOrder(maxNode.Key)
				// book.BuyOrders.Tree.Root = node
				// book.addSellOrder(order)

				// break
			} else {
				// fmt.Println("Removing Node and continue")
				countMatch++

				nodeData.UpdateVolume(-ele.Amount)

				trades = append(trades, Trade{SellOrderID: ele.ID, BuyOrderID: order.ID, Amount: ele.Amount, Price: ele.Price})

				order.Amount -= ele.Amount
				ele.Amount = 0.0
				// order.Amount = math.Floor(order.Amount*100000000) / 100000000
				// node := book.removeBuyOrder(maxNode.Key)
				// book.BuyOrders.Tree.Root = node

			}
		}

		// fmt.Println("countMatch", countMatch, len(nodeOrders))
		if len(nodeOrders) == countMatch {
			node := buyTree.Root.Remove(maxNode.Key) // ob.removeBuyOrder(maxNode.Key, buyTree)
			// fmt.Printf("node removed: %#v %#v\n", node, maxNode)
			buyTree.Root = node
		}

		// fmt.Println("countMatch", countMatch, nodeOrders[countMatch:])
		nodeData.Orders = nodeOrders[countMatch:]
		maxNode.SetData(nodeData)

		// continue
	}
	return trades, noMoreOrders
}

func (ob *OrderBook) removeBuyOrder(key float64) *binarytree.BinaryNode {
	return ob.BuyTree.Root.Remove(key)
}

func (ob *OrderBook) removeSellOrder(key float64) *binarytree.BinaryNode {
	return ob.SellTree.Root.Remove(key)
}

// Print displays book
func (ob *OrderBook) Print() {
	var orderSideBuy []string
	ob.BuyTree.Root.InOrderTraverse(func(i float64) {
		// result = append(result, fmt.Sprintf("%#v", i))
		// fmt.Println("Key", i)
		node := ob.BuyTree.Root.SearchSubTree(i)
		node.Data.(*OrderType).Tree.Root.InOrderTraverse(func(i float64) {
			// result = append(result, fmt.Sprintf("%#v", i))
			// fmt.Println("    value", i)
			res := fmt.Sprintf("%#f -> ", i)
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			// fmt.Printf("subnode: %#v\n", subNode)
			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
			res += fmt.Sprintf("%#f", subNode.Data.(*OrderNode).Volume)
			// fmt.Println("res b", res)
			orderSideBuy = append(orderSideBuy, res)
		})
	})
	fmt.Println()
	fmt.Println("buy orders")
	for _, o := range orderSideBuy {
		fmt.Println(o)
	}

	var orderSideSell []string
	ob.SellTree.Root.InOrderTraverse(func(i float64) {
		// result = append(result, fmt.Sprintf("%#v", i))
		// fmt.Println("Key lol", i)
		node := ob.SellTree.Root.SearchSubTree(i)
		node.Data.(*OrderType).Tree.Root.InOrderTraverse(func(i float64) {
			// result = append(result, fmt.Sprintf("%#v", i))
			// fmt.Println("    value", i)
			res := fmt.Sprintf("%#f -> ", i)
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			// fmt.Printf("subnode: %#v\n", subNode)
			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
			res += fmt.Sprintf("%#f", subNode.Data.(*OrderNode).Volume)
			// fmt.Println("res", res)
			orderSideSell = append(orderSideSell, res)
		})
	})
	fmt.Println()
	fmt.Println("sell orders")
	for _, o := range orderSideSell {
		fmt.Println(o)
	}
}
