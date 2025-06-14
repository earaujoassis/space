package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimStrings(t *testing.T) {
	strs := []string{ "  with spaces  ", "  in  between  " }
	concat := strings.Join(strs, "")
	assert.Equal(t, "  with spaces    in  between  ", concat)
	trimmedStrs := TrimStrings(strs)
	concat = strings.Join(trimmedStrs, "")
	assert.Equal(t, "with spacesin  between", concat)
}
