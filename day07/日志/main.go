package main

import (
	"go-study/day07/日志/mylog"
	"time"
)

func main() {
	log := mylog.NewLog("debug")
	for {
		name := "小米"
		log.Debug("日志内容内容内容内容内容 name:%s, id:%d", name)
		time.Sleep(time.Second)
	}
}
