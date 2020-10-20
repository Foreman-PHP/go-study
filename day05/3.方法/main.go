package main

import "fmt"

// 方法
// func (接收者变量 接收者类型) 方法名 (返回参数) { }
// Go语言中如果标识符首字母是大写就表示对外是公有的

// 结构体
type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 在前面指定接收者 (d dog)
func (d dog) wang() {
	fmt.Printf("%s: 汪汪汪", d.name)
}

func main() {
	d1 := newDog("aaa")
	d1.wang()
}
