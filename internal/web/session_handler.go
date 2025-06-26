package web

import (
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
)

func sessionHandler(c *gin.Context) {
	session := sessions.Default(c)

	userPublicID := session.Get("user_public_id")
	if userPublicID != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	var nextPath = "/"
	var scope = c.Query("scope")
	var grantType = c.Query("grant_type")
	var code = c.Query("code")
	var clientKey = c.Query("client_id")
	var _nextPath = c.Query("_")
	//var state string = c.Query("state")

	if scope == "" || grantType == "" || code == "" || clientKey == "" {
		// Original response:
		// c.String(http.StatusBadRequest, "Missing required parameters")
		c.Redirect(http.StatusFound, "/signin")
		return
	}
	if _nextPath != "" {
		if _nextPath, err := url.QueryUnescape(_nextPath); err == nil {
			nextPath = _nextPath
		}
	}

	repositories := ioc.GetRepositories(c)
	client := repositories.Clients().FindOrCreate(models.DefaultClient)
	if client.Key == clientKey && grantType == shared.AuthorizationCode && scope == models.PublicScope {
		grantToken := repositories.Sessions().FindByToken(code, models.GrantToken)
		if grantToken.IsSavedRecord() {
			session.Set("user_public_id", grantToken.User.PublicID)
			session.Save()
			repositories.Sessions().Invalidate(&grantToken)
			c.Redirect(http.StatusFound, nextPath)
			return
		}
	}

	c.Redirect(http.StatusFound, "/signin")
}
