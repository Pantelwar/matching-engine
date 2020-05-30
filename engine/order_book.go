package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"

	"strconv"

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
	Price  string `json:"price"`
	Amount string `json:"amount"`
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
			b.Price = strconv.FormatFloat(i, 'f', -1, 64)
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			// fmt.Printf("subnode: %#v\n", subNode)
			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
			// res += fmt.Sprintf("%#f", subNode.Data.(*OrderNode).Volume)
			b.Amount = strconv.FormatFloat(subNode.Data.(*OrderNode).Volume, 'f', -1, 64)
			fmt.Println("res b", b)
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
			b.Price = strconv.FormatFloat(i, 'f', -1, 64)
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			// fmt.Printf("subnode: %#v\n", subNode)
			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
			// res += fmt.Sprintf("%#f", subNode.Data.(*OrderNode).Volume)
			// fmt.Println("res", res)
			b.Amount = strconv.FormatFloat(subNode.Data.(*OrderNode).Volume, 'f', -1, 64)
			fmt.Println("res b", b)
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
			res := strconv.FormatFloat(i, 'f', -1, 64) + " -> "
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			// fmt.Printf("subnode: %#v\n", subNode)
			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
			res += strconv.FormatFloat(subNode.Data.(*OrderNode).Volume, 'f', -1, 64)
			// fmt.Println("res", res)
			orderSideSell = append(orderSideSell, res)
		})
	})
	// fmt.Println()
	// fmt.Println("sell orders")
	sells := ""
	for _, o := range orderSideSell {
		// fmt.Println(o)
		sells = o + "\n" + sells
	}

	result = sells + "------------------------------------------\n"

	var orderSideBuy []string
	ob.BuyTree.Root.InOrderTraverse(func(i float64) {
		// result = append(result, fmt.Sprintf("%#v", i))
		// fmt.Println("Key", i)
		node := ob.BuyTree.Root.SearchSubTree(i)
		node.Data.(*OrderType).Tree.Root.InOrderTraverse(func(i float64) {
			// result = append(result, fmt.Sprintf("%#v", i))
			// fmt.Println("    value", i)
			res := strconv.FormatFloat(i, 'f', -1, 64) + " -> "
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			// fmt.Printf("subnode: %#v\n", subNode)
			// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
			res += strconv.FormatFloat(subNode.Data.(*OrderNode).Volume, 'f', -1, 64)
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
			// subTreeNode.Data.(*OrderNode).updateVolume(order.Amount)
			subTreeNode.Data.(*OrderNode).addOrder(order)
			orderNode = subTreeNode.Data.(*OrderNode)
		} else {
			// fmt.Println("not found", order.Price)
			orderNode, _ = subTree.AddOrderInQueue(order)
		}
		// return
	} else {
		// fmt.Println("adding new slab", order.Price)
		orderTypeObj := NewOrderType(order.Type)
		orderNode, _ = orderTypeObj.AddOrderInQueue(order)
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
			// subTreeNode.Data.(*OrderNode).updateVolume(order.Amount)
			subTreeNode.Data.(*OrderNode).addOrder(order)
			orderNode = subTreeNode.Data.(*OrderNode)
		} else {
			// fmt.Println("not found", order.Price)
			orderNode, _ = subTree.AddOrderInQueue(order)
		}
		// return
	} else {
		// fmt.Println("adding new slab", order.Price)
		orderTypeObj := NewOrderType(order.Type)
		orderNode, _ = orderTypeObj.AddOrderInQueue(order)
		ob.SellTree.Insert(searchNodePrice, orderTypeObj)
	}
	ob.orders[order.ID] = orderNode
}

func (ob *OrderBook) removeBuyNode(key float64) error {
	node := ob.BuyTree.Root.Remove(key)
	ob.BuyTree.Root = node
	return nil
}

func (ob *OrderBook) removeSellNode(key float64) error {
	node := ob.SellTree.Root.Remove(key)
	ob.SellTree.Root = node
	return nil
}

func (ob *OrderBook) removeOrder(order *Order) error {
	startPoint := float64(int(math.Ceil(order.Price)) / ob.orderLimitRange * ob.orderLimitRange)
	endPoint := startPoint + float64(ob.orderLimitRange)
	searchNodePrice := (startPoint + endPoint) / 2
	// fmt.Println("search node", startPoint, searchNodePrice, endPoint)
	var node *binarytree.BinaryNode
	if order.Type == Buy {
		node = ob.BuyTree.Root.SearchSubTree(searchNodePrice)
	} else {
		node = ob.SellTree.Root.SearchSubTree(searchNodePrice)
	}
	if node != nil {
		// fmt.Println("slab found", order.Price)
		subTree := node.Data.(*OrderType)
		subTreeNode := subTree.Tree.Root.SearchSubTree(order.Price)
		if subTreeNode != nil {
			fmt.Println("Found node to remove")
			// subTreeNode.Data.(*OrderNode).updateVolume(-order.Amount)
			// subTreeNode.Data.(*OrderNode).Orders = append(subTreeNode.Data.(*OrderNode).Orders[:index], subTreeNode.Data.(*OrderNode).Orders[index+1:]...)
			// if len(subTreeNode.Data.(*OrderNode).Orders) == 0 {
			// orderNode = nil
			n := subTree.Tree.Root.Remove(order.Price)
			subTree.Tree.Root = n
			// }
		} else {
			return errors.New("no Order found")
			// fmt.Println("nothing")
		}
		// return
	} else {
		return errors.New("no Order found")
	}
	return nil
}
