package main

import "fmt"

// 单独声明变量
var name string
var age int

// 批量声明变量
var (
	name2 string
	age2  int
)

// 批量声明常量, 如果某一行后面没有赋值 默认和上一行的值一样
const (
	n1 = 100
	n2
	n3
)

// iota Go语言的常量计数器 只能在常量表达式中使用aaa
// iota 在 const 关键字出现是被重置为0 const中每新增一行常量声明 iota 计数一次
const (
	a1 = iota
	a2
	a3
)

func main() {

	// 给声明的变量赋值
	name = "小明"
	age = 18
	name2 = "小花"
	age2 = 19

	// 简短声明变量 (自推断类型声明)
	name3 := "小王"
	age3 := 20

	fmt.Println(name3)
	fmt.Println(age3)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}
