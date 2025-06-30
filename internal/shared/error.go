package shared

import (
	"fmt"
)

type RequestError struct {
	ErrorType        string
	ErrorDescription string
	State            string
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("request error: %s", e.ErrorDescription)
}

func errorResult(errorType, state string) *RequestError {
	return &RequestError{
		ErrorType:        errorType,
		ErrorDescription: fmt.Sprintf("could not fulfill your request: %s", errorType),
		State:            state,
	}
}

func InvalidRequestResult(state string) *RequestError {
	return errorResult(InvalidRequest, state)
}

func UnauthorizedClientResult(state string) *RequestError {
	return errorResult(UnauthorizedClient, state)
}

func AccessDeniedResult(state string) *RequestError {
	return errorResult(AccessDenied, state)
}

func ServerErrorResult(state string) *RequestError {
	return errorResult(ServerError, state)
}

func InvalidGrantResult(state string) *RequestError {
	return errorResult(InvalidGrant, state)
}

func InvalidScopeResult(state string) *RequestError {
	return errorResult(InvalidScope, state)
}
