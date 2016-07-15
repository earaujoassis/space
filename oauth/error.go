package oauth

import (
    "errors"

    "github.com/earaujoassis/space/utils"
)

func errorResult(errorType, state string) (utils.H, error) {
    return utils.H{
            "error": InvalidRequest,
            "state": state,
    }, errors.New(InvalidRequest)
}

func invalidRequestResult(state string) (utils.H, error) {
    return errorResult(InvalidRequest, state)
}

func unauthorizedClientResult(state string) (utils.H, error) {
    return errorResult(UnauthorizedClient, state)
}

func accessDeniedResult(state string) (utils.H, error) {
    return errorResult(AccessDenied, state)
}

func serverErrorResult(state string) (utils.H, error) {
    return errorResult(ServerError, state)
}
