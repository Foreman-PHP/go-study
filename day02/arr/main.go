package main

import "fmt"

func main() {
	// var 数组变量名 [元素数量]

	// 定义数组一
	var arr1 [3]int
	fmt.Println(arr1)       // [0 0 0] 数组会初始化为int类型的零值
	var arr2 = [3]int{1, 2} // 使用指定的初始值完成初始化 [1 2 0]
	fmt.Println(arr1)
	fmt.Println(arr2)

	// 定义数组二
	var arr3 = [...]int{1, 2, 3}
	fmt.Println(arr1) // [1 2 3] 使用 [...] 自推断数组的长度
	var arr4 = [...]string{"北京", "上海", "广州"}
	fmt.Println(arr3)
	fmt.Println(arr4)

	// 定义数组三
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a) // [0 1 0 5] 指定索引值来初始化数组

	// 遍历数组 方法一
	var arr5 = [...]string{"北京", "上海", "广州"}
	for index, value := range arr5 {
		fmt.Println(index, value)
	}
	// 方法二
	for i := 0; i < len(arr5); i++ {
		fmt.Println(arr5[i])
	}

	// 定义多维数组
	arr6 := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	// 遍历多维数组
	for _, v1 := range arr6 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 练习题
	arr := [...]int{1, 3, 5, 7, 8}
	num := 0
	for _, value := range arr {
		num += value
	}
	fmt.Println(num)

	for i, j := range arr {
		for l, k := range arr[i+1:] {
			if j+k == 8 {
				fmt.Printf("(%d, %d)  ", i, l+i+1)
			}
		}
	}

}
