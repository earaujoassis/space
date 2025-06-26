package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func satelliteHandler(c *gin.Context) {
	session := sessions.Default(c)
	repositories := ioc.GetRepositories(c)
	applicationTokenInterface := session.Get(shared.CookieSessionKey)
	if applicationTokenInterface == nil {
		c.Redirect(http.StatusFound, "/signin")
		return
	}
	applicationToken := utils.StringValue(applicationTokenInterface)
	if applicationToken != "" && !security.ValidToken(applicationToken) {
		c.Redirect(http.StatusFound, "/signout")
		return
	}
	applicationSession := repositories.Sessions().FindByToken(applicationToken, models.ApplicationToken)
	if applicationSession.IsNewRecord() {
		c.Redirect(http.StatusFound, "/signout")
		return
	}
	c.HTML(http.StatusOK, "satellite", utils.H{
		"Title":     " - Mission Control",
		"Satellite": "himalia",
		"Internal":  true,
		"Data":      utils.H{},
	})
}
