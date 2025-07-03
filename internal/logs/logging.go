package logs

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/logs/level"
	"github.com/earaujoassis/space/internal/logs/plugins"
)

const (
	LevelInfo     = level.Info
	LevelError    = level.Error
	LevelCritical = level.Critical
	LevelPanic    = level.Panic
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
		Propagatef(level.Error, "sentry.Init: %s\n", err)
	}
}

func SetupRouter(router *gin.Engine) {
	plugins.SetupSentryForRouter(router)
}

func setLogForLevel(lev level.Level) {
	switch lev {
	case level.Info:
		logger.SetPrefix("[space][INFO    ] ")
	case level.Error:
		logger.SetPrefix("[space][ERROR   ] ")
	case level.Critical:
		logger.SetPrefix("[space][CRITICAL] ")
	case level.Panic:
		logger.SetPrefix("[space][PANIC   ] ")
	}
}

func Propagate(lev level.Level, msg string) {
	setLogForLevel(lev)
	switch lev {
	case level.Error:
		plugins.CaptureMessage(fmt.Sprintf("[space][ERROR] %s", msg), lev)
		logger.Println(msg)
	case level.Critical:
		plugins.CaptureMessage(fmt.Sprintf("[space][FATAL] %s", msg), lev)
		logger.Print(msg)
	case level.Panic:
		plugins.CaptureMessage(fmt.Sprintf("[space][FATAL] %s", msg), lev)
		logger.Println(msg)
		panic(msg)
	default:
		logger.Println(msg)
	}
}

func Propagatef(lev level.Level, msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf(msg, args...)
	setLogForLevel(lev)

	switch lev {
	case level.Error:
		plugins.CaptureMessage(fmt.Sprintf("[space][ERROR] %s", formattedMsg), lev)
		logger.Print(formattedMsg)
	case level.Critical:
		plugins.CaptureMessage(fmt.Sprintf("[space][FATAL] %s", formattedMsg), lev)
		logger.Print(formattedMsg)
	case level.Panic:
		plugins.CaptureMessage(fmt.Sprintf("[space][FATAL] %s", formattedMsg), lev)
		logger.Print(formattedMsg)
		panic(formattedMsg)
	default:
		logger.Print(formattedMsg)
	}
}
