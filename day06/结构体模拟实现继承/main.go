package main

import "fmt"

// 人结构体有 move 方法
type person struct {
	name string
}

func (p person) move() {
	fmt.Printf("%s会走路\n", p.name)
}

// dinner结构体有吃饭方法
type dinner struct {
	cf string
	person
}

func (d dinner) chifan() {
	fmt.Printf("%s在吃饭\n", d.name)
}

func main() {
	x := dinner{
		cf:     "xxxx",
		person: person{"小明"},
	}
	x.chifan()
	x.move()
}
