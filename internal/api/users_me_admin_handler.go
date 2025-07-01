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

func usersMeAdminHandler(c *gin.Context) {
	var uuid = c.PostForm("user_id")
	var providedApplicationKey = c.PostForm("application_key")

	fg := ioc.GetFeatureGate(c)
	if !fg.IsActive("user.adminify") {
		c.JSON(http.StatusForbidden, utils.H{
			"_status":  "error",
			"_message": "User was not updated",
			"error":    "feature is not available at this time",
		})
		return
	}

	cfg := ioc.GetConfig(c)
	if providedApplicationKey != cfg.ApplicationKey {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "User was not updated",
			"error":    "application key is incorrect",
		})
		return
	}

	if !security.ValidUUID(uuid) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "User was not updated",
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
			"_message": "User was not updated",
			"error":    shared.AccessDenied,
		})
		return
	}

	user.Admin = true
	repositories.Users().Save(&user)
	c.JSON(http.StatusNoContent, nil)
}
