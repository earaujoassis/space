package shared

import (
	"errors"
	"fmt"

	"github.com/earaujoassis/space/internal/utils"
)

func errorResult(errorType, state string) (utils.H, error) {
	return utils.H{
		"error":             errorType,
		"error_description": fmt.Sprintf("Could not fulfill your request: %s", errorType),
		"state":             state,
	}, errors.New(errorType)
}

func InvalidRequestResult(state string) (utils.H, error) {
	return errorResult(InvalidRequest, state)
}

func UnauthorizedClientResult(state string) (utils.H, error) {
	return errorResult(UnauthorizedClient, state)
}

func AccessDeniedResult(state string) (utils.H, error) {
	return errorResult(AccessDenied, state)
}

func ServerErrorResult(state string) (utils.H, error) {
	return errorResult(ServerError, state)
}

func InvalidGrantResult(state string) (utils.H, error) {
	return errorResult(InvalidGrant, state)
}

func InvalidScopeResult(state string) (utils.H, error) {
	return errorResult(InvalidScope, state)
}
