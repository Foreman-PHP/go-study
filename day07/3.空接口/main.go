package main

import "fmt"

// interface 接口类型
// interface{} 空接口类型
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 20)
	m1["name"] = "小明"
	m1["age"] = 1000
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)
}
