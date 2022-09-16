package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/Pantelwar/binarytree"
	"github.com/Pantelwar/matching-engine/util"
)

// OrderBook type
type OrderBook struct {
	BuyTree         *binarytree.BinaryTree
	SellTree        *binarytree.BinaryTree
	orderLimitRange int
	orders          map[string]*OrderNode // orderID -> *Order (*list.Element.Value.(*Order))
	mutex           *sync.Mutex
}

// Book ...
type Book struct {
	Buys  []orderinfo `json:"buys"`
	Sells []orderinfo `json:"sells"`
}

type orderinfo struct {
	Price  *util.StandardBigDecimal `json:"price"`
	Amount *util.StandardBigDecimal `json:"amount"`
}

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
			b.Price = util.NewDecimalFromFloat(i)
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
			b.Price = util.NewDecimalFromFloat(i)
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

// BookArray ...
type BookArray struct {
	Buys  [][]string `json:"buys"`
	Sells [][]string `json:"sells"`
}

// GetOrders implements json.Marshaler interface
func (ob *OrderBook) GetOrders(limit int64) *BookArray {
	buys := [][]string{}
	ob.BuyTree.Root.InReverseOrderTraverse(func(i float64) {
		node := ob.BuyTree.Root.SearchSubTree(i)
		node.Data.(*OrderType).Tree.Root.InReverseOrderTraverse(func(i float64) {
			if int64(len(buys)) >= limit && limit != 0 {
				return
			}
			var b []string
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			if subNode != nil {
				price := strconv.FormatFloat(i, 'f', -1, 64)
				b = append(b, price)

				amount := subNode.Data.(*OrderNode).Volume
				b = append(b, amount.String())
				buys = append(buys, b)
			}
		})
	})

	sells := [][]string{}
	ob.SellTree.Root.InOrderTraverse(func(i float64) {
		node := ob.SellTree.Root.SearchSubTree(i)
		node.Data.(*OrderType).Tree.Root.InOrderTraverse(func(i float64) {
			if int64(len(sells)) >= limit && limit != 0 {
				return
			}
			var b []string
			subNode := node.Data.(*OrderType).Tree.Root.SearchSubTree(i)
			if subNode != nil {
				price := strconv.FormatFloat(i, 'f', -1, 64)
				b = append(b, price)

				amount := subNode.Data.(*OrderNode).Volume
				b = append(b, amount.String())
				sells = append(sells, b)
			}
		})
	})

	// res := ob.GetOrders()
	return &BookArray{
		Buys:  buys,
		Sells: sells,
	}
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
			if subNode != nil {
				// fmt.Printf("subnode: %#v\n", subNode)
				// fmt.Printf("volume:%#v, %#v\n\n", subNode.Data.(*OrderNode).Volume, len(subNode.Data.(*OrderNode).Orders))
				vol := subNode.Data.(*OrderNode).Volume.Float64()
				res += strconv.FormatFloat(vol, 'f', -1, 64) //subNode.Data.(*OrderNode).Volume.String() // strings.Trim(subNode.Data.(*OrderNode).Volume.String(), "0")
				// fmt.Println("res", res)
				orderSideSell = append(orderSideSell, res)
			}
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
			if subNode != nil {
				// fmt.Printf("subnode: %#v\n", subNode)
				vol := subNode.Data.(*OrderNode).Volume.Float64()
				res += strconv.FormatFloat(vol, 'f', -1, 64) //subNode.Data.(*OrderNode).Volume.String() // strings.Trim(subNode.Data.(*OrderNode).Volume.String(), "0")
				// fmt.Println("res b", res)
				orderSideBuy = append(orderSideBuy, res)
			}
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
		orderLimitRange: 200000000,
		orders:          make(map[string]*OrderNode),
		mutex:           &sync.Mutex{},
	}
}

// addBuyOrder a buy order to the order book
func (ob *OrderBook) addBuyOrder(order Order) {
	orderPrice := order.Price.Float64()
	startPoint := float64(int(math.Ceil(orderPrice)) / ob.orderLimitRange * ob.orderLimitRange)
	endPoint := startPoint + float64(ob.orderLimitRange)
	searchNodePrice := (startPoint + endPoint) / 2
	// fmt.Println("search node", startPoint, searchNodePrice, endPoint)
	node := ob.BuyTree.Root.SearchSubTree(searchNodePrice)
	var orderNode *OrderNode
	if node != nil {
		// fmt.Println("slab found", order.Price)
		subTree := node.Data.(*OrderType)
		subTreeNode := subTree.Tree.Root.SearchSubTree(orderPrice)
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
	ob.mutex.Lock()
	ob.orders[order.ID] = orderNode
	ob.mutex.Unlock()
}

// addSellOrder a buy order to the order book
func (ob *OrderBook) addSellOrder(order Order) {
	orderPrice := order.Price.Float64()
	startPoint := float64(int(math.Ceil(orderPrice)) / ob.orderLimitRange * ob.orderLimitRange)
	endPoint := startPoint + float64(ob.orderLimitRange)
	searchNodePrice := (startPoint + endPoint) / 2
	// fmt.Println("search node", startPoint, searchNodePrice, endPoint)
	node := ob.SellTree.Root.SearchSubTree(searchNodePrice)
	var orderNode *OrderNode
	if node != nil {
		// fmt.Println("slab found", order.Price)
		subTree := node.Data.(*OrderType)
		subTreeNode := subTree.Tree.Root.SearchSubTree(orderPrice)
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
	ob.mutex.Lock()
	ob.orders[order.ID] = orderNode
	ob.mutex.Unlock()
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
	orderPrice := order.Price.Float64()
	startPoint := float64(int(math.Ceil(orderPrice)) / ob.orderLimitRange * ob.orderLimitRange)
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
		subTreeNode := subTree.Tree.Root.SearchSubTree(orderPrice)
		if subTreeNode != nil {
			fmt.Println("Found node to remove")
			// subTreeNode.Data.(*OrderNode).updateVolume(-order.Amount)
			// subTreeNode.Data.(*OrderNode).Orders = append(subTreeNode.Data.(*OrderNode).Orders[:index], subTreeNode.Data.(*OrderNode).Orders[index+1:]...)
			// if len(subTreeNode.Data.(*OrderNode).Orders) == 0 {
			// orderNode = nil
			n := subTree.Tree.Root.Remove(orderPrice)
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
