package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func getOAuthAuthorizationServerDefinitions(c *gin.Context) utils.H {
	baseURL := shared.GetBaseUrl(c)

	return utils.H{
		"issuer":                                        baseURL,
		"authorization_endpoint":                        fmt.Sprintf("%s%s", baseURL, "/oauth/authorize"),
		"token_endpoint":                                fmt.Sprintf("%s%s", baseURL, "/oauth/token"),
		"scopes_supported":                              []string{"public", "read"},
		"response_types_supported":                      []string{"code"},
		"response_modes_supported":                      []string{"query"},
		"grant_types_supported":                         []string{"authorization_code"},
		"token_endpoint_auth_methods_supported":         []string{"client_secret_basic"},
		"service_documentation":                         "https://github.com/earaujoassis/space",
		"ui_locales_supported":                          []string{"en-US"},
		"revocation_endpoint":                           fmt.Sprintf("%s%s", baseURL, "/oauth/revoke"),
		"revocation_endpoint_auth_methods_supported":    []string{"client_secret_basic"},
		"introspection_endpoint":                        fmt.Sprintf("%s%s", baseURL, "/oauth/introspect"),
		"introspection_endpoint_auth_methods_supported": []string{"client_secret_basic"},
	}
}

func getOAuthAuthorizationServer(c *gin.Context) {
	c.JSON(http.StatusOK, getOAuthAuthorizationServerDefinitions(c))
}
