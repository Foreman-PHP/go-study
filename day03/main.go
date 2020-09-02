package main

import "fmt"

// 函数

// 函数是一段代码的封装
// 把一段逻辑代码抽象出来封装到一个函数中, 给它起一个名字, 每次用到的时候直接用函数名调用就可以了
// 使用函数能够让代码结构更清晰, 更简洁

// 函数的定义
func sum(x, y int) (result int) {
	return x + y
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数但有返回值
func s3() int {
	return 3
}

// 命名返回值就相当在函数中声明了一个变量, 可以在函数中直接使用
func f4(x, y int) (res int) {
	res = x + y
	return // 使用命名返回值 return 后可以省略
}

// 多个返回值
func f5() (int, string) {
	return 1, "小明"
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f6(x string, y ...int) {
	fmt.Println(y) // 是切片类型
}

func main() {

	fmt.Println(f4(1, 2))
}
