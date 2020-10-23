package mylog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	ERROR
)

// 结构体
type Logger struct {
	Level LogLevel
}

func parseLogLevel(s string) LogLevel {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "error":
		return ERROR
	default:
		return DEBUG
	}
}

// 构造函数
func NewLog(levelStr string) Logger {
	level := parseLogLevel(levelStr)
	return Logger{
		Level: level,
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime error")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

func log(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a)
	now := time.Now()
	_, filename, line := getInfo(3)
	fmt.Printf("[%s]--[%s   %d行]-- %s\n", now.Format("2006-01-02 15:04:05"), filename, line, msg)
}

func (l Logger) Debug(format string, a ...interface{}) {
	if l.Level >= DEBUG {
		log(format, a)
	}
}

func (l Logger) Info(format string, a ...interface{}) {
	if l.Level >= INFO {
		log(format, a)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	if l.Level >= ERROR {
		log(format, a)
	}
}
