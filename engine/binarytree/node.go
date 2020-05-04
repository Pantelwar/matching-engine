package binarytree

import (
	"fmt"
	"io"
)

// type Item generic.Type

// BinaryNode ...
type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Key   float64
	Data  interface{}
}

// NewBinaryNode ...
func NewBinaryNode(key float64, data interface{}) *BinaryNode {
	return &BinaryNode{Key: key, Data: data, Left: nil, Right: nil}
}

// Insert ...
func (n *BinaryNode) Insert(key float64, data interface{}) {
	if n == nil {
		return
	} else if key < n.Key {
		// fmt.Println(" key < n.Key", key, n.Key)
		if n.Left == nil {
			n.Left = NewBinaryNode(key, data)
		} else {
			n.Left.Insert(key, data)
		}
	} else if key > n.Key {
		// fmt.Println(" key > n.Key", key, n.Key)
		if n.Right == nil {
			n.Right = NewBinaryNode(key, data)
		} else {
			n.Right.Insert(key, data)
		}
	} else {
		// fmt.Println(" key = n.Key", key, n.Key)

	}
}

func (n *BinaryNode) Print(w io.Writer, ns int, ch rune) {
	if n == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v -> %f\n", ch, n.Key, n.Data)
	n.Left.Print(w, ns+2, 'L')
	n.Right.Print(w, ns+2, 'R')

	// print(w, n.Left, ns+2, 'L')
	// print(w, n.Right, ns+2, 'R')
}

func (n *BinaryNode) PreOrderTraverse(f func(float64)) {
	if n != nil {
		n.Left.PreOrderTraverse(f)
		n.Right.PreOrderTraverse(f)
		f(n.Key)
	}
}

func (n *BinaryNode) Remove(key float64) *BinaryNode {
	if n == nil {
		return nil
	}
	if key < n.Key {
		n.Left = n.Left.Remove(key)
		return n
	}
	if key > n.Key {
		n.Right = n.Right.Remove(key)
		return n
	}
	// key == n.Key
	if n.Left == nil && n.Right == nil {
		n = nil
		return nil
	}
	if n.Left == nil {
		n = n.Right
		return n
	}
	if n.Right == nil {
		n = n.Left
		return n
	}
	leftmostrightside := n.Right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.Left != nil {
			leftmostrightside = leftmostrightside.Left
		} else {
			break
		}
	}
	n.Key, n.Data = leftmostrightside.Key, leftmostrightside.Data
	n.Right = n.Right.Remove(n.Key)
	return n
}

func (n *BinaryNode) SetData(data interface{}) error {
	// if amount <= 0 {
	// 	return errors.New("Invalid Amount")
	// }
	n.Data = data
	return nil
}

func (n *BinaryNode) SearchSubTree(key float64) *BinaryNode {
	if n == nil {
		return n
	}
	if key < n.Key {
		return n.Left.SearchSubTree(key)
	}
	if key > n.Key {
		return n.Right.SearchSubTree(key)
	}
	return n
}
