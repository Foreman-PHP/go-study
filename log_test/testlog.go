package main

import "go-study/mylogs"

func main() {

	log := mylogs.NewLog("info")
	for {
		log.Info("这是一个 info 日志")
		log.Error("这是一个 error 日志")
	}

}
