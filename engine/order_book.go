package engine

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/Pantelwar/binarytree"
)

// OrderBook type
type OrderBook struct {
	BuyTree         *binarytree.BinaryTree
	SellTree        *binarytree.BinaryTree
	orderLimitRange int
	orders          map[string]*OrderNode // orderID -> *Order (*list.Element.Value.(*Order))
}

// Book ...
type Book struct {
	Buys  []orderinfo `json:"buys"`
	Sells []orderinfo `json:"sells"`
}

type orderinfo struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

// // GetOrders ...
// func (ob *OrderBook) GetOrders() Book {
// 	buys := []orderinfo{}
// 	// var orderSideBuy []string
// 	ob.BuyTree.Root.InOrderTraverse(func(i float64) {
// 		// result = append(result, fmt.Sprintf("%#v", i))
// 		// fmt.Println("Key", i)
// 		node := ob.BuyTree.Root.SearchSubTree(i)
// 		node.Data.(*OrderType).Tree.Root.InOrderTraverse(func(i float64) {
// 			// result = append(result, fmt.Sprintf("%#v", i))
// 			// fmt.Println("    value", i)
// 			var b orderinfo
// 			// res := fmt.Sprintf("%#f -> ", i)
// 			b.Price = i
// 			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
// 			// fmt.Printf("subnode: %#v\n", subNode)
// 			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
// 			// res += fmt.Sprintf("%#f", subNode.Data.(*OrderNode).Volume)
// 			b.Amount = subNode.Data.(*OrderNode).Volume
// 			// fmt.Println("res b", res)
// 			// orderSideBuy = append(orderSideBuy, res)
// 			buys = append(buys, b)
// 		})
// 	})
// 	sells := []orderinfo{}
// 	ob.SellTree.Root.InOrderTraverse(func(i float64) {
// 		// result = append(result, fmt.Sprintf("%#v", i))
// 		// fmt.Println("Key lol", i)
// 		node := ob.SellTree.Root.SearchSubTree(i)
// 		node.Data.(*OrderType).Tree.Root.InOrderTraverse(func(i float64) {
// 			// result = append(result, fmt.Sprintf("%#v", i))
// 			// fmt.Println("    value", i)
// 			var b orderinfo
// 			// res := fmt.Sprintf("%#f -> ", i)
// 			b.Price = i
// 			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
// 			// fmt.Printf("subnode: %#v\n", subNode)
// 			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
// 			// res += fmt.Sprintf("%#f", subNode.Data.(*OrderNode).Volume)
// 			// fmt.Println("res", res)
// 			b.Amount = subNode.Data.(*OrderNode).Volume
// 			// fmt.Println("res b", res)
// 			// orderSideBuy = append(orderSideBuy, res)
// 			buys = append(sells, b)
// 		})
// 	})
// 	return Book{
// 		Buys:  buys,
// 		Sells: sells,
// 	}
// }

// MarshalJSON implements json.Marshaler interface
func (ob *OrderBook) MarshalJSON() ([]byte, error) {
	buys := []orderinfo{}
	// var orderSideBuy []string
	ob.BuyTree.Root.InOrderTraverse(func(i float64) {
		// result = append(result, fmt.Sprintf("%#v", i))
		// fmt.Println("Key", i)
		node := ob.BuyTree.Root.SearchSubTree(i)
		node.Data.(*OrderType).Tree.Root.InOrderTraverse(func(i float64) {
			// result = append(result, fmt.Sprintf("%#v", i))
			// fmt.Println("    value", i)
			var b orderinfo
			// res := fmt.Sprintf("%#f -> ", i)
			b.Price = i
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			// fmt.Printf("subnode: %#v\n", subNode)
			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
			// res += fmt.Sprintf("%#f", subNode.Data.(*OrderNode).Volume)
			b.Amount = subNode.Data.(*OrderNode).Volume
			// fmt.Println("res b", res)
			// orderSideBuy = append(orderSideBuy, res)
			buys = append(buys, b)
		})
	})

	sells := []orderinfo{}
	ob.SellTree.Root.InOrderTraverse(func(i float64) {
		// result = append(result, fmt.Sprintf("%#v", i))
		// fmt.Println("Key lol", i)
		node := ob.SellTree.Root.SearchSubTree(i)
		node.Data.(*OrderType).Tree.Root.InOrderTraverse(func(i float64) {
			// result = append(result, fmt.Sprintf("%#v", i))
			// fmt.Println("    value", i)
			var b orderinfo
			// res := fmt.Sprintf("%#f -> ", i)
			b.Price = i
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			// fmt.Printf("subnode: %#v\n", subNode)
			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
			// res += fmt.Sprintf("%#f", subNode.Data.(*OrderNode).Volume)
			// fmt.Println("res", res)
			b.Amount = subNode.Data.(*OrderNode).Volume
			// fmt.Println("res b", res)
			// orderSideBuy = append(orderSideBuy, res)
			buys = append(sells, b)
		})
	})

	// res := ob.GetOrders()
	return json.Marshal(
		&Book{
			Buys:  buys,
			Sells: sells,
		},
	)
}

// String implements Stringer interface
func (ob *OrderBook) String() string {
	result := ""
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
	// fmt.Println()
	// fmt.Println("sell orders")
	for _, o := range orderSideSell {
		// fmt.Println(o)
		result += o + "\n"
	}
	result += "------------------------------------------\n"

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
	// fmt.Println()
	// fmt.Println("buy orders")
	buys := ""
	for _, o := range orderSideBuy {
		// fmt.Println(o)
		buys = o + "\n" + buys
	}
	result += buys
	return result
}

