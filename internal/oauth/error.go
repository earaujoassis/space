package oauth

import (
	"errors"

	"github.com/earaujoassis/space/internal/utils"
)

const (
	errorURI string = "%s?error=%s&state=%s"
)

func errorResult(errorType, state string) (utils.H, error) {
	return utils.H{
		"error": errorType,
		"state": state,
	}, errors.New(errorType)
}

func invalidRequestResult(state string) (utils.H, error) {
	return errorResult(InvalidRequest, state)
}

//lint:ignore U1000 keep it for consistency
func unauthorizedClientResult(state string) (utils.H, error) {
	return errorResult(UnauthorizedClient, state)
}

//lint:ignore U1000 keep it for consistency
func accessDeniedResult(state string) (utils.H, error) {
	return errorResult(AccessDenied, state)
}

func serverErrorResult(state string) (utils.H, error) {
	return errorResult(ServerError, state)
}

func invalidGrantResult(state string) (utils.H, error) {
	return errorResult(InvalidGrant, state)
}

func invalidScopeResult(state string) (utils.H, error) {
	return errorResult(InvalidScope, state)
}
