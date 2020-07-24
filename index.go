package main

import (
	"fmt"
	"go-study/Model"
	"html/template"
	"log"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("注册页面")
		t, _ := template.ParseFiles("register.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）

		number := r.FormValue("number")
		password := r.FormValue("password")

		user := Model.User{}
		res := user.AddUser2(number, password, "213")
		if res != nil {
			fmt.Println("插入失败")
		}
		fmt.Println("注册成功")


	}
}

func main() {
	http.HandleFunc("/", register)
	http.ListenAndServe(":80", nil)
}
