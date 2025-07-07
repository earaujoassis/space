package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func createHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	action := c.MustGet("Action").(models.Action)
	user := c.MustGet("User").(models.User)
	if user.ID != action.UserID || !user.Admin {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Service was not created",
			"error":    shared.AccessDenied,
		})
		return
	}

	service := models.Service{
		Name:         c.PostForm("name"),
		Description:  c.PostForm("description"),
		CanonicalURI: c.PostForm("canonical_uri"),
		LogoURI:      c.PostForm("logo_uri"),
		Type:         models.PublicService,
	}
	err := repositories.Services().Create(&service)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Service was not created",
			"error":    "cannot create Service",
			"service":  service,
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
