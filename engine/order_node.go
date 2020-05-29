package engine

import "github.com/shopspring/decimal"

// OrderNode ...
type OrderNode struct {
	Orders []*Order        `json:"orders"`
	Volume decimal.Decimal `json:"volume"`
}

// NewOrderNode returns new OrderNode struct
func NewOrderNode() *OrderNode {
	return &OrderNode{Orders: []*Order{}, Volume: decimal.NewFromFloat(0.0)}
}

// addOrder adds order to node
func (on *OrderNode) addOrder(order Order) {
	on.updateVolume(order.Amount)
	on.Orders = append(on.Orders, &order)
}

// updateVolume updates volume
func (on *OrderNode) updateVolume(value decimal.Decimal) {
	on.Volume = on.Volume.Add(value)
}

// removeOrder removes order from OrderNode array
func (on *OrderNode) removeOrder(index int) {
	on.updateVolume(on.Orders[index].Amount.Neg())
	on.Orders = append(on.Orders[:index], on.Orders[index+1:]...)
}

// // MarshalJSON implements json.Marshaler interface
// func (on *OrderNode) MarshalJSON() ([]byte, error) {

// }
