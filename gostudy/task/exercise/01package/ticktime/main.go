package main

import (
	"fmt"
	"time"
)

// TickTime 定时器的一个demo方法
func TickTime() {
	Ticker := time.Tick(time.Second * 2)
	for t := range Ticker {
		fmt.Println(t)
	}
}

func main() {
	// TickTime()

	t := time.Now()
	fmt.Println(t.Format("2006/01/02"))
	fmt.Println(t.Format("2006/01/02 15:04:05"))

}
