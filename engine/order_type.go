package engine

import (
	"errors"

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
func (ot *OrderType) AddOrderInQueue(order Order) (*OrderNode, error) {
	if ot.Type != order.Type {
		return nil, errors.New("Invalid order type")
	}
	orderNode := NewOrderNode()
	orderNode.Orders = append(orderNode.Orders, &order)
	orderNode.Volume = order.Amount
	ot.Tree.Insert(order.Price, orderNode)
	return orderNode, nil
}
