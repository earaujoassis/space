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

func sessionsListHandler(c *gin.Context) {
	var userUUID = c.Param("user_id")

	if !security.ValidUUID(userUUID) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Sessions are not available",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	repositories := ioc.GetRepositories(c)
	action := c.MustGet("Action").(models.Action)
	currentSession := c.MustGet("CurrentSession").(models.Session)
	user := repositories.Users().FindByUUID(userUUID)
	if user.ID != action.UserID {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Sessions are not available",
			"error":    shared.AccessDenied,
		})
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"_status":  "success",
		"_message": "Sessions are available",
		"sessions": repositories.Sessions().ApplicationSessionsWithActive(user, currentSession),
	})
}
