package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/utils"
)

func groupsPatchHandler(c *gin.Context) {
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

	repositories := ioc.GetRepositories(c)

	user := repositories.Users().FindByUUID(userUUID)
	if user.IsNewRecord() {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Groups not available",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	client := repositories.Clients().FindByUUID(clientUUID)
	if client.IsNewRecord() {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Groups not available",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	group := repositories.Groups().FindOrCreate(user, client)
	tags := c.PostFormArray("tags")
	group.Tags = utils.TrimStrings(tags)
	err := repositories.Groups().Save(&group)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Groups not patched",
			"error":    "validation failed",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
