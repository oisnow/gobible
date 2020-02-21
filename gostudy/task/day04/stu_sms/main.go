package main

import (
	"fmt"
	"os"
)

func main() {
	var allstu = StuMgr{}
	var stu = Student{}

	for {
		var i int
		ShowMenu()
		fmt.Print("请输入指令->")
		fmt.Scanln(&i)
		var (
			id   int
			name string
			age  int
		)
		switch i {

		case 1:

			fmt.Println("添加学生信息ing......")
			fmt.Println("输入学生ID")
			fmt.Scanln(&id)
			fmt.Println("输入学生姓名")
			fmt.Scanln(&name)
			fmt.Println("输入学生年龄")
			fmt.Scanln(&age)
			allstu.AddStu(stu.NewStu(name, age, id))
		case 2:
			fmt.Println("修改学信息ing......")
			fmt.Println("输入要修改学生姓名")
			fmt.Scanln(&name)
			fmt.Println("输入要修改学生ID")
			fmt.Scanln(&id)
			fmt.Println("输入要修改学生年龄")
			fmt.Scanln(&age)
			allstu.UpdateStu(stu.NewStu(name, age, id))
		case 3:
			fmt.Println("删除学生信息ing......")
			fmt.Println("输入学生ID")
			fmt.Scanln(&id)
			fmt.Println("输入学生姓名确认")
			fmt.Scanln(&name)
			allstu.DelStu(stu.NewStu(name, age, id))
		case 4:
			fmt.Println("展示学生信息ing......")
			allstu.ShowStu()
		case 5:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入无效")
		}
	}
}
