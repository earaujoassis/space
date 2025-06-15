package oauth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func revokeHandler(c *gin.Context) {
	var token = c.PostForm("token")
	var tokenTypeHint = c.PostForm("token_type_hint")
	var session models.Session

	authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)
	key, secret := shared.BasicAuthDecode(authorizationBasic)
	repositories := ioc.GetRepositories(c)
	client := repositories.Clients().Authentication(key, secret)
	if client.IsNewRecord() {
		c.Header("WWW-Authenticate", "Basic realm=\"OAuth\"")
		c.JSON(http.StatusUnauthorized, utils.H{
			"error":             shared.InvalidClient,
			"error_description": "Client authentication failed",
		})
		return
	}

	if token == "" {
		c.JSON(http.StatusBadRequest, utils.H{
			"error":             shared.InvalidClient,
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
		session = repositories.Sessions().FindByToken(token, models.AccessToken)
	case models.RefreshToken:
		session = repositories.Sessions().FindByToken(token, models.RefreshToken)
	}

	if session.IsNewRecord() {
		session = repositories.Sessions().FindByToken(token, models.AccessToken)
		if session.IsNewRecord() {
			session = repositories.Sessions().FindByToken(token, models.RefreshToken)
		}
	}

	if session.ID != 0 {
		repositories.Sessions().Invalidate(&session)
	}

	c.Status(http.StatusOK)
}
