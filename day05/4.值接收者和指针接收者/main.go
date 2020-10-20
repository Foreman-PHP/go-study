package main

import "fmt"

// 人 结构体
type person struct {
	name string
	age  int
}

// 方法 值接收者
func (p person) guoshengri() {
	p.age++
}

// 方法 指针接收者
func (p *person) guonian() {
	p.age++
}

// 构造函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("小明", 18)
	p1.guoshengri()
	fmt.Println(p1.age) // 打印出来还是 18 岁, 因为guoshengri()方法 是值接收者 只是把内存中的值拷贝了一份

	p1.guonian()
	fmt.Println(p1.age) // 19岁 因为guonian() 是指针接收者把内存地址传进去
}

// 什么时候使用指针接收者
// 1. 需要修改接收者中的值
// 2. 接收者是拷贝代价比较大的对象
// 3. 保证一致性, 如果有某个方法使用了指针接收者, 那么其他的方法也应该使用指针接收者
