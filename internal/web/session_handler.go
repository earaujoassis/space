package web

import (
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func sessionHandler(c *gin.Context) {
	session := sessions.Default(c)
	repositories := ioc.GetRepositories(c)

	applicationTokenInterface := session.Get(shared.CookieSessionKey)
	applicationToken := utils.StringValue(applicationTokenInterface)
	applicationSession := repositories.Sessions().FindByToken(applicationToken, models.ApplicationToken)
	if applicationTokenInterface != nil && applicationSession.IsSavedRecord() {
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

	grantToken := repositories.Sessions().FindByToken(code, models.GrantToken)
	if grantToken.IsNewRecord() {
		c.Redirect(http.StatusFound, "/signin")
		return
	}
	repositories.Sessions().Invalidate(&grantToken)

	if _nextPath != "" {
		if _nextPath, err := url.QueryUnescape(_nextPath); err == nil {
			nextPath = _nextPath
		}
	}

	client := repositories.Clients().FindOrCreate(models.DefaultClient)
	if client.Key == clientKey && grantType == shared.AuthorizationCode && scope == models.PublicScope {
		applicationSession = models.Session{
			User:      grantToken.User,
			Client:    client,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Scopes:    models.PublicScope,
			TokenType: models.ApplicationToken,
		}

		err := repositories.Sessions().Create(&applicationSession)
		if err == nil {
			session.Set(shared.CookieSessionKey, applicationSession.Token)
			session.Save()
			c.Redirect(http.StatusFound, nextPath)
			return
		}
	}

	c.Redirect(http.StatusFound, "/signin")
}
