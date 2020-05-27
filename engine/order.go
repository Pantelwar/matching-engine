package engine

import (
	"encoding/json"
	"fmt"
	"math/big"

	"gopkg.in/go-playground/validator.v9"
)

// Order describes the struct of the order
type Order struct {
	Amount *big.Float `json:"amount"` // validate:"side_validate"`
	Price  *big.Float `json:"price"`  // validate:"side_validate"`
	ID     string     `json:"id"`     // validate:"required"`
	Type   Side       `json:"type"`   //  validate:"side_validate"`
}

func sideValidation(fl validator.FieldLevel) bool {
	// fmt.Println("sideValidation", fl.Field().Interface())
	if fl.Field().Interface() != Buy && fl.Field().Interface() != Sell {
		return false
	}
	return true
}

func bigFloatValidation(fl validator.FieldLevel) bool {
	// fmt.Println("bigFloatValidation", fl.Field().Float(), fl.Field().Float() <= 0)
	if fl.Field().Float() <= 0 {
		return false
	}
	return true
}

// NewOrder returns *Order
func NewOrder(id string, orderType Side, amount, price string) *Order {
	amountBig, _ := new(big.Float).SetString(amount)
	priceBig, _ := new(big.Float).SetString(price)
	o := &Order{ID: id, Type: orderType, Amount: amountBig, Price: priceBig}
	validate := validator.New()
	validate.RegisterValidation("side_validate", sideValidation)
	// validate.RegisterValidation("big_float_validation", bigFloatValidation)
	err := validate.Struct(o)
	if err != nil {
		// fmt.Println("error", err)
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
	// fmt.Printf("Final order: %#v\n", order)
	// validate := validator.New()
	// validate.RegisterValidation("side_validate", sideValidation)
	// // validate.RegisterValidation("big_float_validation", bigFloatValidation)
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
		Amount float64 `json:"amount" validate:"big_float_validation"`
		Price  float64 `json:"price" validate:"big_float_validation"`
		ID     string  `json:"id" validate:"required"`
		Type   Side    `json:"type" validate:"side_validate"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	validate := validator.New()
	validate.RegisterValidation("side_validate", sideValidation)
	validate.RegisterValidation("big_float_validation", bigFloatValidation)
	err := validate.Struct(obj)
	if err != nil {
		// fmt.Println("error", err)
		return err
	}

	order.ID = obj.ID
	order.Amount = new(big.Float).SetFloat64(obj.Amount)
	order.Price = new(big.Float).SetFloat64(obj.Price)
	order.Type = obj.Type
	// fmt.Printf("order Value: %#v\n", order)
	return nil
}

// MarshalJSON implements json.Marshaler interface
func (order *Order) MarshalJSON() ([]byte, error) {
	orderAmount, _ := order.Amount.Float64()
	orderPrice, _ := order.Price.Float64()
	return json.Marshal(
		&struct {
			Amount float64 `json:"amount"`
			Price  float64 `json:"price"`
			ID     string  `json:"id"`
			Type   Side    `json:"type"`
		}{
			Amount: orderAmount,
			Price:  orderPrice,
			ID:     order.ID,
			Type:   order.Type,
		},
	)
}
