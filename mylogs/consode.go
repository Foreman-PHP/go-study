package mylogs

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	INFO LogLevel = iota
	ERROR
)

// 日志结构体
type Logger struct {
	Level LogLevel
}

// 构造函数
func NewLog(s string) Logger {
	level, _ := parseLogLevel(s)
	return Logger{
		Level: level,
	}
}

func parseLogLevel(s string) (LogLevel, error) {
	switch s {
	case "info":
		return INFO, nil
	case "error":
		return ERROR, nil
	default:
		return ERROR, nil
	}
}

func (l Logger) enable(level LogLevel) bool {
	return l.Level == level
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s], %s\n", time, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%v], %v\n", time, msg)
	}
}
