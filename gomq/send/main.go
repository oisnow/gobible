package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

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

func activeMqProducer(c chan string, queue string, conn *stomp.Conn) {
	for {
		err := conn.Send(queue, "text/plain", []byte(<-c))
		fmt.Println("send active mq..." + queue)
		if err != nil {
			fmt.Println("active mq message send erorr: " + err.Error())
		}
	}
}

func main() {
	host := "127.0.0.1"
	port := "61613"
	queue := "gomq"

	activeMq := connActiveMq(host, port)
	defer activeMq.Disconnect()
	c := make(chan string)
	go activeMqProducer(c, queue, activeMq)

	for i := 0; i <= 2; i++ {

		c <- "hello world" + strconv.Itoa(i)
	}

}
