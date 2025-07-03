package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/utils"
)

func usersMePasswordHandler(c *gin.Context) {
	var bearer = c.PostForm("_")
	var newPassword = c.PostForm("new_password")
	var passwordConfirmation = c.PostForm("password_confirmation")

	if !security.ValidRandomString(bearer) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "User password was not updated",
			"error":    "must use valid token field",
		})
		return
	}

	repositories := ioc.GetRepositories(c)
	action := repositories.Actions().Authentication(bearer)
	if action.UUID == "" || !action.GrantsWriteAbility() || !action.CanUpdateUser() {
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "User password was not updated",
			"error":    "invalid token field",
		})
		return
	}

	user := repositories.Users().FindByID(action.UserID)
	if user.IsNewRecord() {
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "User password was not updated",
			"error":    "token string not valid",
		})
		return
	}

	if newPassword != passwordConfirmation {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "User password was not updated",
			"error":    "new password and password confirmation must match each other",
		})
		return
	}

	err := repositories.Users().SetPassword(&user, newPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "User password was not updated",
			"error":    "invalid password update attempt",
			"user":     user,
		})
		return
	}
	repositories.Users().Save(&user)
	repositories.Actions().Delete(action)
	c.JSON(http.StatusNoContent, nil)
}
