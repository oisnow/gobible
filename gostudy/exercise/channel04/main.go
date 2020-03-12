package main

import (
	"fmt"
	"math/rand"
	"time"
)

//使用goroutine和channel编写消费者与生产者模型

//生产者产生6位随机数字
//消费者计算6位数之和

//一个生产者 20个消费者

// var chanitem = make(chan items)

type items struct {
	id  int32
	num int32
}

type res struct {
	items
	sum int32
}

//生产者
func producer(ch chan *items) {

	var id int32
	rand.Seed(time.Now().UnixNano())

	for {
		tmpitems := &items{
			id:  id,
			num: rand.Int31(),
		}
		ch <- tmpitems

	}
}

//消费者
func consumer(ch chan *items, chres chan *res) {

	tmpitems := <-ch
	tmpnum := tmpitems.num
	tmpsum := calcsum(tmpnum)
	tmpres := &res{
		items: *tmpitems,
		sum:   tmpsum,
	}
	chres <- tmpres
}

func calcsum(num int32) (sum int32) {
	// var sum int32 = 0
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return
}

//打印结果
func printres() {

}

func main() {
	var a int32 = 12345
	fmt.Println(calcsum(a))
}
