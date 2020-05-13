package engine

// OrderNode ...
type OrderNode struct {
	Orders []*Order
	Volume float64
}

// NewOrderNode returns new OrderNode struct
func NewOrderNode() *OrderNode {
	return &OrderNode{Orders: []*Order{}, Volume: 0.0}
}

// AddOrder adds order to node
func (on *OrderNode) AddOrder(order Order) {
	on.Orders = append(on.Orders, &order)
}

// UpdateVolume updates volume
func (on *OrderNode) UpdateVolume(value float64) {
	on.Volume += value
}
