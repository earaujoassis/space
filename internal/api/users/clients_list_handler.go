package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func clientsListHandler(c *gin.Context) {
	var uuid = c.Param("user_id")

	if !security.ValidUUID(uuid) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "User's clients unavailable",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	action := c.MustGet("Action").(models.Action)
	repositories := ioc.GetRepositories(c)
	user := repositories.Users().FindByUUID(uuid)
	if user.IsNewRecord() || user.ID != action.UserID {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "User's clients unavailable",
			"error":    shared.AccessDenied,
		})
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"_status":  "success",
		"_message": "User's clients available",
		"clients":  repositories.Users().ActiveClients(user),
	})
}
