package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int
	name string
}

var allStudent map[int]*student

func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func main() {
	allStudent = make(map[int]*student, 30)
	for {
		fmt.Println("学生管理系统(函数版)")
		fmt.Println(`
1. 查询学生
2. 新增学生
3. 删除学生
4. 退出
                     `)
		fmt.Print("请输入:")
		var number int
		fmt.Scanln(&number)
		fmt.Printf("你选择了%d\n", number)
		switch number {
		case 1:
			showStudent()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("选择错误")
		}
	}

}

func showStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

func addStudent() {
	var number int
	var name string
	fmt.Print("请输入学号:")
	fmt.Scanln(&number)
	fmt.Print("请输入姓名:")
	fmt.Scanln(&name)
	newStu := newStudent(number, name)
	allStudent[number] = newStu
}

func delStudent() {
	fmt.Print("请输入学号:")
	var number int
	fmt.Scanln(&number)
	delete(allStudent, number)
}
