package engine

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
	"gopkg.in/go-playground/validator.v9"
)

// Order describes the struct of the order
type Order struct {
	Amount decimal.Decimal `json:"amount"` // validate:"gt=0"`
	Price  decimal.Decimal `json:"price"`  // validate:"gt=0"`
	ID     string          `json:"id"`     // validate:"required"`
	Type   Side            `json:"type"`   //  validate:"side_validate"`
}

func sideValidation(fl validator.FieldLevel) bool {
	if fl.Field().Interface() != Buy && fl.Field().Interface() != Sell {
		return false
	}
	return true
}

// NewOrder returns *Order
func NewOrder(id string, orderType Side, amount, price decimal.Decimal) *Order {
	o := &Order{ID: id, Type: orderType, Amount: amount, Price: price}
	// validate := validator.New()
	// validate.RegisterValidation("side_validate", sideValidation)
	// err := validate.Struct(o)
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return nil
	// }
	return o //&Order{ID: id, Type: orderType, Amount: amount, Price: price}
}

// FromJSON create the Order struct from json string
func (order *Order) FromJSON(msg []byte) error {
	err := json.Unmarshal(msg, order)
	if err != nil {
		return err
	}
	// validate := validator.New()
	// validate.RegisterValidation("side_validate", sideValidation)
	// err = validate.Struct(order)
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return err
	// }
	return nil
}

// ToJSON returns json string of the order
func (order *Order) ToJSON() ([]byte, error) {
	str, err := json.Marshal(order)
	return str, err
}

// String implements Stringer interface
func (order *Order) String() string {
	return fmt.Sprintf("\"%s\":\n\tside: %v\n\tquantity: %s\n\tprice: %s\n", order.ID, order.Type, order.Amount.String(), order.Price.String())
}

// UnmarshalJSON implements json.Unmarshaler interface
func (order *Order) UnmarshalJSON(data []byte) error {
	obj := struct {
		Type   Side    `json:"type"`
		ID     string  `json:"id"`
		Amount float64 `json:"amount"`
		Price  float64 `json:"price"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	order.Type = obj.Type
	order.ID = obj.ID
	order.Price = decimal.NewFromFloat(obj.Price)
	order.Amount = decimal.NewFromFloat(obj.Amount)
	return nil
}
