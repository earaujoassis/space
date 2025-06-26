package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/utils"
)

func signinHandler(c *gin.Context) {
	fg := ioc.GetFeatureGate(c)
	c.HTML(http.StatusOK, "satellite", utils.H{
		"Title":             " - Sign In",
		"Satellite":         "ganymede",
		"UserCreateEnabled": fg.IsActive("user.create"),
	})
}
