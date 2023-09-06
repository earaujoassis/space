package utils

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/earaujoassis/space/internal/logs"
)

func RecoverHandler() {
	if rec := recover(); rec != nil {
		logs.Propagatef(logs.Error, "%+v\n%s\n", errors.New(fmt.Sprintf("%v", rec)), string(debug.Stack()))
		os.Exit(1)
	}
}
