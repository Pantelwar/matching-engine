package engine

// OrderNode ...
type OrderNode struct {
	Orders []*Order
	Volume float64
}

func NewOrderNode() *OrderNode {
	return &OrderNode{Orders: []*Order{}, Volume: 0.0}
}

func (on *OrderNode) AddOrder(order Order) {
	on.Orders = append(on.Orders, &order)
}

func (on *OrderNode) UpdateVolume(value float64) {
	on.Volume += value
}
