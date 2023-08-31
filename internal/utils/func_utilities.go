package utils

import (
	"strings"
)

func TrimStrings(strs []string) []string {
	trimmed := make([]string, len(strs))
	for i := range strs {
		trimmed[i] = strings.TrimSpace(strs[i])
	}

	return trimmed
}
