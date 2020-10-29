package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用goroutine和channel编写消费者与生产者模型

//生产者产生6位随机数字
//消费者计算6位数之和

//一个生产者 20个消费者

var chanitem chan *items
var chanres chan *res
var wg sync.WaitGroup

// var chanitem = make(chan items)

type items struct {
	id  int32
	num int32
}

type res struct {
	items *items
	sum   int32
}

//生产者
func producer(ch chan *items) {

	var id int32
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		id++
		tmpitems := &items{
			id:  id,
			num: rand.Int31(),
		}
		ch <- tmpitems

	}
	close(ch)
}

//消费者
func consumer(ch chan *items, chres chan *res) {
	defer wg.Done()
	for tmpitems := range ch {
		tmpsum := calcsum(tmpitems.num)
		tmpres := &res{
			items: tmpitems,
			sum:   tmpsum,
		}
		chres <- tmpres
		time.Sleep(time.Millisecond * 100)
	}
}

func calcsum(num int32) (sum int32) {
	// var sum int32 = 0
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return
}

func strartworker(n int, ch chan *items, chres chan *res) {
	for i := 0; i < n; i++ {
		go consumer(ch, chres)
		// fmt.Println("------------------------")
	}
}

//打印结果
func printres(res chan *res) {
	for {
		ret, ok := <-res
		if !ok {
			break
		} else {
			fmt.Printf("id:%d,num:%d,sum:%d \n", ret.items.id, ret.items.num, ret.sum)
		}

	}

}

func main() {

	chanitem = make(chan *items, 10000)
	chanres = make(chan *res, 10000)
	go producer(chanitem)
	wg.Add(5)
	strartworker(5, chanitem, chanres)
	wg.Wait()
	close(chanres)
	printres(chanres)
}
