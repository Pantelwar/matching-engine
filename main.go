package main

import (
	"fmt"
	"matching-engine/engine"
)

func main() {
	ob := engine.NewOrderBook()
	fmt.Println("Orderbook", ob)
}
