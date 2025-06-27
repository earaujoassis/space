package feature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFeatureAvailable(t *testing.T) {
	assert.True(t, IsFeatureAvailable("user.create"))
	assert.True(t, IsFeatureAvailable("user.adminify"))
	assert.False(t, IsFeatureAvailable("another"))
}
