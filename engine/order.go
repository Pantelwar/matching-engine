package engine

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

// Order describes the struct of the order
type Order struct {
	Amount string `json:"amount" validate:"gt=0"`
	Price  string `json:"price" validate:"gt=0"`
	ID     string `json:"id" validate:"required"`
	Type   Side   `json:"type"  validate:"side_validate"`
}

func sideValidation(fl validator.FieldLevel) bool {
	// fmt.Println("sideValidation", fl.Field().Interface(), fl.Field().Interface() == Buy.String(), fl.Field().Interface() == Sell.String(), fl.Field().Interface() != Buy, fl.Field().Interface() != Sell)
	if fl.Field().Interface() != Buy && fl.Field().Interface() != Sell {
		return false
	}
	return true
}

// NewOrder returns *Order
func NewOrder(id string, orderType Side, amount, price string) *Order {
	o := &Order{ID: id, Type: orderType, Amount: amount, Price: price}
	// fmt.Println("Amount", o.Amount)
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
	return fmt.Sprintf("\"%s\":\n\tside: %v\n\tquantity: %s\n\tprice: %s\n", order.ID, order.Type, order.Amount, order.Price)
}

// UnmarshalJSON implements json.Unmarshaler interface
func (order *Order) UnmarshalJSON(data []byte) error {
	obj := struct {
		Amount float64 `json:"amount" validate:"gt=0"`
		Price  float64 `json:"price" validate:"gt=0"`
		ID     string  `json:"id" validate:"required"`
		Type   Side    `json:"type" validate:"side_validate"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	// fmt.Println("side print", obj.Type)
	validate := validator.New()
	validate.RegisterValidation("side_validate", sideValidation)
	err := validate.Struct(obj)
	if err != nil {
		fmt.Println("validation error", err)
		return err
	}

	priceString := strconv.FormatFloat(obj.Price, 'f', -1, 64)
	amountString := strconv.FormatFloat(obj.Amount, 'f', -1, 64)
	order.Amount = amountString
	order.Price = priceString
	order.ID = obj.ID
	order.Type = obj.Type

	return nil
}
