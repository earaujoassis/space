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

func usersClientsRevokeHandler(c *gin.Context) {
	var userUUID = c.Param("user_id")
	var clientUUID = c.Param("client_id")

	if !security.ValidUUID(userUUID) || !security.ValidUUID(clientUUID) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Client application irrevocable",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	action := c.MustGet("Action").(models.Action)
	repositories := ioc.GetRepositories(c)
	user := repositories.Users().FindByUUID(userUUID)
	if user.IsNewRecord() || user.ID != action.UserID {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Client application irrevocable",
			"error":    shared.AccessDenied,
		})
		return
	}

	client := repositories.Clients().FindByUUID(clientUUID)
	repositories.Sessions().RevokeAccess(client, user)
	c.JSON(http.StatusNoContent, nil)
}
