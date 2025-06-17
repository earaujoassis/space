package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimStrings(t *testing.T) {
	strs := []string{"  with spaces  ", "  in  between  "}
	concat := strings.Join(strs, "")
	assert.Equal(t, "  with spaces    in  between  ", concat)
	trimmedStrs := TrimStrings(strs)
	concat = strings.Join(trimmedStrs, "")
	assert.Equal(t, "with spacesin  between", concat)
}

func TestTrimStringsWithEmptyString(t *testing.T) {
	srts := strings.Split("", " ")
	assert.Equal(t, 1, len(srts))
	trimmed := TrimStrings(srts)
	assert.Equal(t, 1, len(trimmed))
}

func TestScopes(t *testing.T) {
	scopes := "  test  one two  three   "
	processed := Scopes(scopes)
	assert.Equal(t, 4, len(processed))
	assert.Equal(t, "test", processed[0])
	assert.Equal(t, "one", processed[1])
	assert.Equal(t, "two", processed[2])
	assert.Equal(t, "three", processed[3])
	scopes = "three two one test"
	processed = Scopes(scopes)
	assert.Equal(t, 4, len(processed))
	assert.Equal(t, "three", processed[0])
	assert.Equal(t, "two", processed[1])
	assert.Equal(t, "one", processed[2])
	assert.Equal(t, "test", processed[3])
}

func TestURIs(t *testing.T) {
	uris := "http://localhost:8080  \n   http://example.com   "
	processed := URIs(uris)
	assert.Equal(t, 2, len(processed))
	assert.Equal(t, "http://localhost:8080", processed[0])
	assert.Equal(t, "http://example.com", processed[1])
	uris = "http://localhost:8080\nhttp://example.com"
	processed = URIs(uris)
	assert.Equal(t, 2, len(processed))
	assert.Equal(t, "http://localhost:8080", processed[0])
	assert.Equal(t, "http://example.com", processed[1])
}
