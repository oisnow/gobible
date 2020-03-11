package main

import "fmt"

func main() {

	//定义一个channel类型
	//定义内部存放的数据类型 int
	var ch1 chan int
	var ch2 chan string

	fmt.Println("ch1", ch1)
	fmt.Println("ch2", ch2)

	//初始化channel
	ch3 := make(chan int, 1)
	defer close(ch3)
	//管道的操作：接受 发送 关闭
	// 发送和接受是一个符号 <-
	ch3 <- 10 //把10发送给ch3
	// <-ch3 //从ch3接受值，并直接丢弃
	ret := <-ch3 //使用变量ret 接受ch3的值10
	fmt.Println(ret)
}
