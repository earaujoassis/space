package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

func requiresConformance(c *gin.Context) {
	host := fmt.Sprintf("%s://%s", shared.Scheme(c.Request), c.Request.Host)
	correctXRequestedBy := c.Request.Header.Get("X-Requested-By") == "SpaceApi"
	// WARNING The Origin header attribute sometimes is not sent; we should not block these requests
	sameOriginPolicy := c.Request.Header.Get("Origin") == "" || host == c.Request.Header.Get("Origin")
	if correctXRequestedBy && sameOriginPolicy {
		c.Next()
	} else {
		c.JSON(http.StatusBadRequest, utils.H{
			"error": "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy",
		})
		c.Abort()
	}
}

// The following Authorization method is used by the web client, with an action token
func actionTokenBearerAuthorization(c *gin.Context) {
	authorizationBearer := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)

	if !security.ValidToken(authorizationBearer) {
		c.JSON(http.StatusBadRequest, utils.H{
			"error": "must use valid token string",
		})
		c.Abort()
		return
	}

	action := services.ActionAuthentication(authorizationBearer)
	if action.UUID == "" || !services.ActionGrantsReadAbility(action) {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"error": shared.AccessDenied,
		})
		c.Abort()
		return
	}
	c.Set("Action", action)
	c.Next()
}
