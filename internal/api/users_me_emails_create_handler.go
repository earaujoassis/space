package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func usersMeEmailsCreateHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	action := c.MustGet("Action").(models.Action)
	user := c.MustGet("User").(models.User)
	if user.ID != action.UserID {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Email was not created",
			"error":    shared.AccessDenied,
		})
		return
	}

	email := models.Email{
		User:    user,
		Address: c.PostForm("address"),
	}
	err := repositories.Emails().Create(&email)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Email was not created",
			"error":    "cannot create Email",
			"email":    email,
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
