package services

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/utils"
)

func createHandler(c *gin.Context) {
	service := models.Service{
		Name:         c.PostForm("name"),
		Description:  c.PostForm("description"),
		CanonicalURI: c.PostForm("canonical_uri"),
		LogoURI:      c.PostForm("logo_uri"),
		Type:         models.PublicService,
	}
	repositories := ioc.GetRepositories(c)
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
