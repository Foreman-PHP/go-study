package mylogs

import (
	"fmt"
	"time"
)

type Logger struct {
}

func NewLog() Logger {
	return Logger{}
}

var now = time.Now()

func (l Logger) Debug(msg string) {

	fmt.Println(msg, now.Format("2006-01-02 15:04:05"))
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Waring(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	fmt.Println(msg)
}
