package engine

// OrderNode ...
type OrderNode struct {
	Orders []*Order `json:"orders"`
	Volume string   `json:"volume"`
}

// NewOrderNode returns new OrderNode struct
func NewOrderNode() *OrderNode {
	return &OrderNode{Orders: []*Order{}, Volume: "0"}
}

// addOrder adds order to node
func (on *OrderNode) addOrder(order Order) {
	on.updateVolume(order.Amount)
	on.Orders = append(on.Orders, &order)
}

// updateVolume updates volume
func (on *OrderNode) updateVolume(value string) {
	on.Volume = Add(on.Volume, value)
}

// removeOrder removes order from OrderNode array
func (on *OrderNode) removeOrder(index int) {
	amount := Sub("0", on.Orders[index].Amount)
	on.updateVolume(amount)
	on.Orders = append(on.Orders[:index], on.Orders[index+1:]...)
}

// // MarshalJSON implements json.Marshaler interface
// func (on *OrderNode) MarshalJSON() ([]byte, error) {

// }
