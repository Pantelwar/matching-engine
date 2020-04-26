package engine

import "encoding/json"

// Order describes the struct of the order
type Order struct {
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	ID     string `json:"id"`
	Type   string `json:"type"`
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