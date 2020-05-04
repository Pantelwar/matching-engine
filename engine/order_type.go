package engine

import (
	"matching-engine/engine/binarytree"
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
	arr := []*Order{&order}
	ot.Tree.Insert(order.Price, arr)
	return nil
}
