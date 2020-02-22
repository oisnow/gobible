package main

/*
使用面向对象编写一个简单的学生管理系统，包含功能学生信息的增删改查。
*/

import (
	"fmt"
	"os"
)

// InPutInfo 捕获输入的学生信息
func InPutInfo() *Student {
	var (
		id   int
		name string
		age  int
	)

	fmt.Println("输入学生的详细信息")
	fmt.Print("输入学生姓名——>")
	fmt.Scanln(&name)
	fmt.Print("输入学生学号——>")
	fmt.Scanln(&id)
	fmt.Print("输入学生年龄——>")
	fmt.Scanln(&age)
	stu := NewStu(name, age, id)
	return stu
}

func main() {
	var allstu = StuMgr{}
	// var stu = Student{}

	for {
		var i int
		ShowMenu()
		// fmt.Print("请输入指令->")
		fmt.Scanln(&i)

		switch i {

		case 1:

			fmt.Println("添加学生信息ing......")
			stu := InPutInfo()
			allstu.AddStu(stu)
		case 2:
			fmt.Println("修改学信息ing......")
			stu := InPutInfo()
			allstu.UpdateStu(stu)
		case 3:
			fmt.Println("删除学生信息ing......")
			stu := InPutInfo()
			allstu.DelStu(stu)
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
