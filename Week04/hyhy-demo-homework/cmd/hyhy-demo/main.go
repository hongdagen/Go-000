package main

import (
	v1 "github.com/hongdagen/Go-000/Week_04/hyhy-demo-homework/api/hyhy-demo/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	orderSave:= InitOrderSave()
	listen, _ := net.Listen("tcp", "127.0.0.1:8181")
	server := grpc.NewServer()
	v1.RegisterOrderSaveServiceServer(server, orderSave)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("RPC server listen failed. err: %s\n", err.Error())
	}
}
