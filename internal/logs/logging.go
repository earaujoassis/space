package logs

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

type Level int8

const (
	Info Level = iota
	Error
	Critical
	Panic
)

func init() {
	logger = log.New(os.Stdout, "", log.Ldate | log.Ltime | log.LUTC)
}

func setLogForLevel(level Level) {
	switch level {
	case Info:
		logger.SetPrefix("[space][INFO    ] ")
	case Error:
		logger.SetPrefix("[space][ERROR   ] ")
	case Critical:
		logger.SetPrefix("[space][CRITICAL] ")
	case Panic:
		logger.SetPrefix("[space][PANIC   ] ")
	}
}

func Propagate(level Level, msg string) {
	setLogForLevel(level)
	switch level {
	case Panic:
		logger.Println(msg)
		panic(msg)
	default:
		logger.Println(msg)
	}
}

func Propagatef(level Level, msg string, args ...interface{}) {
	setLogForLevel(level)
	switch level {
	case Panic:
		logger.Printf(msg, args...)
		panic(fmt.Sprintf(msg, args...))
	default:
		logger.Printf(msg, args...)
	}
}
