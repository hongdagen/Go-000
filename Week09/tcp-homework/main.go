package main

import (
	"context"
	"fmt"
	"github/hongdagen/Go-000/Week09/server"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	cancel, _ := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	msgChan := make(chan server.Message)

	go server.NewTcpServer(cancel, server.TCPInfo{
		Ip:   "127.0.0.1",
		Port: 8765,
	}, msgChan)

	go func() {
		for {
			select {
			case <-signalChan:
				fmt.Println("received os signal")
			}
		}
	}()

	time.Sleep(1 * time.Second)
	go func() {
		conn, _ := net.Dial("tcp", "127.0.0.1:8765")
		defer conn.Close()
		for i := 0; i < 10; i++ {
			conn.Write([]byte(strconv.Itoa(time.Now().Nanosecond())))
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(1 * time.Hour)
}
