package oauth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

func revokeHandler(c *gin.Context) {
	var token = c.PostForm("token")
	var tokenTypeHint = c.PostForm("token_type_hint")
	var session models.Session

	authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)
	client := ClientAuthentication(authorizationBasic)
	if client.ID == 0 {
		c.Header("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"error":             InvalidClient,
			"error_description": "Client authentication failed",
		})
		return
	}

	if token == "" {
		c.JSON(http.StatusBadRequest, utils.H{
			"error":             InvalidClient,
			"error_description": "Missing token parameter",
		})
		return
	}

	if !security.ValidToken(token) {
		c.Status(http.StatusOK)
		return
	}

	switch tokenTypeHint {
	case models.AccessToken:
		session = services.FindSessionByToken(token, models.AccessToken)
	case models.RefreshToken:
		session = services.FindSessionByToken(token, models.RefreshToken)
	}

	if session.ID == 0 {
		session = services.FindSessionByToken(token, models.AccessToken)
		if session.ID == 0 {
			session = services.FindSessionByToken(token, models.RefreshToken)
		}
	}

	if session.ID != 0 {
		services.InvalidateSession(session)
	}

	c.Status(http.StatusOK)
}
