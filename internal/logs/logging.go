package logs

import (
	"fmt"
	"log"
	"os"

	"github.com/earaujoassis/space/internal/logs/plugins"
)

type Level int8

const (
	Info Level = iota
	Error
	Critical
	Panic
)

type Options struct {
	Environment string
	Release     string
	SentryUrl   string
}

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC)
}

func Setup(opts Options) {
	err := plugins.SetupSentry(opts.Environment, opts.Release, opts.SentryUrl)
	if err != nil {
		Propagatef(Error, "sentry.Init: %s\n", err)
	}
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
	case Error, Critical:
		plugins.CaptureException(msg)
		logger.Println(msg)
	case Panic:
		plugins.CaptureException(msg)
		logger.Println(msg)
		panic(msg)
	default:
		logger.Println(msg)
	}
}

func Propagatef(level Level, msg string, args ...interface{}) {
	setLogForLevel(level)
	switch level {
	case Error, Critical:
		plugins.CaptureException(fmt.Sprintf(msg, args...))
		logger.Printf(msg, args...)
		panic(fmt.Sprintf(msg, args...))
	case Panic:
		plugins.CaptureException(fmt.Sprintf(msg, args...))
		logger.Printf(msg, args...)
		panic(fmt.Sprintf(msg, args...))
	default:
		logger.Printf(msg, args...)
	}
}
