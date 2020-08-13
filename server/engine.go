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
	book map[string]*engine.OrderBook
}

// NewEngine returns Engine object
func NewEngine() *Engine {
	return &Engine{book: map[string]*engine.OrderBook{}}
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

	if req.GetPair() == "" {
		fmt.Println("Invalid pair")
		return nil, errors.New("Invalid pair")
	}

	var pairBook *engine.OrderBook
	if val, ok := e.book[req.GetPair()]; ok {
		pairBook = val
	} else {
		pairBook = engine.NewOrderBook()
		e.book[req.GetPair()] = pairBook
	}

	ordersProcessed, partialOrder := pairBook.Process(order)

	ordersProcessedString, err := json.Marshal(ordersProcessed)

	// if order.Type.String() == "sell" {
	fmt.Println("pair:", req.GetPair())
	fmt.Println(pairBook)
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

// Cancel implements EngineServer interface
func (e *Engine) Cancel(ctx context.Context, req *engineGrpc.Order) (*engineGrpc.Order, error) {
	order := &engine.Order{ID: req.GetID()}

	if order.ID == "" {
		fmt.Println("Invalid JSON")
		return nil, errors.New("Invalid JSON")
	}

	if req.GetPair() == "" {
		fmt.Println("Invalid pair")
		return nil, errors.New("Invalid pair")
	}

	var pairBook *engine.OrderBook
	if val, ok := e.book[req.GetPair()]; ok {
		pairBook = val
	} else {
		pairBook = engine.NewOrderBook()
		e.book[req.GetPair()] = pairBook
	}

	order = pairBook.CancelOrder(order.ID)

	fmt.Println("pair:", req.GetPair())
	fmt.Println(pairBook)

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

	if req.GetPair() == "" {
		fmt.Println("Invalid pair")
		return nil, errors.New("Invalid pair")
	}

	var pairBook *engine.OrderBook
	if val, ok := e.book[req.GetPair()]; ok {
		pairBook = val
	} else {
		pairBook = engine.NewOrderBook()
		e.book[req.GetPair()] = pairBook
	}

	ordersProcessed, partialOrder := pairBook.ProcessMarket(order)

	ordersProcessedString, err := json.Marshal(ordersProcessed)

	// if order.Type.String() == "sell" {
	fmt.Println("pair:", req.GetPair())
	fmt.Println(pairBook)
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
