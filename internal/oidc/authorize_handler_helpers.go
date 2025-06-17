package oidc

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func processResponseForAuthorizeHandlerIDToken(c *gin.Context, result utils.H, err error) {
	responseMode := c.Query("response_mode")
	redirectURI := c.Query("redirect_uri")
	if err != nil {
		switch responseMode {
		case shared.FormPostReponseType:
			c.HTML(http.StatusOK, "form_post.error", utils.H{
				"Callback":         redirectURI,
				"Error":            result["error"],
				"ErrorDescription": result["error_description"],
			})
		default:
			location := fmt.Sprintf(shared.ErrorFragmentURI, redirectURI, result["error"], result["state"])
			// Previous return: c.HTML(http.StatusFound, location)
			c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
				"Title":     " - Authorization Error",
				"Internal":  true,
				"ProceedTo": location,
				"ErrorCode": result["error"],
			})
		}
	} else {
		switch responseMode {
		case shared.FormPostReponseType:
			c.HTML(http.StatusOK, "form_post.id_token.success", utils.H{
				"Callback": redirectURI,
				"IDToken":  result["id_token"],
				"State":    result["state"],
			})
		// case shared.FragmentResponseType:
		default:
			location := fmt.Sprintf("%s#id_token=%s&state=%s",
				redirectURI, result["id_token"], result["state"])
			c.Redirect(http.StatusFound, location)
		}
	}
}

func processResponseForIDTokenAccessDenied(c *gin.Context) {
	responseMode := c.Query("response_mode")
	redirectURI := c.Query("redirect_uri")
	state := c.Query("state")

	switch responseMode {
	case shared.FormPostReponseType:
		c.HTML(http.StatusOK, "form_post.error", utils.H{
			"Callback":         redirectURI,
			"Error":            shared.AccessDenied,
			"ErrorDescription": "User request to deny access",
		})
	// case shared.FragmentResponseType:
	default:
		location := fmt.Sprintf(shared.ErrorFragmentURI, redirectURI, shared.AccessDenied, state)
		c.Redirect(http.StatusFound, location)
		c.Redirect(http.StatusFound, location)
	}
}

func validateResponseModeForIDToken(c *gin.Context) error {
	responseMode := c.Query("response_mode")

	if responseMode == "" ||
		responseMode == shared.FragmentResponseType ||
		responseMode == shared.FormPostReponseType {
		return nil
	}

	c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
		"Title":     " - Authorization Error",
		"Internal":  true,
		"ProceedTo": nil,
		"ErrorCode": shared.InvalidResponseMode,
	})

	return fmt.Errorf("%s", shared.InvalidResponseMode)
}

func validateScopeForIDToken(c *gin.Context) error {
	responseMode := c.Query("response_mode")
	redirectURI := c.Query("redirect_uri")
	state := c.Query("state")
	scope := c.Query("scope")

	scopes := utils.Scopes(scope)
	if scope != "" && models.HasValidScopes(scopes) && strings.Contains(scope, models.OpenIDScope) {
		return nil
	}

	switch responseMode {
	case shared.FormPostReponseType:
		c.HTML(http.StatusOK, "form_post.error", utils.H{
			"Callback":         redirectURI,
			"Error":            shared.InvalidScope,
			"ErrorDescription": "User request to deny access",
		})
	// case shared.FragmentResponseType:
	default:
		location := fmt.Sprintf(shared.ErrorFragmentURI, redirectURI, shared.InvalidScope, state)
		c.Redirect(http.StatusFound, location)
		c.Redirect(http.StatusFound, location)
	}

	return fmt.Errorf(shared.InvalidScope)
}

func processResponseForAuthorizeHandlerCode(c *gin.Context, result utils.H, err error) {
	responseMode := c.Query("response_mode")
	redirectURI := c.Query("redirect_uri")
	if err != nil {
		switch responseMode {
		case shared.FormPostReponseType:
			c.HTML(http.StatusOK, "form_post.error", utils.H{
				"Callback":         redirectURI,
				"Error":            result["error"],
				"ErrorDescription": result["error_description"],
			})
		default:
			location := fmt.Sprintf(shared.ErrorQueryURI, redirectURI, result["error"], result["state"])
			// Previous return: c.HTML(http.StatusFound, location)
			c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
				"Title":     " - Authorization Error",
				"Internal":  true,
				"ProceedTo": location,
				"ErrorCode": result["error"],
			})
		}
	} else {
		switch responseMode {
		case shared.FormPostReponseType:
			c.HTML(http.StatusOK, "form_post.code.success", utils.H{
				"Callback": redirectURI,
				"Code":     result["code"],
				"State":    result["state"],
			})
		// case shared.QueryResponseType:
		default:
			location := fmt.Sprintf("%s?code=%s&state=%s",
				redirectURI, result["code"], result["state"])
			c.Redirect(http.StatusFound, location)
		}
	}
}

func processResponseForCodeAccessDenied(c *gin.Context) {
	responseMode := c.Query("response_mode")
	redirectURI := c.Query("redirect_uri")
	state := c.Query("state")

	switch responseMode {
	case shared.FormPostReponseType:
		c.HTML(http.StatusOK, "form_post.error", utils.H{
			"Callback":         redirectURI,
			"Error":            shared.AccessDenied,
			"ErrorDescription": "User request to deny access",
		})
	// case shared.QueryResponseType:
	default:
		location := fmt.Sprintf(shared.ErrorQueryURI, redirectURI, shared.AccessDenied, state)
		c.Redirect(http.StatusFound, location)
		c.Redirect(http.StatusFound, location)
	}
}

func validateResponseModeForCode(c *gin.Context) error {
	responseMode := c.Query("response_mode")

	if responseMode == "" ||
		responseMode == shared.QueryResponseType ||
		responseMode == shared.FormPostReponseType {
		return nil
	}

	c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
		"Title":     " - Authorization Error",
		"Internal":  true,
		"ProceedTo": nil,
		"ErrorCode": shared.InvalidResponseMode,
	})

	return fmt.Errorf("%s", shared.InvalidResponseMode)
}

func validateScopeForCode(c *gin.Context) error {
	responseMode := c.Query("response_mode")
	redirectURI := c.Query("redirect_uri")
	state := c.Query("state")
	scope := c.Query("scope")

	scopes := utils.Scopes(scope)
	if scope != "" && models.HasValidScopes(scopes) && strings.Contains(scope, models.OpenIDScope) {
		return nil
	}

	switch responseMode {
	case shared.FormPostReponseType:
		c.HTML(http.StatusOK, "form_post.error", utils.H{
			"Callback":         redirectURI,
			"Error":            shared.InvalidScope,
			"ErrorDescription": "User request to deny access",
		})
	// case shared.QueryResponseType:
	default:
		location := fmt.Sprintf(shared.ErrorQueryURI, redirectURI, shared.InvalidScope, state)
		c.Redirect(http.StatusFound, location)
		c.Redirect(http.StatusFound, location)
	}

	return fmt.Errorf(shared.InvalidScope)
}
