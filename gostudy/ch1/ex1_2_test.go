package main

import (
	"os"
	"testing"
)

//“练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。”

func TestEx1_2(t *testing.T) {
	for i := 0; i < len(os.Args[:]); i++ {
		t.Logf("the key:%d--the val:%s", i, os.Args[i])
	}
}
