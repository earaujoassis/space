package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidRequestResult(t *testing.T) {
	result := InvalidRequestResult("state")
	assert.Equal(t, "state", result.State)
	assert.Equal(t, "invalid_request", result.ErrorType)
}

func TestUnauthorizedClientResult(t *testing.T) {
	result := UnauthorizedClientResult("state")
	assert.Equal(t, "state", result.State)
	assert.Equal(t, "unauthorized_client", result.ErrorType)
}

func TestAccessDeniedResult(t *testing.T) {
	result := AccessDeniedResult("state")
	assert.Equal(t, "state", result.State)
	assert.Equal(t, "access_denied", result.ErrorType)
}

func TestServerErrorResult(t *testing.T) {
	result := ServerErrorResult("state")
	assert.Equal(t, "state", result.State)
	assert.Equal(t, "server_error", result.ErrorType)
}

func TestInvalidGrantResult(t *testing.T) {
	result := InvalidGrantResult("state")
	assert.Equal(t, "state", result.State)
	assert.Equal(t, "invalid_grant", result.ErrorType)
}

func TestInvalidScopeResult(t *testing.T) {
	result := InvalidScopeResult("state")
	assert.Equal(t, "state", result.State)
	assert.Equal(t, "invalid_scope", result.ErrorType)
}