// NewOrderBook Returns new order book
func NewOrderBook() *OrderBook {
	bTree := binarytree.NewBinaryTree()
	sTree := binarytree.NewBinaryTree()
	bTree.ToggleSplay(true)
	sTree.ToggleSplay(true)

	return &OrderBook{
		BuyTree:         bTree,
		SellTree:        sTree,
		orderLimitRange: 100,
		orders:          make(map[string]*OrderNode),
	}
}

// addBuyOrder a buy order to the order book
func (ob *OrderBook) addBuyOrder(order Order) {
	startPoint := float64(int(math.Ceil(order.Price)) / ob.orderLimitRange * ob.orderLimitRange)
	endPoint := startPoint + float64(ob.orderLimitRange)
	searchNodePrice := (startPoint + endPoint) / 2
	// fmt.Println("search node", startPoint, searchNodePrice, endPoint)
	node := ob.BuyTree.Root.SearchSubTree(searchNodePrice)
	var orderNode *OrderNode
	if node != nil {
		// fmt.Println("slab found", order.Price)
		subTree := node.Data.(*OrderType)
		subTreeNode := subTree.Tree.Root.SearchSubTree(order.Price)
		if subTreeNode != nil {
			// order.Amount += node.Datorder.Priceorder.Pricea.(Order).Amount
			// fmt.Println("node found", order.Price)
			subTreeNode.Data.(*OrderNode).UpdateVolume(order.Amount)
			subTreeNode.Data.(*OrderNode).AddOrder(order)
			orderNode = subTreeNode.Data.(*OrderNode)
		} else {
			// fmt.Println("not found", order.Price)
			orderNode = subTree.AddOrderInQueue(order)
		}
		// return
	} else {
		// fmt.Println("adding new slab", order.Price)
		orderTypeObj := NewOrderType(order.Type)
		orderNode = orderTypeObj.AddOrderInQueue(order)
		ob.BuyTree.Insert(searchNodePrice, orderTypeObj)
	}
	// fmt.Println("ors", orderNode)
	ob.orders[order.ID] = orderNode
}

// addSellOrder a buy order to the order book
func (ob *OrderBook) addSellOrder(order Order) {
	startPoint := float64(int(math.Ceil(order.Price)) / ob.orderLimitRange * ob.orderLimitRange)
	endPoint := startPoint + float64(ob.orderLimitRange)
	searchNodePrice := (startPoint + endPoint) / 2
	// fmt.Println("search node", startPoint, searchNodePrice, endPoint)
	node := ob.SellTree.Root.SearchSubTree(searchNodePrice)
	var orderNode *OrderNode
	if node != nil {
		// fmt.Println("slab found", order.Price)
		subTree := node.Data.(*OrderType)
		subTreeNode := subTree.Tree.Root.SearchSubTree(order.Price)
		if subTreeNode != nil {
			// order.Amount += node.Datorder.Priceorder.Pricea.(Order).Amount
			// fmt.Println("node found", order.Price)
			subTreeNode.Data.(*OrderNode).UpdateVolume(order.Amount)
			subTreeNode.Data.(*OrderNode).AddOrder(order)
			orderNode = subTreeNode.Data.(*OrderNode)
		} else {
			// fmt.Println("not found", order.Price)
			orderNode = subTree.AddOrderInQueue(order)
		}
		// return
	} else {
		// fmt.Println("adding new slab", order.Price)
		orderTypeObj := NewOrderType(order.Type)
		orderNode = orderTypeObj.AddOrderInQueue(order)
		ob.SellTree.Insert(searchNodePrice, orderTypeObj)
	}
	ob.orders[order.ID] = orderNode
}

func (ob *OrderBook) removeBuyOrder(key float64) *binarytree.BinaryNode {
	return ob.BuyTree.Root.Remove(key)
}

func (ob *OrderBook) removeSellOrder(key float64) *binarytree.BinaryNode {
	return ob.SellTree.Root.Remove(key)
}

// func (ob *OrderBook) removeOrder(order *Order, index int) error {
// 	startPoint := float64(int(math.Ceil(order.Price)) / ob.orderLimitRange * ob.orderLimitRange)
// 	endPoint := startPoint + float64(ob.orderLimitRange)
// 	searchNodePrice := (startPoint + endPoint) / 2
// 	// fmt.Println("search node", startPoint, searchNodePrice, endPoint)
// 	node := ob.BuyTree.Root.SearchSubTree(searchNodePrice)
// 	// var orderNode *OrderNode
// 	if node != nil {
// 		// fmt.Println("slab found", order.Price)
// 		subTree := node.Data.(*OrderType)
// 		subTreeNode := subTree.Tree.Root.SearchSubTree(order.Price)
// 		if subTreeNode != nil {
// 			fmt.Println("Found node to remove")
// 			subTreeNode.Data.(*OrderNode).UpdateVolume(-order.Amount)
// 			subTreeNode.Data.(*OrderNode).Orders = append(subTreeNode.Data.(*OrderNode).Orders[:index], subTreeNode.Data.(*OrderNode).Orders[index+1:]...)
// 			if len(subTreeNode.Data.(*OrderNode).Orders) == 0 {
// 				// orderNode = nil
// 				n := subTree.Tree.Root.Remove(order.Price)
// 				subTree.Tree.Root = n
// 			}
// 		} else {
// 			fmt.Println("nothing")
// 		}
// 		// return
// 	} else {
// 		fmt.Println("nothing #2")
// 	}
// 	return nil
// }
