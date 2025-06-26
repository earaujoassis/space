package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/utils"
)

func signupHandler(c *gin.Context) {
	fg := ioc.GetFeatureGate(c)
	c.HTML(http.StatusOK, "satellite", utils.H{
		"Title":             " - Sign Up",
		"Satellite":         "io",
		"UserCreateEnabled": fg.IsActive("user.create"),
		"Data": utils.H{
			"feature.gates": utils.H{
				"user.create": fg.IsActive("user.create"),
			},
		},
	})
}
