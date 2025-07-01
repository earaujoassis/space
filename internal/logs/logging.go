package logs

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

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

func SetupRouter(router *gin.Engine) {
	plugins.SetupSentryForRouter(router)
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
	case Error:
		plugins.CaptureMessage(fmt.Sprintf("[space][ERROR] %s", msg))
		logger.Println(msg)
	case Critical:
		plugins.CaptureMessage(fmt.Sprintf("[space][FATAL] %s", msg))
		logger.Print(msg)
	case Panic:
		plugins.CaptureMessage(fmt.Sprintf("[space][FATAL] %s", msg))
		logger.Println(msg)
		panic(msg)
	default:
		logger.Println(msg)
	}
}

func Propagatef(level Level, msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf(msg, args...)
	setLogForLevel(level)

	switch level {
	case Error:
		plugins.CaptureMessage(fmt.Sprintf("[space][ERROR] %s", formattedMsg))
		logger.Print(formattedMsg)
	case Critical:
		plugins.CaptureMessage(fmt.Sprintf("[space][FATAL] %s", formattedMsg))
		logger.Print(formattedMsg)
	case Panic:
		plugins.CaptureMessage(fmt.Sprintf("[space][FATAL] %s", formattedMsg))
		logger.Print(formattedMsg)
		panic(formattedMsg)
	default:
		logger.Print(formattedMsg)
	}
}
