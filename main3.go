package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"matching-engine/engine"
	"matching-engine/engine/binarytree"
)

func print(w io.Writer, node *binarytree.BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v     ->    %v\n", ch, node.Price, node.Amount)
	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}

func main() {
	book := engine.NewOrderBook()
	fmt.Println("BOOK", book)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// create a signal channel to know when we are done
	done := make(chan bool)

	startTime := time.Now()
	// started := false
	count := 0
	sellPrice := 7000.0
	sellAmount := 10.0
	go func() {
		for i := 0.0; i < 5; i++ {
			// rand.Seed(10)
			var price float64
			price = 7000.0 + math.Ceil(rand.Float64()*500) // (rand.Float64()*5)+7000
			amount := 1 + rand.Float64()*5                 // (rand.Float64()*5)+7000
			amount = math.Floor(amount*100000000) / 100000000
			sellAmount += amount

			orderType := "buy"
			// fmt.Println("type:", orderType)
			orderString := fmt.Sprintf("{\"id\":\"qwe\", \"type\": \"%s\", \"amount\":%f, \"price\":%f }", orderType, amount, price)

			// fmt.Println("orderString:", orderString)
			var order engine.Order
			// decode the message
			err := order.FromJSON([]byte(orderString))
			if err != nil {
				fmt.Println("JSON Parse Error =: ", err)
				continue
			}
			if order.Amount == 0 || order.Price == 0 {
				fmt.Println("Invalid JSON")
				continue
			}

			book.Process(order)
			count += 1
		}

		elapsed := time.Since(startTime)
		fmt.Println("Executionn started after:", elapsed, count)

		// fmt.Println("old Book", sellAmount)
		// book.Print()

		count = 0
		startTime = time.Now()
		for i := 0; i < 1; i++ {
			orderType := "sell"
			// fmt.Println("type:", orderType)
			orderString := fmt.Sprintf("{\"id\":\"qwe\", \"type\": \"%s\", \"amount\":%f, \"price\":%f }", orderType, sellAmount, sellPrice)

			// fmt.Println("orderString:", orderString)
			var order engine.Order

			// decode the message
			err := order.FromJSON([]byte(orderString))
			if err != nil {
				fmt.Println("JSON Parse Error =: ", err)
				continue
			}
			if order.Amount == 0 || order.Price == 0 {
				fmt.Println("Invalid JSON")
				continue
			}
			// fmt.Printf("msg %#v %#v\n", order, err)

			// process the order
			book.Process(order)
			count += 1
			// fmt.Printf("Trades %#v\n", trades)
			// fmt.Printf("Trades %d\n", len(trades))

			// go printOrders(book, &order)
		}
		done <- true
	}()

	// wait until we are done
	<-done
	// fmt.Println()
	fmt.Println("new Book")
	book.Print()

	// fmt.Println()

	// node := book.BuyOrders.Tree.Root.Remove(7303.0)
	// var result []string
	// book.BuyOrders.Tree.Root.PreOrderTraverse(func(i float64) {
	// 	result = append(result, fmt.Sprintf("%f", i))
	// })
	// fmt.Printf("node: %#v", node)

	// fmt.Println()
	// book.Print()
	// n := book.Search(book.Root, 7213.0)
	// print(os.Stdout, book.Root, 0, 'M')
	// print(os.Stdout, n, 0, 'M')
	elapsed := time.Since(startTime)
	fmt.Println("Complete", elapsed, count)
	fmt.Println("Complete")
}
