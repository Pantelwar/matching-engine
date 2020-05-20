package engine

import (
	"encoding/json"
	"fmt"
)

// Order describes the struct of the order
type Order struct {
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	ID     string  `json:"id"`
	Type   Side    `json:"type"`
}

// NewOrder returns *Order
func NewOrder(id string, orderType Side, amount, price float64) *Order {
	return &Order{ID: id, Type: orderType, Amount: amount, Price: price}
}

// FromJSON create the Order struct from json string
func (order *Order) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, order)
}

// ToJSON returns json string of the order
func (order *Order) ToJSON() []byte {
	str, _ := json.Marshal(order)
	return str
}

// String implements Stringer interface
func (order *Order) String() string {
	return fmt.Sprintf("\n\"%s\":\n\tside: %v\n\tquantity: %f\n\tprice: %f\n\n", order.ID, order.Type, order.Amount, order.Price)
}
