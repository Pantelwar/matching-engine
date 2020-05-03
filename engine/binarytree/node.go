package binarytree

import (
	"errors"
	"fmt"
	"io"
)

// type Item generic.Type

// BinaryNode ...
type BinaryNode struct {
	Left   *BinaryNode
	Right  *BinaryNode
	Key    float64
	Data   interface{}
	Price  float64
	Amount float64
	ID     string
}

// NewBinaryNode ...
func NewBinaryNode(key float64, data interface{}) *BinaryNode { //price float64, amount float64) *BinaryNode {
	return &BinaryNode{Key: key, Data: data, Left: nil, Right: nil}

	// return &BinaryNode{Price: price, Amount: amount, Left: nil, Right: nil}
}

// Insert ...
func (n *BinaryNode) Insert(key float64, data interface{}) { //price float64, amount float64) {

	if n == nil {
		return
	} else if price < n.Price {
		// fmt.Println(" price < n.Price", price, n.Price)
		if n.Left == nil {
			n.Left = NewBinaryNode(price, amount)
		} else {
			n.Left.Insert(price, amount)
		}
	} else if price > n.Price {
		// fmt.Println(" price > n.Price", price, n.Price)
		if n.Right == nil {
			n.Right = NewBinaryNode(price, amount)
		} else {
			n.Right.Insert(price, amount)
		}
	} else {
		// fmt.Println(" price = n.Price", price, n.Price)

	}
	// if n == nil {
	// 	return
	// } else if price < n.Price {
	// 	// fmt.Println(" price < n.Price", price, n.Price)
	// 	if n.Left == nil {
	// 		n.Left = NewBinaryNode(price, amount)
	// 	} else {
	// 		n.Left.Insert(price, amount)
	// 	}
	// } else if price > n.Price {
	// 	// fmt.Println(" price > n.Price", price, n.Price)
	// 	if n.Right == nil {
	// 		n.Right = NewBinaryNode(price, amount)
	// 	} else {
	// 		n.Right.Insert(price, amount)
	// 	}
	// } else {
	// 	// fmt.Println(" price = n.Price", price, n.Price)

	// }
}

func (n *BinaryNode) Print(w io.Writer, ns int, ch rune) {
	if n == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v -> %f\n", ch, n.Price, n.Amount)
	n.Left.Print(w, ns+2, 'L')
	n.Right.Print(w, ns+2, 'R')

	// print(w, n.Left, ns+2, 'L')
	// print(w, n.Right, ns+2, 'R')
}

func (n *BinaryNode) PreOrderTraverse(f func(float64)) {
	if n != nil {
		n.Left.PreOrderTraverse(f)
		n.Right.PreOrderTraverse(f)
		f(n.Price)
	}
}

func (n *BinaryNode) Remove(price float64) *BinaryNode {
	if n == nil {
		return nil
	}
	if price < n.Price {
		n.Left = n.Left.Remove(price)
		return n
	}
	if price > n.Price {
		n.Right = n.Right.Remove(price)
		return n
	}
	// price == n.Price
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
	n.Price, n.Amount = leftmostrightside.Price, leftmostrightside.Amount
	n.Right = n.Right.Remove(n.Price)
	return n
}

func (n *BinaryNode) SetAmount(amount float64) error {
	if amount <= 0 {
		return errors.New("Invalid Amount")
	}
	n.Amount = amount
	return nil
}

func (n *BinaryNode) SearchSubTree(price float64) *BinaryNode {
	if n == nil {
		return n
	}
	if price < n.Price {
		return n.Left.SearchSubTree(price)
	}
	if price > n.Price {
		return n.Right.SearchSubTree(price)
	}
	return n
}
