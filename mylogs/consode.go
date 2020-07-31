package mylogs

import (
	"fmt"
	"time"
)

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		funcName, _, line := GetInfo(2)
		fmt.Printf("[%s], [%s], [%s, %d] \n", time, msg, funcName, line)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%v], %v\n", time, msg)
	}
}
