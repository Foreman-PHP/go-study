package mylogs

import (
	"fmt"
	"path"
	"runtime"
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

func GetInfo(n int) (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("getInfo falied\n")
		return
	}

	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(file)
	return
}
