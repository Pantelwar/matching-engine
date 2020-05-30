package engine

import "github.com/shopspring/decimal"

// OrderNode ...
type OrderNode struct {
	Orders []*Order `json:"orders"`
	Volume float64  `json:"volume"`
}

// NewOrderNode returns new OrderNode struct
func NewOrderNode() *OrderNode {
	return &OrderNode{Orders: []*Order{}, Volume: 0.0}
}

// addOrder adds order to node
func (on *OrderNode) addOrder(order Order) {
	on.updateVolume(order.Amount)
	on.Orders = append(on.Orders, &order)
	// fmt.Println("orders", on.Volume, on.Orders)
}

// updateVolume updates volume
func (on *OrderNode) updateVolume(value float64) {
	amount := decimal.NewFromFloat(on.Volume).Add(decimal.NewFromFloat(value))
	// fmt.Println("amount", amount)
	on.Volume, _ = amount.Float64()
	// fmt.Println("amount more before", on.Volume, value)
	// if value < 0 {
	// 	value = math.Floor(value*100000000) / 100000000
	// } else {
	// 	value = math.Ceil(value*100000000) / 100000000
	// }
	// on.Volume += value
	// fmt.Println("amount before", on.Volume, value)
	// on.Volume = math.Ceil(on.Volume*100000000) / 100000000

	// fmt.Println("amount", on.Volume)
	// on.Volume, _ = decimal.NewFromFloat(on.Volume).Float64()

	// on.Volume = math.Floor(on.Volume*100000000) / 100000000

	// v := new(big.Float).SetFloat64(on.Volume)
	// val := new(big.Float).SetFloat64(value)
	// fmt.Println("amount before", on.Volume, v.String(), val.String())

	// on.Volume, _ = v.Add(v, val).Float64()

	// fmt.Println("amount", on.Volume)
	// fmt.Println()
}

// removeOrder removes order from OrderNode array
func (on *OrderNode) removeOrder(index int) {
	on.updateVolume(-on.Orders[index].Amount)
	on.Orders = append(on.Orders[:index], on.Orders[index+1:]...)
}

// // MarshalJSON implements json.Marshaler interface
// func (on *OrderNode) MarshalJSON() ([]byte, error) {

// }
