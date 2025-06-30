package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func clientsProfileHandler(c *gin.Context) {
	clientUUID := c.Param("client_id")
	repositories := ioc.GetRepositories(c)
	action := c.MustGet("Action").(models.Action)
	user := c.MustGet("User").(models.User)
	if user.ID != action.UserID || !user.Admin {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Client was not updated",
			"error":    shared.AccessDenied,
		})
		return
	}

	if !security.ValidUUID(clientUUID) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Client was not updated",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	client := repositories.Clients().FindByUUID(clientUUID)
	canonicalURI := c.PostForm("canonical_uri")
	redirectURI := c.PostForm("redirect_uri")
	scopes := c.PostForm("scopes")
	if canonicalURI != "" {
		client.CanonicalURI = utils.URIs(canonicalURI)
	}
	if redirectURI != "" {
		client.RedirectURI = utils.URIs(redirectURI)
	}
	if scopes != "" {
		client.Scopes = strings.Join(utils.Scopes(scopes), " ")
	}
	err := repositories.Clients().Save(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Client was not patched",
			"error":    fmt.Sprintf("%v", err),
		})
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
