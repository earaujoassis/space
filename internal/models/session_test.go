package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidSessionModel(t *testing.T) {
	var session Session = Session{
		Scopes:       PublicScope,
		TokenType:    PublicClient,
	}

	assert.False(t, IsValid("validate", session), "should return false for the invalid session")
}

func TestHasValidScopes(t *testing.T) {
	assert.False(t, HasValidScopes([]string{ WriteScope }))
	assert.True(t, HasValidScopes([]string{ PublicScope, OpenIDScope }))
	assert.True(t, HasValidScopes([]string{ PublicScope, OpenIDScope, ProfileScope }))
	assert.True(t, HasValidScopes([]string{ PublicScope, ReadScope, OpenIDScope, ProfileScope }))
	assert.False(t, HasValidScopes([]string{ PublicScope, ReadScope, WriteScope, OpenIDScope, ProfileScope }))
}
