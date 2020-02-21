package main

import (
	"fmt"
)

//Student 定义学生struct
type Student struct {
	Name string
	Age  int
	ID   int
}

// StuMgr 结构体 定义所有学生列表 使用slice格式
type StuMgr struct {
	AllStudent []*Student
}

// var AllStudent = make([]*Student, 0, 200)

//NewStu 初始化stu信息
func (s *Student) NewStu(name string, age, id int) *Student {
	return &Student{
		Name: name,
		Age:  age,
		ID:   id,
	}

}

// ShowStu Student的 添加学生方法5
func (s *StuMgr) ShowStu() {
	// fmt.Println(len(s.AllStudent))
	if len(s.AllStudent) == 0 {
		println("SMS中暂时没有数据")
	}
	for _, stu := range s.AllStudent {
		// fmt.Println(stu)
		fmt.Printf("ID:%d,姓名:%s,年龄:%d \n", stu.ID, stu.Name, stu.Age)
	}
}

// AddStu Student的 添加学生方法
func (s *StuMgr) AddStu(stu *Student) {
	for _, st := range s.AllStudent {
		if st.Name == stu.Name {
			fmt.Println("该学生信息已经存在")
			return
		}
	}
	s.AllStudent = append(s.AllStudent, stu)
	fmt.Printf("[ ID:%d,姓名:%s,年龄:%d ] 已添加到SMS系统 \n", stu.ID, stu.Name, stu.Age)

}

// DelStu Student的 添加学生方法
func (s *StuMgr) DelStu(stu *Student) {
	for key, st := range s.AllStudent {
		if st.ID == stu.ID && st.Name == stu.Name {
			s.AllStudent = append(s.AllStudent[:key], s.AllStudent[key+1:]...)
			fmt.Printf("[ ID:%d,姓名:%s,年龄:%d ] 已完成SMS系统中的删除 \n", stu.ID, stu.Name, stu.Age)
			return
		}

	}
	fmt.Printf("[ ID:%d,姓名:%s ] 该学生信息不存在 \n", stu.ID, stu.Name)

}

// UpdateStu Student的 修改学生方法
func (s *StuMgr) UpdateStu(stu *Student) {
	for _, st := range s.AllStudent {
		if st.Name == stu.Name {
			st.Name = stu.Name
			st.Age = stu.Age
			st.ID = stu.ID
			fmt.Printf("[ ID:%d,姓名:%s,年龄:%d ] 已完成SMS系统中的更新 \n", stu.ID, stu.Name, stu.Age)
			return
		}

	}
	fmt.Printf("[ ID:%d,姓名:%s,年龄:%d ] 该学生信息不存在 \n", stu.ID, stu.Name, stu.Age)

}

// ShowMenu 菜单展示函数
func ShowMenu() {
	fmt.Println("===================================================================")
	fmt.Println("welcome to SMS(Student Manager System)!")
	fmt.Println("1.添加学生")
	fmt.Println("2.修改学生")
	fmt.Println("3.删除学生")
	fmt.Println("4.展示学生")
	fmt.Println("5.退出")
	fmt.Println("===================================================================")
}
