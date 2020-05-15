package engine

import "encoding/json"

// Trade describes the trade structure
type Trade struct {
	BuyOrderID  string  `json:"buy_order_id"`
	SellOrderID string  `json:"sell_order_id"`
	Amount      float64 `json:"amount"`
	Price       float64 `json:"price"`
}

// FromJSON create the Trade struct from json string
func (trade *Trade) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, trade)
}

// ToJSON returns json string of the Trade
func (trade *Trade) ToJSON() []byte {
	str, _ := json.Marshal(trade)
	return str
}
