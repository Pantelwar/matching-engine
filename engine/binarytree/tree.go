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
func (t *BinaryTree) Insert(key float64, data interface{}) *BinaryTree {
	if t.Root == nil {
		t.Root = NewBinaryNode(key, data)
	} else {
		t.Root.Insert(key, data)
	}
	return t
}

func (t *BinaryTree) LessThan(n *BinaryNode, key float64) *BinaryNode {
	if n == nil {
		return n
	}
	fmt.Println("...", n.Key, key)
	if key < n.Key {
		return t.LessThan(n.Left, key)
	}
	// if key > n.Key {
	// 	return t.SearchSubTree(n.Right, key)
	// }
	return n.Left
}

func (t *BinaryTree) GreatThan(n *BinaryNode, key float64) *BinaryNode {
	if n == nil {
		return n
	}
	if key > n.Key {
		return t.GreatThan(n.Right, key)
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
