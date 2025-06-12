package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/utils"
)

func satelliteHandler(c *gin.Context) {
	session := sessions.Default(c)
	userPublicID := session.Get("user_public_id")
	if userPublicID == nil {
		c.Redirect(http.StatusFound, "/signin")
		return
	}
	c.HTML(http.StatusOK, "satellite", utils.H{
		"Title":     " - Mission Control",
		"Satellite": "himalia",
		"Internal":  true,
		"Data":      utils.H{},
	})
}
