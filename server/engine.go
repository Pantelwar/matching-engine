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

	ordersProcessedString, err := json.Marshal(ordersProcessed)

	fmt.Println(e.book)

	if err != nil {
		return nil, err
	}

	if partialOrder != nil {
		var partialOrderString []byte
		partialOrderString, err = json.Marshal(partialOrder)
		return &engineGrpc.OutputOrders{OrdersProcessed: string(ordersProcessedString), PartialOrder: string(partialOrderString)}, nil
	}
	return &engineGrpc.OutputOrders{OrdersProcessed: string(ordersProcessedString), PartialOrder: "null"}, nil
}

// Cancel implements EngineServer interface
func (e *Engine) Cancel(ctx context.Context, req *engineGrpc.Order) (*engineGrpc.Order, error) {
	order := &engine.Order{ID: req.GetID()}

	if order.ID == "" {
		fmt.Println("Invalid JSON")
		return nil, errors.New("Invalid JSON")
	}

	order = e.book.CancelOrder(order.ID)

	fmt.Println(e.book)

	if order == nil {
		return nil, errors.New("NoOrderPresent")
	}

	orderEngine := &engineGrpc.Order{}

	orderEngine.ID = order.ID
	orderEngine.Amount = order.Amount.String()
	orderEngine.Price = order.Price.String()
	orderEngine.Type = engineGrpc.Side(engineGrpc.Side_value[order.Type.String()])

	return orderEngine, nil
}

// ProcessMarket implements EngineServer interface
func (e *Engine) ProcessMarket(ctx context.Context, req *engineGrpc.Order) (*engineGrpc.OutputOrders, error) {
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

	if order.Amount.Cmp(bigZero) == 0 {
		fmt.Println("Invalid JSON")
		return nil, errors.New("Invalid JSON")
	}

	ordersProcessed, partialOrder := e.book.ProcessMarket(order)

	ordersProcessedString, err := json.Marshal(ordersProcessed)

	// if order.Type.String() == "sell" {
	fmt.Println(e.book)
	// }

	if err != nil {
		return nil, err
	}

	if partialOrder != nil {
		var partialOrderString []byte
		partialOrderString, err = json.Marshal(partialOrder)
		return &engineGrpc.OutputOrders{OrdersProcessed: string(ordersProcessedString), PartialOrder: string(partialOrderString)}, nil
	}
	return &engineGrpc.OutputOrders{OrdersProcessed: string(ordersProcessedString), PartialOrder: "null"}, nil
}
