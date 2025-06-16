package utils

import (
	"slices"
	"strings"
)

func TrimStrings(strs []string) []string {
	trimmed := make([]string, len(strs))
	for i := range strs {
		trimmed[i] = strings.TrimSpace(strs[i])
	}

	return trimmed
}

func removeEmptyStrings(slice []string) []string {
	return slices.DeleteFunc(slice, func(s string) bool {
		return s == ""
	})
}

func Scopes(scopes string) []string {
	strs := TrimStrings(strings.Split(strings.TrimSpace(scopes), " "))
	return removeEmptyStrings(strs)
}

func URIs(uris string) []string {
	strs := TrimStrings(strings.Split(strings.TrimSpace(uris), "\n"))
	return removeEmptyStrings(strs)
}
