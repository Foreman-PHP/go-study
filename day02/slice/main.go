package main

import "fmt"

/**
重要:
切片是指向了一个底层的数组
切片的长度就是它元素的个数
切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
切片是一个引用类型, 修改了底层数组的值, 切片的值也跟着改变
*/
func main() {

	// 声明切片
	//var s1 []int
	//var s2 []string
	//fmt.Println(s1, s2)
	//fmt.Println(s1 == nil) // true 还没初始化 没有在内存中开辟空间
	//fmt.Println(s2 == nil) // true
	//
	//// 初始化
	//s1 = []int{1, 2, 3}
	//s2 = []string{"学校", "校训", "xx"}
	//fmt.Println(s1, s2)
	//fmt.Println(s1 == nil) // false 已初始化
	//fmt.Println(s2 == nil)
	//
	//// 长度和容量
	//fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	//fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))
	//
	//// 2. 由数组得到切片
	//a1 := [...]int{1, 2, 3, 7, 8, 12, 31}
	//s3 := a1[0:4] // 基于数组切割, 左包含右不不包含
	//fmt.Println(s3)
	//s4 := a1[:4] // 相当于 a1[0:4]
	//s5 := a1[3:] // 相当于 a1[3:len(a1)] 到最后一位
	//s6 := a1[:]  // 相当于 a1[0:len(a1)] 到最后一位
	//fmt.Println(s4, s5, s6)
	//
	//fmt.Printf("len(s4):%d cap(s4):%d\n", len(s4), cap(s4)) // len(s4):4 cap(s4):7 切片的容量是底层数组从切片的第一个元素到最后一个元素的数量

	// 使用make函数创建切片
	s7 := make([]int, 5, 10) // 切片的长度是5个长度 容量是10
	fmt.Println(s7)
	//切片的赋值
	s7[0] = 100
	fmt.Println(s7)
	// 遍历切片
	for _, v := range s7 {
		fmt.Println(v)
	}

	// 使用 append() 函数为切片追加元素
	s9 := []string{"北京", "上海", "天津"}
	// 调用 append() 函数必须用原来的切换变量来接收返回值
	// append()追加元素, 原来的底层数组放不下的时候 Go底层就会把底层数组换一个
	s9 = append(s9, "欢迎你")
	fmt.Println(s9)

}
