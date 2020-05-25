package engine

import (
	"encoding/json"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// Order describes the struct of the order
type Order struct {
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	ID     string  `json:"id"`
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
	// order.Type = Sell
	fmt.Println("o.Type", order.Type)
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
	return fmt.Sprintf("\n\"%s\":\n\tside: %v\n\tquantity: %f\n\tprice: %f\n\n", order.ID, order.Type, order.Amount, order.Price)
}
