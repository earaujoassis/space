package api

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

func usersSessionsRevokeHandler(c *gin.Context) {
	var userUUID = c.Param("user_id")
	var sessionUUID = c.Param("session_id")

	if !security.ValidUUID(userUUID) || !security.ValidUUID(sessionUUID) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Session irrevocable",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	repositories := ioc.GetRepositories(c)
	action := c.MustGet("Action").(models.Action)
	user := repositories.Users().FindByUUID(userUUID)
	session := repositories.Sessions().FindByUUID(sessionUUID)
	if user.IsNewRecord() || user.ID != action.UserID || user.ID != session.User.ID {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Session irrevocable",
			"error":    shared.AccessDenied,
		})
		return
	}

	repositories.Sessions().Invalidate(&session)
	c.JSON(http.StatusNoContent, nil)
}
