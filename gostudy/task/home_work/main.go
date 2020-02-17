package main

import (
	"fmt"
	"os"
)

/*
1.打印菜单
2.等待用户输入菜单选型
3.添加书籍函数
4.修改书籍函数
5.展示书籍函数
6.退出 os.Exit(0)

使用函数实现一个简单的图书管理系统
每本书有书名、作者、架构、上架信息
用户可以在控制台添加书籍、修改书籍、打印所有书籍信息
*/

func main() {

	for {
		ShowMenu()
		var input int
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Println("正在添加书籍信息")
			AddBookInfo()
		case 2:
			fmt.Println("2.修改书籍信息")
		case 3:
			fmt.Println("所有书籍展示如下:")
			ShowAllBooks()
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}

}
