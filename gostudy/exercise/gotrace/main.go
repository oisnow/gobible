package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("/Users/haoyunpeng/gowork/src/github.com/gobible/gostudy/exercise/gotrace/trace.out")
	if err != nil {
		panic("ceate file err")
	}

	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic("trace error")
	}

	defer trace.Stop()

	fmt.Println("start a new trace")

	//go tool trace trace.out
}
