package main

import (
	"go-study/mylogs"
)

func main() {
	log := mylogs.NewLog("error")
	for {
		log.Info("这是一个 info 日志")
		log.Error("这是一个 error 日志")
	}

}
