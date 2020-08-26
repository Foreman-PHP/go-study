package main

import "fmt"

func main() {
	// map类似于其他语言中的哈希表或字典, 以 key-value 的形式存储

	// 创建 map
	m := make(map[string]int)
	m["age"] = 18
	fmt.Println(m)

	// 方式二
	ages := map[string]int{"age": 18}
	fmt.Printf("%v", ages)

}
