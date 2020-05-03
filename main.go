package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"matching-engine/engine"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func printOrders(book *engine.OrderBook, order *engine.Order) {
	printArray := []string{}
	if len(book.BuyOrders) <= len(book.SellOrders) && order.Type == "sell" {
		for _, order := range book.SellOrders {
			printArray = append(printArray, " | "+fmt.Sprintf("%f - %f", order.Price, order.Amount))
		}

		for i, order := range book.BuyOrders {
			printArray[i] = fmt.Sprintf("\n%f - %f", order.Amount, order.Price) + printArray[i]
		}

	} else if len(book.BuyOrders) >= len(book.SellOrders) && order.Type == "buy" {
		for _, order := range book.BuyOrders {
			printArray = append(printArray, fmt.Sprintf("\n%f - %f", order.Amount, order.Price))
		}

		for i, order := range book.SellOrders {
			printArray[i] += " | " + fmt.Sprintf("%f - %f", order.Price, order.Amount)
		}
	}

	fmt.Println("````````````````````````````````````````````")
	for _, order := range printArray {
		fmt.Printf("%s\n", order)
	}
}
func printBook(book *engine.OrderBook) {
	printArray := []string{}
	fmt.Println("orders", len(book.BuyOrders), len(book.SellOrders))
	if len(book.BuyOrders) <= len(book.SellOrders) {
		for _, order := range book.SellOrders {
			printArray = append(printArray, " | "+fmt.Sprintf("%f - %f", order.Price, order.Amount))
		}

		for i, order := range book.BuyOrders {
			printArray[i] = fmt.Sprintf("\n%f - %f", order.Amount, order.Price) + printArray[i]
		}

	} else if len(book.BuyOrders) >= len(book.SellOrders) {
		for _, order := range book.BuyOrders {
			printArray = append(printArray, fmt.Sprintf("\n%f - %f", order.Amount, order.Price))
		}

		for i, order := range book.SellOrders {
			printArray[i] += " | " + fmt.Sprintf("%f - %f", order.Price, order.Amount)
		}
	}

	fmt.Println("````````````````````````````````````````````")
	for _, order := range printArray {
		fmt.Printf("%s\n", order)
	}
}
func main() {

	// // create the consumer and listen for new order messages
	// consumer := createConsumer()

	// // create the producer of trade messages
	// producer := createProducer()

	// create the order book
	book := engine.NewOrderBook()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// create a signal channel to know when we are done
	done := make(chan bool)

	startTime := time.Now()
	// started := false
	count := 0
	// sellPrice := 7000.0
	sellAmount := 0.0
	go func() {
		for i := 0; i < 10; i++ {
			// rand.Seed(10)
			var price float64
			price = 7000.0 + math.Ceil(rand.Float64()*500) // (rand.Float64()*5)+7000
			amount := 1 + rand.Float64()*5                 // (rand.Float64()*5)+7000
			amount = math.Floor(amount*100000000) / 100000000
			sellAmount += amount
			// fmt.Println("amount buy", amount, price )

			// types := []string{"buy", "sell"}//,"buy", "sell","buy", "sell","buy", "sell"}
			// n := rand.Intn(len(types))// % len(types)
			// orderType := types[n]
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
			// fmt.Printf("msg %#v %#v\n", order, err)

			// process the order
			//trades :=
			book.Process(order)
			count += 1
			// fmt.Printf("Trades %#v\n", trades)

			// go printOrders(book, &order)
		}
		elapsed := time.Since(startTime)
		fmt.Println("Executionn started after:", elapsed, count)
		// count = 0
		// startTime = time.Now()
		// for i := 0; i < 1; i++ {
		// 	orderType := "sell"
		// 	// fmt.Println("type:", orderType)
		// 	orderString := fmt.Sprintf("{\"id\":\"qwe\", \"type\": \"%s\", \"amount\":%f, \"price\":%f }", orderType, sellAmount, sellPrice)

		// 	// fmt.Println("orderString:", orderString)
		// 	var order engine.Order

		// 	// decode the message
		// 	err := order.FromJSON([]byte(orderString))
		// 	if err != nil {
		// 		fmt.Println("JSON Parse Error =: ", err)
		// 		continue
		// 	}
		// 	if order.Amount == 0 || order.Price == 0 {
		// 		fmt.Println("Invalid JSON")
		// 		continue
		// 	}
		// 	// fmt.Printf("msg %#v %#v\n", order, err)

		// 	// process the order
		// 	book.Process(order)
		// 	count +=1
		// 	// fmt.Printf("Trades %#v\n", trades)
		// 	// fmt.Printf("Trades %d\n", len(trades))

		// 	// go printOrders(book, &order)
		// }
		// printBook(book)
		done <- true
	}()
	// start processing orders
	// go func() {
	// 	for {
	// 		fmt.Println("\nRunnning")
	// 		elapsed := time.Since(startTime)
	// 		fmt.Println("Complete", elapsed)
	// 		select {
	// 		case err := <-consumer.Errors():
	// 			fmt.Println("consumer.Errors()", err)
	// 		case ntf := <-consumer.Notifications():
	// 			fmt.Printf("Rebalanced: %+v\n", ntf)
	// 		case msg := <-consumer.Messages():
	// 			// msgCount++
	// 			// fmt.Printf("Receiveing message => Key: %s, Value: %s\n", string(msg.Key), string(msg.Value))
	// 			count+=1
	// 			if !started {
	// 				startTime = time.Now()
	// 				started = true
	// 			}
	// 			var order engine.Order
	// 			// decode the message
	// 			err := order.FromJSON(msg.Value)
	// 			if err != nil {
	// 				consumer.MarkOffset(msg, "")
	// 				fmt.Println("JSON Parse Error =: ", err)
	// 				continue
	// 			}
	// 			if order.Amount == 0 || order.Price == 0 {
	// 				consumer.MarkOffset(msg, "")
	// 				fmt.Println("Invalid JSON")
	// 				continue
	// 			}
	// 			fmt.Printf("msg %#v %#v\n", order, err)

	// 			// process the order
	// 			trades := book.Process(order)
	// 			fmt.Printf("Trades %#v\n", trades)

	// 			printArray := []string{}
	// 			if len(book.BuyOrders) <= len(book.SellOrders) && order.Type == "sell" {
	// 				for _, order := range book.SellOrders {
	// 					printArray = append(printArray, " | "+fmt.Sprintf("%f - %f", order.Price, order.Amount))
	// 				}

	// 				for i, order := range book.BuyOrders {
	// 					printArray[i] = fmt.Sprintf("\n%f - %f", order.Amount, order.Price) + printArray[i]
	// 				}

	// 			} else if len(book.BuyOrders) >= len(book.SellOrders) && order.Type == "buy" {
	// 				for _, order := range book.BuyOrders {
	// 					printArray = append(printArray, fmt.Sprintf("\n%f - %f", order.Amount, order.Price))
	// 				}

	// 				for i, order := range book.SellOrders {
	// 					printArray[i] += " | " + fmt.Sprintf("%f - %f", order.Price, order.Amount)
	// 				}
	// 			}

	// 			for _, order := range printArray {
	// 				fmt.Printf("%s\n", order)
	// 			}

	// 			// send trades to message queue
	// 			for _, trade := range trades {
	// 				rawTrade := trade.ToJSON()
	// 				fmt.Println("Raw Trade", string(rawTrade))
	// 				producer.Input() <- &sarama.ProducerMessage{
	// 					Topic: "trades",
	// 					Value: sarama.ByteEncoder(rawTrade),
	// 				}
	// 			}
	// 			consumer.MarkOffset(msg, "")
	// 		case <-signals:
	// 			fmt.Println("Interrupt is detected")
	// 			done <- true
	// 		}
	// 	}
	// }()

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
