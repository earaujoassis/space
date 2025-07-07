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

func profileHandler(c *gin.Context) {
	var uuid = c.Param("user_id")

	if !security.ValidUUID(uuid) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "User instropection failed",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	action := c.MustGet("Action").(models.Action)
	repositories := ioc.GetRepositories(c)
	user := repositories.Users().FindByUUID(uuid)
	if user.ID != action.UserID {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "User instropection failed",
			"error":    shared.AccessDenied,
		})
		return
	}

	c.JSON(http.StatusOK, utils.H{
		"_status":  "success",
		"_message": "User instropection fulfilled",
		"user": utils.H{
			"is_admin":            user.Admin,
			"username":            user.Username,
			"first_name":          user.FirstName,
			"last_name":           user.LastName,
			"email":               user.Email,
			"email_verified":      user.EmailVerified,
			"timezone_identifier": user.TimezoneIdentifier,
		},
	})
}
