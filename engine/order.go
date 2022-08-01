package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/Pantelwar/matching-engine/util"
)

// Order describes the struct of the order
type Order struct {
	Amount *util.StandardBigDecimal `json:"amount"` // validate:"gt=0"`
	Price  *util.StandardBigDecimal `json:"price"`  // validate:"gt=0"`
	ID     string                   `json:"id"`     // validate:"required"`
	Type   Side                     `json:"type"`   //  validate:"side_validate"`
}

// func sideValidation(fl validator.FieldLevel) bool {
// 	if fl.Field().Interface() != Buy && fl.Field().Interface() != Sell {
// 		return false
// 	}
// 	return true
// }

// NewOrder returns *Order
func NewOrder(id string, orderType Side, amount, price *util.StandardBigDecimal) *Order {
	o := &Order{ID: id, Type: orderType, Amount: amount, Price: price}
	return o
}

// FromJSON create the Order struct from json string
func (order *Order) FromJSON(msg []byte) error {
	err := json.Unmarshal(msg, order)
	if err != nil {
		return err
	}
	return nil
}

// ToJSON returns json string of the order
func (order *Order) ToJSON() ([]byte, error) {
	str, err := json.Marshal(order)
	return str, err
}

// String implements Stringer interface
func (order *Order) String() string {
	amount := order.Amount.Float64()
	price := order.Price.Float64()

	return fmt.Sprintf("\"%s\":\n\tside: %v\n\tquantity: %s\n\tprice: %s\n", order.ID, order.Type, strconv.FormatFloat(amount, 'f', -1, 64), strconv.FormatFloat(price, 'f', -1, 64))
}

// UnmarshalJSON implements json.Unmarshaler interface
func (order *Order) UnmarshalJSON(data []byte) error {
	obj := struct {
		Type   Side   `json:"type"`   // validate:"side_validate"`
		ID     string `json:"id"`     // validate:"required"`
		Amount string `json:"amount"` // validate:"required"`
		Price  string `json:"price"`  // validate:"required"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		fmt.Println("Damn errr", err)
		return err
	}

	if obj.ID == "" {
		return errors.New("ID is not present")
	}
	if obj.Type == "" {
		return errors.New("invalid order type")
	}

	var err error
	order.Price, err = util.NewDecimalFromString(obj.Price) //.Quantize(8)
	if err != nil {
		fmt.Println("price", order.Price, err.Error())
		return errors.New("invalid order price")
	}
	order.Amount, err = util.NewDecimalFromString(obj.Amount) //.Quantize(8)
	if err != nil {
		return errors.New("invalid order amount")
	}

	order.Type = obj.Type
	order.ID = obj.ID

	price := order.Price.Float64()
	if price <= 0 {
		return errors.New("Order price should be greater than zero")
	}
	amount := order.Amount.Float64()
	if amount <= 0 {
		return errors.New("Order amount should be greater than zero")
	}
	return nil
}

// MarshalJSON implements json.Marshaler interface
func (order *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Type   string `json:"type"`
			ID     string `json:"id"`
			Amount string `json:"amount"`
			Price  string `json:"price"`
		}{
			Type:   order.Type.String(),
			ID:     order.ID,
			Amount: order.Amount.String(),
			Price:  order.Price.String(),
		},
	)
}
