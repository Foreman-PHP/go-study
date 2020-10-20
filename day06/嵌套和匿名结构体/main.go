package main

import "fmt"

// 匿名结构体
type test struct {
	string
	int
}

type address struct {
	city string
}

// 匿名嵌套结构体
type person struct {
	name string
	age  int
	address
}

// 嵌套结构体
type company struct {
	name string
	addr address
}

func main() {
	// 初始化匿名结构
	t := test{
		string: "小米",
		int:    10,
	}
	// 打印其中的一个值
	fmt.Println(t.string)

	// 初始化嵌套结构体
	p1 := person{
		name: "小明",
		age:  10,
		address: address{
			"北京",
		},
	}

	fmt.Println(p1.city) // 直接可以访问匿名结构体中的字段

}
