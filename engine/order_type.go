package engine

import (
	"github.com/Pantelwar/binarytree"
)

// OrderType defines tree side
type OrderType struct {
	Tree *binarytree.BinaryTree
	Type Side `json:"type"`
}

// NewOrderType returns OrderType struct
func NewOrderType(orderSide Side) *OrderType {
	bTree := binarytree.NewBinaryTree()
	bTree.ToggleSplay(true)
	return &OrderType{Tree: bTree, Type: orderSide}
}

// AddOrderInQueue adds order to the tree
func (ot *OrderType) AddOrderInQueue(order Order) error {
	orderNode := NewOrderNode()
	orderNode.Orders = append(orderNode.Orders, &order)
	orderNode.Volume = order.Amount
	ot.Tree.Insert(order.Price, orderNode)
	return nil
}
