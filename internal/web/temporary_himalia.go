package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/utils"
)

func himaliaHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "satellite", utils.H{
		"Title":     " - Mission Control",
		"Satellite": "himalia",
		"Internal":  true,
		"Data": utils.H{},
	})
}
