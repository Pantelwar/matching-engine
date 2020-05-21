package engine

import (
	"math"

	"github.com/Pantelwar/binarytree"
)

// CancelOrder remove the order from book and returns
func (ob *OrderBook) CancelOrder(id string) *Order {
	orderNode := ob.orders[id]
	// fmt.Printf("orderNode: %v\n", orderNode.Orders)
	// orderNode.Orders = []*Order{}
	for i, order := range orderNode.Orders {
		if order.ID == id {
			// ob.orders[id].AddOrder(*order)
			// ob.removeOrder(order, i)
			orderNode.UpdateVolume(-order.Amount)
			orderNode.Orders = append(orderNode.Orders[:i], orderNode.Orders[i+1:]...)
			if len(orderNode.Orders) == 0 {
				// orderNode = nil
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
				// var orderNode *OrderNode
				if node != nil {
					// fmt.Println("slab found", order.Price)
					subTree := node.Data.(*OrderType)
					subTreeNode := subTree.Tree.Root.SearchSubTree(order.Price) // same as orderNode
					if subTreeNode != nil {
						n := subTree.Tree.Root.Remove(order.Price)
						subTree.Tree.Root = n
					}
				}
			}
			delete(ob.orders, id)
			return order
		}
	}
	return nil
}
