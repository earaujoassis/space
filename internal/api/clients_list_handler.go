package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func clientsListHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	action := c.MustGet("Action").(models.Action)
	user := c.MustGet("User").(models.User)
	if user.ID != action.UserID || !user.Admin {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Clients are not available",
			"error":    shared.AccessDenied,
		})
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"_status":  "success",
		"_message": "Clients are available",
		"clients":  repositories.Clients().GetActive(),
	})
}
