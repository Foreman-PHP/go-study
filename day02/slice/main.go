package main

import "fmt"

func main() {

	// 声明切片
	var s1 []int
	fmt.Println(s1)

	// 从数组中创建
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 9}
	s2 := a[5]
	fmt.Println(s2)
}
