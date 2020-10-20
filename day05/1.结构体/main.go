package main

import "fmt"

// 结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var xiaoming person
	xiaoming.age = 10
	xiaoming.name = "小明"
	xiaoming.gender = "男"
	xiaoming.hobby = []string{"篮球", "足球"}
	fmt.Println(xiaoming)
}
