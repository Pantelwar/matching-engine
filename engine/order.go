package engine

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

// Order describes the struct of the order
type Order struct {
	Amount float64 `json:"amount" validate:"gt=0"`
	Price  float64 `json:"price" validate:"gt=0"`
	ID     string  `json:"id" validate:"required"`
	Type   Side    `json:"type"  validate:"side_validate"`
}

func sideValidation(fl validator.FieldLevel) bool {
	if fl.Field().Interface() != Buy && fl.Field().Interface() != Sell {
		return false
	}
	return true
}

// NewOrder returns *Order
func NewOrder(id string, orderType Side, amount, price float64) *Order {
	o := &Order{ID: id, Type: orderType, Amount: amount, Price: price}
	validate := validator.New()
	validate.RegisterValidation("side_validate", sideValidation)
	err := validate.Struct(o)
	if err != nil {
		fmt.Println("error", err)
		return nil
	}
	return o //&Order{ID: id, Type: orderType, Amount: amount, Price: price}
}

// FromJSON create the Order struct from json string
func (order *Order) FromJSON(msg []byte) error {
	err := json.Unmarshal(msg, order)
	if err != nil {
		return err
	}
	validate := validator.New()
	validate.RegisterValidation("side_validate", sideValidation)
	err = validate.Struct(order)
	if err != nil {
		fmt.Println("error", err)
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
	return fmt.Sprintf("\"%s\":\n\tside: %v\n\tquantity: %s\n\tprice: %s\n", order.ID, order.Type, strconv.FormatFloat(order.Amount, 'f', -1, 64), strconv.FormatFloat(order.Price, 'f', -1, 64))
}
