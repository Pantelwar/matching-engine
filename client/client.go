package main

import (
	"context"
	"log"
	"time"

	engineGrpc "github.com/Pantelwar/matching-engine/engineGrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		log.Fatalf("connect grpc client fail. %v", err)
	}
	defer conn.Close()

	client := engineGrpc.NewEngineClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//limit order
	ord := &engineGrpc.Order{
		Type:   1,
		ID:     "10029",
		Amount: "1.20231",
		Price:  "2212.00",
		Pair:   "ethusdt",
	}

	outOrder, err := client.Process(ctx, ord)
	if err != nil {
		log.Fatalf("post order fail. %v", err)
	}

	log.Printf("order info: %v", outOrder)

	//fetch order
	bookInput := &engineGrpc.BookInput{
		Pair:  "btcusdt",
		Limit: 1000,
	}
	outBook, err := client.FetchBook(ctx, bookInput)

	log.Printf("orderBook: %v", outBook)
}
