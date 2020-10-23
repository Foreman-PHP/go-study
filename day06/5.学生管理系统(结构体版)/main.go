package main

import (
	"fmt"
	"os"
)

var smr studentMgr

func showMenu() {
	smr = studentMgr{allStudent: make(map[string]student, 100)}

	for {
		fmt.Println("学生管理系统(结构体版)")
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
			smr.SelStudent()
		case 2:
			smr.AddStudent()
		case 3:
			smr.DelStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("选择错误")
		}
	}
}

func main() {
	showMenu()
}
