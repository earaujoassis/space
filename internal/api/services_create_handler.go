package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func servicesCreateHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	session := sessions.Default(c)
	action := c.MustGet("Action").(models.Action)
	userPublicID := session.Get("user_public_id")
	user := repositories.Users().FindByPublicID(userPublicID.(string))
	if userPublicID == nil || user.IsNewRecord() || user.ID != action.UserID || !user.Admin {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Service was not created",
			"error":    shared.AccessDenied,
		})
		return
	}

	serviceName := c.PostForm("name")
	serviceDescription := c.PostForm("description")
	canonicalURI := c.PostForm("canonical_uri")
	logoURI := c.PostForm("logo_uri")

	service := repositories.Services().Create(serviceName,
		serviceDescription,
		canonicalURI,
		logoURI)

	if service.IsNewRecord() {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Service was not created",
			"error":    "cannot create Service",
			"service":  service,
		})
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
