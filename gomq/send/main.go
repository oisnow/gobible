package main

import (
	"fmt"
	"gobible/gomq/readconfig"
	"log"
	"net"
	"os"
	"strconv"
	"time"

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

	quelen, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "quelen")
	if err != nil {
		log.Fatalf("can not get value（%s）：%s", "key_default", err)
		os.Exit(1)
	}

	fmt.Print("host,port,queue", host, port, queue)

	activeMq := connActiveMq(host, port)
	defer activeMq.Disconnect()
	c := make(chan string)
	go activeMqProducer(c, queue, activeMq)
	quelens, err := strconv.Atoi(quelen)
	for i := 0; i <= quelens; i++ {

		c <- "hello world" + strconv.Itoa(i)
	}
	time.Sleep(time.Second * 10)
}
