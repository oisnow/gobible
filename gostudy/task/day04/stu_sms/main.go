package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		var i int
		ShowMenu()
		fmt.Print("请输入指令")
		fmt.Scanln(&i)
		switch i {
		case 1:
			fmt.Println("12", "1")
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("3")
		case 4:
			fmt.Println("4")
		case 5:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入无效")
		}
	}
}
