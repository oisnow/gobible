package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"gobible/gomq/readconfig"

	"github.com/Unknwon/goconfig"
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

func acticeCustomer(queue string, conn *stomp.Conn, timeout time.Duration) {
	for {
		sub, _ := conn.Subscribe(queue, stomp.AckMode(stomp.AckAuto))
		for {
			select {
			case v := <-sub.C:

				//the 'v.body' type is []byte Convert string
				fmt.Println(string(v.Body))

			case <-time.After(time.Second * timeout):
				return
			}
		}
	}
}

func main() {
	filePath := "../config/config.ini"
	cfg := readconfig.GetConfigFile(filePath)

	host, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "host")
	if err != nil {
		log.Fatalf("can not get value（%s）：%s", "key_default", err)
		os.Exit(1)
	}

	port, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "port")
	if err != nil {
		log.Fatalf("can not get value（%s）：%s", "key_default", err)
		os.Exit(1)
	}

	queue, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "queue")
	if err != nil {
		log.Fatalf("can not get value（%s）：%s", "key_default", err)
		os.Exit(1)
	}

	timeout, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "timeout")
	if err != nil {
		log.Fatalf("can not get value（%s）：%s", "key_default", err)
		os.Exit(1)
	}

	activeMq := connActiveMq(host, port)
	defer activeMq.Disconnect()
	timeoutint, err := strconv.Atoi(timeout)
	acticeCustomer(queue, activeMq, time.Duration(timeoutint))

}
