package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-stomp/stomp"
)

func connActiveMq(host, port string) (stompConn *stomp.Conn) {
	stompConn, err := stomp.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		fmt.Println("connect to active_mq server service, error: " + err.Error())
		os.Exit(1)
	}

	return stompConn
}

func acticeCustomer(queue string, conn *stomp.Conn) {
	for {
		sub, _ := conn.Subscribe(queue, stomp.AckMode(stomp.AckAuto))
		for {
			select {
			case v := <-sub.C:

				//the 'v.body' type is []byte Convert string
				fmt.Println(string(v.Body))

			case <-time.After(time.Second * 100):
				return
			}
		}
	}
}

func main() {
	host := "127.0.0.1"
	port := "61613"
	queue := "gomq"

	activeMq := connActiveMq(host, port)
	defer activeMq.Disconnect()
	acticeCustomer(queue, activeMq)

}
