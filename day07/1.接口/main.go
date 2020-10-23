package main

import "fmt"

// 接口是一种特殊的类型
type speaker interface {
	speak()
}

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func da(x speaker) {
	x.speak()
}

func main() {
	c1 := cat{}
	c1.speak()

	d1 := dog{}
	d1.speak()
}
