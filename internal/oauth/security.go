package oauth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/utils"
)

// clientAuthentication authenticates a client application, extracting the key-secret pair;
//
//	and returns a client entry/model, given the key-secret pair
func clientAuthentication(authorizationHeader string) models.Client {
	key, secret := utils.BasicAuthDecode(authorizationHeader)
	return services.ClientAuthentication(key, secret)
}

// The following Authorization method is used by OAuth clients only
func clientBasicAuthorization(c *gin.Context) {
	authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)

	if !security.ValidBase64(authorizationBasic) {
		c.JSON(http.StatusBadRequest, utils.H{
			"error": "must use valid Authorization string",
		})
		c.Abort()
		return
	}

	client := clientAuthentication(authorizationBasic)
	if client.ID == 0 {
		c.Header("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"error": AccessDenied,
		})
		c.Abort()
		return
	}
	c.Set("Client", client)
	c.Next()
}

// The following Authorization method is used by the OAuth clients, with an OAuth session token
func oAuthTokenBearerAuthorization(c *gin.Context) {
	authorizationBearer := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)

	if !security.ValidToken(authorizationBearer) {
		c.JSON(http.StatusBadRequest, utils.H{
			"error": "must use valid token string",
		})
		c.Abort()
		return
	}

	session := AccessAuthentication(authorizationBearer)
	if session.ID == 0 || !services.SessionGrantsReadAbility(session) {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"error": AccessDenied,
		})
		c.Abort()
		return
	}
	c.Set("Session", session)
	c.Next()
}
