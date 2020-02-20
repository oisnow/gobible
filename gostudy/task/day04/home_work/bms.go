package main

import "fmt"

//Book 定义书籍结构体
type Book struct {
	title   string
	author  string
	price   float32
	publish bool
}

//AllBooks 初始化所有书籍的切片
var AllBooks = make([]*Book, 0, 200)

//ShowMenu 定义展示菜单函数
func ShowMenu() {
	fmt.Println("****************BMS 欢迎您！*******************")
	fmt.Println("1.添加书籍信息")
	fmt.Println("2.修改书籍信息")
	fmt.Println("3.展示所有书籍")
	fmt.Println("4.退出")
	fmt.Println("***********************************************")
	fmt.Println("")
}

// NewBook 书籍构造函数
func NewBook(title, author string, price float32, publish bool) *Book {
	return &Book{
		title:   title,
		author:  author,
		price:   price,
		publish: publish,
	}
}

// UpdateBook 修改书籍书籍函数
func UpdateBook() {
	var (
		title   string
		author  string
		price   float32
		publish bool
	)
	fmt.Println("修改书籍信息……")
	fmt.Print("输入书籍名称:")
	fmt.Scanln(&title)
	for _, val := range AllBooks {
		if val.title != title {
			goto err
		}
		goto update
	}

update:
	fmt.Print("输入书籍作者:")
	fmt.Scanln(&author)
	fmt.Print("输入书籍价格:")
	fmt.Scanln(&price)
	fmt.Print("该书籍是否发布:")
	fmt.Scanln(&publish)
	for _, val := range AllBooks {
		if val.title == title {
			val.title = title
			val.author = author
			val.price = price
			val.publish = publish
		}

	}
err:
	fmt.Println("书库中没有该书籍")

	return
}

// AddBookInfo 定义添加书籍函数
func AddBookInfo() {
	var (
		title   string
		author  string
		price   float32
		publish bool
	)
	fmt.Print("输入书籍名称:")
	fmt.Scanln(&title)
	fmt.Print("输入书籍作者:")
	fmt.Scanln(&author)
	fmt.Print("输入书籍价格:")
	fmt.Scanln(&price)
	fmt.Print("该书籍是否发布:")
	fmt.Scanln(&publish)
	book := NewBook(title, author, price, publish)
	AllBooks = append(AllBooks, book)
	fmt.Println("书籍添加成功!")

}

//ShowAllBooks 展示所有书籍的函数
func ShowAllBooks() {
	if len(AllBooks) == 0 {
		fmt.Println("书库中还没有书籍!")
		return
	}
	for _, val := range AllBooks {
		fmt.Printf("书籍名称:%s，书籍作者:%s,书籍价格:%f，书籍是否上架:%t \n", val.title, val.author, val.price, val.publish)
	}
}
