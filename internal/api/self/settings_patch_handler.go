package self

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

const (
	especialKeyEmailAddress string = "notifications.system-email-notifications.email-address"
)

func settingsPatchHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	action := c.MustGet("Action").(models.Action)
	user := c.MustGet("User").(models.User)
	if user.ID != action.UserID {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Settings are not available",
			"error":    shared.AccessDenied,
		})
		return
	}

	key := c.PostForm("key")
	parts := strings.Split(key, ".")
	if len(parts) != 3 {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Setting was not created",
			"error":    "invalid setting key",
		})
		return
	}

	switch key {
	case especialKeyEmailAddress:
		email := c.PostForm("value")
		owner := repositories.Users().HoldsEmail(user, email)
		if !owner {
			c.JSON(http.StatusBadRequest, utils.H{
				"_status":  "error",
				"_message": "Setting was not created",
				"error":    "invalid setting",
			})
			return
		}
		setting := repositories.Settings().FindOrDefault(user, parts[0], parts[1], parts[2])
		setting.Value = email
		var err error
		if setting.IsNewRecord() {
			err = repositories.Settings().Create(&setting)
		} else {
			err = repositories.Settings().Save(&setting)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.H{
				"_status":  "error",
				"_message": "Setting was not created",
				"error":    "invalid setting",
			})
			return
		}
	default:
		setting := repositories.Settings().FindOrDefault(user, parts[0], parts[1], parts[2])
		setting.Value = c.PostForm("value")
		var err error
		if setting.IsNewRecord() {
			err = repositories.Settings().Create(&setting)
		} else {
			err = repositories.Settings().Save(&setting)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.H{
				"_status":  "error",
				"_message": "Setting was not created",
				"error":    "invalid setting",
			})
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)
}
