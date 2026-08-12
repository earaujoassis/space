package helpers

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/utils"
)

func RequirePermission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if slices.Contains(permissions, "role:admin") {
			user := c.MustGet("User").(models.User)
			if !user.Admin {
				c.JSON(http.StatusForbidden, utils.H{
					"error": "authorization error",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
