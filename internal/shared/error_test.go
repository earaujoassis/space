package shared

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidRequestResult(t *testing.T) {
	result, err := InvalidRequestResult("state")
	assert.Equal(t, "state", result["state"])
	assert.Equal(t, "invalid_request", result["error"])
	assert.Equal(t, "invalid_request", fmt.Sprintf("%s", err))
}

func TestUnauthorizedClientResult(t *testing.T) {
	result, err := UnauthorizedClientResult("state")
	assert.Equal(t, "state", result["state"])
	assert.Equal(t, "unauthorized_client", result["error"])
	assert.Equal(t, "unauthorized_client", fmt.Sprintf("%s", err))
}

func TestAccessDeniedResult(t *testing.T) {
	result, err := AccessDeniedResult("state")
	assert.Equal(t, "state", result["state"])
	assert.Equal(t, "access_denied", result["error"])
	assert.Equal(t, "access_denied", fmt.Sprintf("%s", err))
}

func TestServerErrorResult(t *testing.T) {
	result, err := ServerErrorResult("state")
	assert.Equal(t, "state", result["state"])
	assert.Equal(t, "server_error", result["error"])
	assert.Equal(t, "server_error", fmt.Sprintf("%s", err))
}

func TestInvalidGrantResult(t *testing.T) {
	result, err := InvalidGrantResult("state")
	assert.Equal(t, "state", result["state"])
	assert.Equal(t, "invalid_grant", result["error"])
	assert.Equal(t, "invalid_grant", fmt.Sprintf("%s", err))
}

func TestInvalidScopeResult(t *testing.T) {
	result, err := InvalidScopeResult("state")
	assert.Equal(t, "state", result["state"])
	assert.Equal(t, "invalid_scope", result["error"])
	assert.Equal(t, "invalid_scope", fmt.Sprintf("%s", err))
}
