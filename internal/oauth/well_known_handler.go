package oauth

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/utils"
)

func getOAuthAuthorizationServerDefinitions(c *gin.Context) utils.H {
	baseURL := getBaseUrl(c)

	return utils.H{
		"issuer": baseURL,
		"authorization_endpoint": fmt.Sprintf("%s%s", baseURL, "/oauth/authorize"),
		"token_endpoint": fmt.Sprintf("%s%s", baseURL, "/oauth/token"),
		"response_types_supported": []string{ "code" },
		"grant_types_supported": []string{ "authorization_code" },
		"scopes_supported": []string{ "openid", "public", "read", "write" },
		"token_endpoint_auth_methods_supported": []string{ "client_secret_basic" },
		"response_modes_supported": []string{ "query" },
		"revocation_endpoint": fmt.Sprintf("%s%s", baseURL, "/oauth/revoke"),
		"introspection_endpoint": fmt.Sprintf("%s%s", baseURL, "/oauth/introspect"),
		"service_documentation": "https://github.com/earaujoassis/space",
		"ui_locales_supported": []string{ "en-US" },
	}
}

func getOAuthAuthorizationServer(c *gin.Context) {
	c.JSON(http.StatusOK, getOAuthAuthorizationServerDefinitions(c))
}
