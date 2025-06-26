package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
)

func healthCheckHandler(c *gin.Context) {
	db, err := ioc.GetDB(c).DB()
	if err != nil {
		c.String(http.StatusOK, "unhealthy")
		return
	}
	if err := db.Ping(); err == nil {
		c.String(http.StatusOK, "healthy")
	} else {
		c.String(http.StatusOK, "unhealthy")
	}
}
