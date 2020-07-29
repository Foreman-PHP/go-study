package main

import "go-study/mylogs"

func main() {
	log := mylogs.NewLog()

	for {
		log.Debug("这是一个 debug 日志")
		log.Info("这是一个 info 日志")
		log.Error("这是一个 Error 日志")
	}
}
