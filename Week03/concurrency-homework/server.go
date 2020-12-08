package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type serverImpl struct {
	svr *http.Server
}

func newServer() *serverImpl {
	s := &http.Server{
		Addr: ":8811",
	}
	return &serverImpl{svr: s}
}

func (this *serverImpl) start() error {
	fmt.Println("server start")
	return this.svr.ListenAndServe()
}

func (this *serverImpl) stop(ctx context.Context) {
	fmt.Println("开始停止server.")
	this.svr.Shutdown(ctx)
}

func main() {

	server := newServer()

	stop := make(chan struct{})
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done")
			server.stop(ctx)
			stop <- struct{}{}

		}()
		return server.start()
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
	fmt.Println("server is over")
}
