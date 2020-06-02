package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"matching-engine/engine"
	engineGrpc "matching-engine/engineGrpc"

	"github.com/ericlagergren/decimal"
)

// Engine ...
type Engine struct {
	book *engine.OrderBook
}

// NewEngine returns Engine object
func NewEngine() *Engine {
	return &Engine{book: engine.NewOrderBook()}
}

// Process implements EngineServer interface
func (e *Engine) Process(ctx context.Context, req *engineGrpc.Order) (*engineGrpc.OutputOrders, error) {
	fmt.Println("process", req.GetPrice(), req.GetAmount(), req.GetID(), req.GetType())
	bigZero, _ := new(decimal.Big).SetString("0.0")
	orderString := fmt.Sprintf("{\"id\":\"%s\", \"type\": \"%s\", \"amount\": \"%s\", \"price\": \"%s\" }", req.GetID(), req.GetType(), req.GetAmount(), req.GetPrice())

	var order engine.Order
	// decode the message
	// fmt.Println("Orderstring =: ", orderString)
	err := order.FromJSON([]byte(orderString))
	if err != nil {
		fmt.Println("JSON Parse Error =: ", err)
		return nil, err
	}

	if order.Amount.Cmp(bigZero) == 0 || order.Price.Cmp(bigZero) == 0 {
		fmt.Println("Invalid JSON")
		return nil, errors.New("Invalid JSON")
	}

	ordersProcessed, partialOrder := e.book.Process(order)
	// fmt.Printf("\nordersProcessed: %v, \n\npartialOrder: %v\n", ordersProcessed, partialOrder)

	ordersProcessedString, err := json.Marshal(ordersProcessed)

	fmt.Println(e.book)
	if err != nil {
		return nil, err
	}
	// fmt.Println("ordersProcessedString", string(ordersProcessedString))

	if partialOrder != nil {
		var partialOrderString []byte
		partialOrderString, err = json.Marshal(partialOrder)
		return &engineGrpc.OutputOrders{OrdersProcessed: string(ordersProcessedString), PartialOrder: string(partialOrderString)}, nil
	}
	return &engineGrpc.OutputOrders{OrdersProcessed: string(ordersProcessedString), PartialOrder: ""}, nil

}
