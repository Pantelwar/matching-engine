package binarytree

import "fmt"

// BinaryTree ...
type BinaryTree struct {
	Root *BinaryNode
}

// NewBinaryTree ...
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

// Insert ...
func (t *BinaryTree) Insert(price float64, amount float64) *BinaryTree {
	if t.Root == nil {
		t.Root = NewBinaryNode(price, amount)
	} else {
		t.Root.Insert(price, amount)
	}
	return t
}

func (t *BinaryTree) LessThan(n *BinaryNode, price float64) *BinaryNode {
	if n == nil {
		return n
	}
	fmt.Println("...", n.Price, price)
	if price < n.Price {
		return t.LessThan(n.Left, price)
	}
	// if price > n.Price {
	// 	return t.SearchSubTree(n.Right, price)
	// }
	return n.Left
}

func (t *BinaryTree) GreatThan(n *BinaryNode, price float64) *BinaryNode {
	if n == nil {
		return n
	}
	if price > n.Price {
		return t.GreatThan(n.Right, price)
	}
	return n
}

func (t *BinaryTree) Max() *BinaryNode {
	n := t.Root
	if n == nil {
		return nil
	}
	for {
		if n.Right == nil {
			return n
		}
		n = n.Right
	}
}
