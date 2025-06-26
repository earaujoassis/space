package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/utils"
)

func errorHandler(c *gin.Context) {
	errorReason := c.Query("error")

	c.HTML(http.StatusOK, "error.generic", utils.H{
		"Title":       " - Unexpected Error",
		"Internal":    true,
		"ErrorReason": errorReason,
	})
}
