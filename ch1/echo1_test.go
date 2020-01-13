package main

import (
	"os"
	"testing"
)

//打印运行该程序的所有参数，此处和书中有所不同，书中为go run *.go, 此处为go test *.go -v
func TestEcho(t *testing.T) {
	var s, sep string
	// t.Log(os.Args[2])
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "@@@"
	}
	t.Log(s)
}
