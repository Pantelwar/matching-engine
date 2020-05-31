package engine

import (
	"github.com/ericlagergren/decimal"
)

// OrderNode ...
type OrderNode struct {
	Orders []*Order     `json:"orders"`
	Volume *decimal.Big `json:"volume"`
}

// NewOrderNode returns new OrderNode struct
func NewOrderNode() *OrderNode {
	vol, _ := new(decimal.Big).SetString("0.0")
	return &OrderNode{Orders: []*Order{}, Volume: vol}
}

// addOrder adds order to node
func (on *OrderNode) addOrder(order Order) {
	on.updateVolume(order.Amount)
	on.Orders = append(on.Orders, &order)
}

// updateVolume updates volume
func (on *OrderNode) updateVolume(value *decimal.Big) {
	on.Volume = new(decimal.Big).Add(on.Volume, value)
	// fmt.Println("onVolume", on.Volume)
}

// removeOrder removes order from OrderNode array
func (on *OrderNode) removeOrder(index int) {
	on.updateVolume(new(decimal.Big).Neg(on.Orders[index].Amount))
	on.Orders = append(on.Orders[:index], on.Orders[index+1:]...)
}

// // MarshalJSON implements json.Marshaler interface
// func (on *OrderNode) MarshalJSON() ([]byte, error) {

// }
