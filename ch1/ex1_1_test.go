package main

import (
	"os"
	"testing"
)

//“练习 1.1：  修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字”
func TestEx1_1(t *testing.T) {
	t.Log(os.Args[0])
}
