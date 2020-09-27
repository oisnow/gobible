package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/Unknwon/goconfig"

	"github.com/go-stomp/stomp"
)

func getConfigFile(filePath string) (configFile *goconfig.ConfigFile) {
	cfg, err := goconfig.LoadConfigFile(filePath)
	if err != nil {
		log.Fatal("can not load config")
	}
	return cfg
}

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
	filePath := "config.ini"
	cfg := getConfigFile(filePath)

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

	activeMq := connActiveMq(host, port)
	defer activeMq.Disconnect()
	acticeCustomer(queue, activeMq)

}
