package self

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/utils"
)

func settingsListHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	user := c.MustGet("User").(models.User)

	c.JSON(http.StatusOK, utils.H{
		"_status":  "success",
		"_message": "Settings are available",
		"settings": repositories.Settings().ReduceForUser(user),
	})
}
