package engine

import (
	"encoding/json"
	"reflect"
)

// Side of the order
type Side int

// Sell (asks) or Buy (bids)
const (
	Sell Side = iota
	Buy
)

// UnmarshalJSON implements interface for json unmarshal
func (s *Side) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"buy"`:
		*s = Buy
	case `"sell"`:
		*s = Sell
	default:
		return &json.UnsupportedValueError{
			Value: reflect.New(reflect.TypeOf(data)),
			Str:   string(data),
		}
	}

	return nil
}

// Order describes the struct of the order
type Order struct {
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	ID     string  `json:"id"`
	Type   Side    //string  `json:"type"`
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
