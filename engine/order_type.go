package engine

import (
	"github.com/Pantelwar/binarytree"
)

type OrderType struct {
	Tree *binarytree.BinaryTree
	Type string
}

func NewOrderType(orderSide string) *OrderType {
	bTree := binarytree.NewBinaryTree()
	return &OrderType{Tree: bTree, Type: orderSide}
}

func (ot *OrderType) AddOrderInQueue(order Order) error {
	// arr := []*Order{&order}
	orderNode := NewOrderNode()
	orderNode.Orders = append(orderNode.Orders, &order)
	orderNode.Volume = order.Amount
	ot.Tree.Insert(order.Price, orderNode)
	return nil
}
