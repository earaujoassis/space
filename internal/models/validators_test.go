package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasValidScopes(t *testing.T) {
	assert.False(t, HasValidScopes([]string{WriteScope}))
	assert.True(t, HasValidScopes([]string{PublicScope, OpenIDScope}))
	assert.True(t, HasValidScopes([]string{PublicScope, OpenIDScope, ProfileScope}))
	assert.True(t, HasValidScopes([]string{PublicScope, ReadScope, OpenIDScope, ProfileScope}))
	assert.False(t, HasValidScopes([]string{PublicScope, ReadScope, WriteScope, OpenIDScope, ProfileScope}))
}
