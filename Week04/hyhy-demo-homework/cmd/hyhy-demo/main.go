package main

import (
	"context"
	"errors"
	"fmt"
	v1 "github.com/hongdagen/Go-000/Week_04/hyhy-demo-homework/api/hyhy-demo/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func newServerImpl(network, addr string) (*grpc.Server, error) {
	listen, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal("listen start failed")
	}
	server := grpc.NewServer()
	if err := server.Serve(listen); err != nil {
		log.Fatalf("RPC server listen failed. err: %s\n", err.Error())
		return nil, err
	}
	return server, nil
}

var (
	server *grpc.Server
	err    error
)

func main() {
	stop := make(chan struct{})
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		go func() {
			<-ctx.Done()
			fmt.Println("got ctx done")
			server.Stop()
			stop <- struct{}{}
		}()
		orderSave := InitOrderSave()
		server, err = newServerImpl("tcp", "127.0.0.1:8181")
		if err != nil {
			return err
		}
		v1.RegisterOrderSaveServiceServer(server, orderSave)
		return nil
	})

	g.Go(func() error {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got signal")
				return ctx.Err()
			case <-signalChan:
				return errors.New("quit")
			}
		}
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("%v\n", err)
	}
	<-stop
	fmt.Println("server done")

}
