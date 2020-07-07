package main

import (
	"fmt"
	engineGrpc "matching-engine/engineGrpc"
	"matching-engine/server"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()
	cs := server.NewEngine()
	engineGrpc.RegisterEngineServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9093")
	if err != nil {
		e := fmt.Errorf("Unable to listen server, err: %v", err)
		fmt.Println("err", e)
		os.Exit(1)
	}
	fmt.Println("grpc server listening to :9093")
	gs.Serve(l)
}
