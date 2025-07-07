package self

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/utils"
)

func emailsCreateHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	user := c.MustGet("User").(models.User)
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
