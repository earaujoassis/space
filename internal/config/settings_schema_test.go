package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/earaujoassis/space/test/utils"
)

func TestLoadSettingsSchema(t *testing.T) {
	if err := utils.EnsureProjectRoot(); err != nil {
		t.Fatalf("Failed to change to project root: %v", err)
	}

	schema := LoadSettingsSchema()
	authenticationNotification, ok := schema["notifications.system-email-notifications.authentication"].([]interface{})
	require.True(t, ok, fmt.Sprintf("%v", authenticationNotification))
	assert.Equal(t, "bool", authenticationNotification[0])
	assert.Equal(t, true, authenticationNotification[1])
	emailAddress, ok := schema["notifications.system-email-notifications.email-address"].([]interface{})
	require.True(t, ok)
	assert.Equal(t, "string", emailAddress[0])
	tokenIntrospection, ok := schema["notifications.client-application-email-notifications.token-introspection"].([]interface{})
	require.True(t, ok)
	assert.Equal(t, "bool", tokenIntrospection[0])
	assert.Equal(t, false, tokenIntrospection[1])
	userinfoIntrospection, ok := schema["notifications.client-application-email-notifications.userinfo-introspection"].([]interface{})
	require.True(t, ok)
	assert.Equal(t, "bool", userinfoIntrospection[0])
	assert.Equal(t, false, userinfoIntrospection[1])
}
