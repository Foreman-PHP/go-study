package main

import "fmt"

// 结构体
type person struct {
	name string
	age  int
}

// 构造函数
// 当结构体比较大的时候尽量使用结构体指针, 减少程序的内存开销
// 结构体小的时候可以使用值拷贝
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	liuyang := newPerson("刘洋", 18)
	fmt.Println(liuyang)
}
