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

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		_ = fmt.Errorf("Unable to listen server, err: %v", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
