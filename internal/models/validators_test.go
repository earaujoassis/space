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
	assert.True(t, HasValidScopes([]string{PublicScope, ReadScope, OpenIDScope, ProfileScope, EmailScope}))
	assert.False(t, HasValidScopes([]string{PublicScope, ReadScope, OpenIDScope, ProfileScope, EmailScope, WriteScope}))
}

func TestAnyProhibitiveScopeForPublicClient(t *testing.T) {
	assert.False(t, anyProhibitiveScopeForPublicClient(PublicScope))
	assert.False(t, anyProhibitiveScopeForPublicClient(OpenIDScope))
	assert.False(t, anyProhibitiveScopeForPublicClient("public openid"))
	assert.False(t, anyProhibitiveScopeForPublicClient("  public  openid  "))
	assert.False(t, anyProhibitiveScopeForPublicClient("  public  openid  "))
	assert.True(t, anyProhibitiveScopeForPublicClient(ReadScope))
	assert.True(t, anyProhibitiveScopeForPublicClient(WriteScope))
	assert.True(t, anyProhibitiveScopeForPublicClient(ProfileScope))
	assert.True(t, anyProhibitiveScopeForPublicClient(EmailScope))
	assert.True(t, anyProhibitiveScopeForPublicClient("read write"))
	assert.True(t, anyProhibitiveScopeForPublicClient("profile email"))
	assert.True(t, anyProhibitiveScopeForPublicClient("public read"))
	assert.True(t, anyProhibitiveScopeForPublicClient("openid profile email"))
}
