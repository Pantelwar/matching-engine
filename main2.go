package main

import (
	"fmt"
	"log"

	// "os"
	// "os/signal"
	"math"
	"math/rand"
	"time"

	// "reflect"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	ob "github.com/muzykantov/orderbook"
	"github.com/shopspring/decimal"
)

// func printOrders(book *engine.OrderBook, order *engine.Order) {
// 	printArray := []string{}
// 	if len(book.BuyOrders) <= len(book.SellOrders) && order.Type == "sell" {
// 			for _, order := range book.SellOrders {
// 				printArray = append(printArray, " | "+fmt.Sprintf("%f - %f", order.Price, order.Amount))
// 			}

// 			for i, order := range book.BuyOrders {
// 				printArray[i] = fmt.Sprintf("\n%f - %f", order.Amount, order.Price) + printArray[i]
// 			}

// 	} else if len(book.BuyOrders) >= len(book.SellOrders) && order.Type == "buy" {
// 			for _, order := range book.BuyOrders {
// 				printArray = append(printArray, fmt.Sprintf("\n%f - %f", order.Amount, order.Price))
// 			}

// 			for i, order := range book.SellOrders {
// 				printArray[i] += " | " + fmt.Sprintf("%f - %f", order.Price, order.Amount)
// 			}
// 	}

// 	fmt.Println("````````````````````````````````````````````")
// 	for _, order := range printArray {
// 		fmt.Printf("%s\n", order)
// 	}
// }
// func printBook(book *engine.OrderBook) {
// 	printArray := []string{}
// 	if len(book.BuyOrders) <= len(book.SellOrders) {
// 			for _, order := range book.SellOrders {
// 				printArray = append(printArray, " | "+fmt.Sprintf("%f - %f", order.Price, order.Amount))
// 			}

// 			for i, order := range book.BuyOrders {
// 				printArray[i] = fmt.Sprintf("\n%f - %f", order.Amount, order.Price) + printArray[i]
// 			}

// 	} else if len(book.BuyOrders) >= len(book.SellOrders){
// 			for _, order := range book.BuyOrders {
// 				printArray = append(printArray, fmt.Sprintf("\n%f - %f", order.Amount, order.Price))
// 			}

// 			for i, order := range book.SellOrders {
// 				printArray[i] += " | " + fmt.Sprintf("%f - %f", order.Price, order.Amount)
// 			}
// 	}

// 	fmt.Println("````````````````````````````````````````````")
// 	for _, order := range printArray {
// 		fmt.Printf("%s\n", order)
// 	}
// }
func main() {

	// // create the consumer and listen for new order messages
	// consumer := createConsumer()

	// // create the producer of trade messages
	// producer := createProducer()

	// create the order book
	book := ob.NewOrderBook()
	fmt.Println("Orderbook", book)
	// price, _ := decimal.NewFromString("10.3")
	// amount, _ := decimal.NewFromString("12.0")
	// book.ProcessLimitOrder(ob.Sell, "uinqueID", decimal.New(55.0, 0), decimal.New(100, 0))
	// fmt.Println("1",book)
	// book.ProcessLimitOrder(ob.Buy, "uinqueID1", decimal.New(7, 0), decimal.New(120, 0))
	// fmt.Println("2",book)
	// book.ProcessLimitOrder(ob.Buy, "uinqueID2", decimal.New(3, 0), decimal.New(120, 0))
	// fmt.Println("3",book)

	// signals := make(chan os.Signal, 1)
	// signal.Notify(signals, os.Interrupt)

	// create a signal channel to know when we are done
	done := make(chan bool)

	startTime := time.Now()
	// started := false
	count := 0
	sellPrice := 7000.0
	sellAmount := 10.0
	go func() {
		for i := 0.0; i < 100000; i++ {
			// 	// rand.Seed(10)
			var price float64
			price = 7000.0 + i             // math.Ceil(rand.Float64()*500) // (rand.Float64()*5)+7000
			amount := 1 + rand.Float64()*5 // (rand.Float64()*5)+7000
			amount = math.Floor(amount*100000000) / 100000000
			sellAmount += amount
			// 	fmt.Println(reflect.TypeOf(price), price)
			amountD, _ := decimal.NewFromString(fmt.Sprintf("%f", amount))
			priceD, _ := decimal.NewFromString(fmt.Sprintf("%f", price))
			// fmt.Println("amount buy", amountD, priceD)

			book.ProcessLimitOrder(ob.Buy, fmt.Sprintf("a%d", i), amountD, priceD)
			// fmt.Println("a b c d", a, b, c,d)

			count += 1
			// 	// fmt.Printf("Trades %#v\n", trades)

			// 	// go printOrders(book, &order)
		}
		// fmt.Println("buys",book)

		elapsed := time.Since(startTime)
		fmt.Println("Executionn started after:", elapsed, count)
		count = 0
		startTime = time.Now()
		for i := 500000; i < 500001; i++ {
			amountSell, _ := decimal.NewFromString(fmt.Sprintf("%f", sellAmount))
			priceSell, _ := decimal.NewFromString(fmt.Sprintf("%f", sellPrice))
			// fmt.Println("amount", amountSell, priceSell)
			// a, b, c, d :=
			book.ProcessLimitOrder(ob.Sell, fmt.Sprintf("%d", i), amountSell, priceSell)
			// fmt.Println("a:", a)
			// fmt.Println("b:", b)
			// fmt.Println("c:", c)
			// fmt.Println("d:", d)
		}
		// printBook(book)
		fmt.Println("final: ", book)

		done <- true
	}()

	// wait until we are done
	<-done
	elapsed := time.Since(startTime)
	fmt.Println("Complete", elapsed, count)
	fmt.Println("Complete")
}

//
// Create the consumer
//

func createConsumer() *cluster.Consumer { //sarama.PartitionConsumer {
	// define our configuration to the cluster
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.CommitInterval = 1 * time.Second

	// create the consumer
	consumer, err := cluster.NewConsumer([]string{"localhost:9092"}, "myconsumer", []string{"test"}, config)
	if err != nil {
		log.Fatal("Unable to connect consumer to kafka cluster")
	}

	// go handleErrors(consumer)
	// go handleNotifications(consumer)
	return consumer
}

// func handleErrors(consumer *cluster.Consumer) {
// 	for err := range consumer.Errors() {
// 		log.Printf("Error: %s\n", err.Error())
// 	}
// }

// func handleNotifications(consumer *cluster.Consumer) {
// 	for ntf := range consumer.Notifications() {
// 		log.Printf("Rebalanced: %+v\n", ntf)
// 	}
// }

//
// Create the producer
//

func createProducer() sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Fatal("Unable to connect producer to kafka server")
	}
	return producer
}
