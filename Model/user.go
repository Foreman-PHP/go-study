package Model

import (
	"fmt"
	"go-study/utlis"
)

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
}

func (user *User) AddUser() error {

	sql := "insert into users(username, password, email) values(?, ?, ?)"
	inst, err := utlis.Db.Prepare(sql)
	if err != nil {
		fmt.Println("预编译出现错误", err)
	}

	_, err2 := inst.Exec("admin", "admin", "123@qq.com")
	if err2 != nil {
		fmt.Println("执行错误", err2)
		return err2
	}
	return nil
}

func (user *User) AddUser2(username string, password string, email string) error {
	sql := "insert into users(username, password, email) values(?, ?, ?)"
	_, err := utlis.Db.Exec(sql, username, password, email)
	if err != nil {
		fmt.Println("执行错误", err)
	}
	return nil
}
