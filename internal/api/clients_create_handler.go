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

func clientsCreateHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)
	session := sessions.Default(c)
	action := c.MustGet("Action").(models.Action)
	userPublicID := session.Get("user_public_id")
	user := repositories.Users().FindByPublicID(userPublicID.(string))
	if userPublicID == nil || user.IsNewRecord() || user.ID != action.UserID || !user.Admin {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Client was not created",
			"error":    shared.AccessDenied,
		})
		return
	}

	client := models.Client{
		Name:         c.PostForm("name"),
		Description:  c.PostForm("description"),
		Scopes:       models.PublicScope,
		CanonicalURI: utils.URIs(c.PostForm("canonical_uri")),
		RedirectURI:  utils.URIs(c.PostForm("redirect_uri")),
		Type:         models.ConfidentialClient,
	}

	err := repositories.Clients().Create(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Client was not created",
			"error":    fmt.Sprintf("%v", err),
			"client":   client,
		})
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
