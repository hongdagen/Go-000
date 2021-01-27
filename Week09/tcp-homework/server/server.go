package server

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

type TCPInfo struct {
	Ip   string
	Port int
}

type Message string

func NewTcpServer(ctx context.Context, tcpInfo TCPInfo, msgChan chan Message) {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", tcpInfo.Ip, tcpInfo.Port))
	defer listen.Close()
	if err != nil {
		log.Fatalf("start tcp server err:%v", err.Error())
	}



	for {
		conn, errConn := listen.Accept()
		if errConn != nil {
			log.Printf("accept error :%v", errConn.Error())
			continue
		}
		go read(ctx, conn, msgChan)
		go write(ctx, conn, msgChan)
	}

}

func read(ctx context.Context, conn net.Conn, msgChan chan Message) {
	for {
		select {
		case <-ctx.Done():
			log.Println("read function received context done")
			conn.Close()
		default:
			reader := bufio.NewReader(conn)
			var buf [1024]byte
			n, errReader := reader.Read(buf[:])
			if errReader != nil {
				log.Printf("read function reader error:%v", errReader)
				break
			}
			log.Printf("read msg: %v",string(buf[:n]))
			msgChan <- byte2Message(buf[:n])
		}
	}
}

func write(ctx context.Context, conn net.Conn, msgChan chan Message) {
	for {
		select {
		case <-ctx.Done():
			log.Println("write function received context done")
			conn.Close()
		default:
			msg := <-msgChan
			time.Sleep(1 * time.Second)
			log.Printf("write msg:%v", msg)
			_, errSend := conn.Write(message2Byte(msg))
			if errSend != nil {
				log.Printf("write msg error:%v, msg:%v", errSend.Error(), msg)
				break
			}
		}
	}
}

func byte2Message(b []byte) Message {
	return Message(string(b))
}

func message2Byte(msg Message) []byte {
	return []byte(msg)
}
