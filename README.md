# Matching Engine
Improved matching engine written in Go (Golang)


## Features
* Standard price-time priority
* Supports limit orders
* High performance (above 500k trades per second)
* Optimal memory usage

## Installation

```javascript
go get github.com/Pantelwar/matching-engine/engine
```

## Usage
To start using order book you need to create object:

```javascript
import (
  "fmt" 
  ob "github.com/Pantelwar/matching-engine/engine"
)

func main() {
  orderBook := ob.NewOrderBook()
  fmt.Println(orderBook)
}
```

Then you be able to use next primary functions:

```javascript
func (ob *OrderBook) Process(order Order, orderside string) []Trade {...}
```

## About primary functions
#### Process

```javascript
// ProcessLimitOrder places new order to the OrderBook
// Arguments:
//      order   -   type Order struct {
//                      Amount float64  `json:"amount"`
//                      Price  float64  `json:"price"`
//                      ID     string   `json:"id"`
//                      Type   string   `json:"type"`
//                  }
//                  amount  -   how much quantity you want to sell or buy
//                  price   -   no more expensive (or cheaper) this price
//                  id      -   unique order ID in depth
//                  type    -   what do you want to do ("buy" or "sell")
// Return:
//      trades  -   []Trade{} - array of trades
//                      type Trade struct {
//                          BuyOrderID  string  `json:"buy_order_id"`
//                          SellOrderID string  `json:"sell_order_id"`
//                          Amount      float64 `json:"amount"`
//                          Price       float64 `json:"price"`
//                      }
//                  buy_order_id    -   order ID of the buy order
//                  sell_order_id   -   order ID of the sell order
//                  amount          -   how musch quantity is consumed in the execution
//                  price           -   at what price execution takes place
                  
func (ob *OrderBook) Process(order Order, orderside string) []Trade {...}

```