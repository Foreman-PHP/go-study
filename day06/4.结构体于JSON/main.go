package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name" db:"user_name" ini:"ini_name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "小明",
		Age:  10,
	}
	b, _ := json.Marshal(p1) // 转 json
	fmt.Printf("%#v\n", string(b))

	// json转
	str := `{"name":"小明", "age": 18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
}
