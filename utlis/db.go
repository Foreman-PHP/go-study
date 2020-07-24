package utlis

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/homestead?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
