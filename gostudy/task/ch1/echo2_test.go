package main

import (
	"os"
	"testing"
)

//使用range循环展示 打印运行该程序的所有参数，此处和书中有所不同，书中为go run *.go, 此处为go test *.go -v
func TestEcho2(t *testing.T) {

	var s, sep string
	for _, args := range os.Args[:] {
		s += sep + args
		sep = "@@"
	}
	t.Log(s)
}
