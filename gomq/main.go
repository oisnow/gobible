package main

import (
	"fmt"
	"gobible/vendor/github.com/go-stomp/stomp"
	"net"
	"os"
)

func connActiveMq(host, port string) (stompConn *stomp.Conn) { // @todo 实现断开重连
	stompConn, err := stomp.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		fmt.Println("connect to active_mq server service, error: " + err.Error())
		os.Exit(1)
	}

	return stompConn
}

func main() {

}
