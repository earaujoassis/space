package utils

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/earaujoassis/space/internal/logs"
)

func RecoverHandler() {
	if rec := recover(); rec != nil {
		logs.Propagatef(logs.LevelError, "%+v\n%s\n", fmt.Errorf("%v", rec), string(debug.Stack()))
		os.Exit(1)
	}
}
