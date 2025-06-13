package oauth

import (
	"errors"

	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func errorResult(errorType, state string) (utils.H, error) {
	return utils.H{
		"error": errorType,
		"state": state,
	}, errors.New(errorType)
}

func invalidRequestResult(state string) (utils.H, error) {
	return errorResult(shared.InvalidRequest, state)
}

//lint:ignore U1000 keep it for consistency
func unauthorizedClientResult(state string) (utils.H, error) {
	return errorResult(shared.UnauthorizedClient, state)
}

//lint:ignore U1000 keep it for consistency
func accessDeniedResult(state string) (utils.H, error) {
	return errorResult(shared.AccessDenied, state)
}

func serverErrorResult(state string) (utils.H, error) {
	return errorResult(shared.ServerError, state)
}

func invalidGrantResult(state string) (utils.H, error) {
	return errorResult(shared.InvalidGrant, state)
}

func invalidScopeResult(state string) (utils.H, error) {
	return errorResult(shared.InvalidScope, state)
}
