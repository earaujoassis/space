package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQueryString(t *testing.T) {
	result := ParseQueryString("https://localhost?key1=value1&key2=value2")
	assert.Equal(t, "value1", result["key1"])
	assert.Equal(t, "value2", result["key2"])
	result = ParseQueryString("")
	assert.Zero(t, len(result))
	result = ParseQueryString("https://localhost?")
	assert.Zero(t, len(result))
}

func TestParseFragmentString(t *testing.T) {
	result := ParseFragmentString("https://localhost#key1=value1&key2=value2")
	assert.Equal(t, "value1", result["key1"])
	assert.Equal(t, "value2", result["key2"])
	result = ParseFragmentString("")
	assert.Zero(t, len(result))
	result = ParseFragmentString("https://localhost#")
	assert.Zero(t, len(result))
}
