package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}
type eater interface {
	eat(string)
}

type cat struct{}

func (c cat) move() {
	fmt.Println("猫移动~")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	c1 := cat{}
	c1.move()
}
