package services

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/utils"
)

func listHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)

	c.JSON(http.StatusOK, utils.H{
		"_status":  "success",
		"_message": "Services are available",
		"services": repositories.Services().GetAll(),
	})
}
