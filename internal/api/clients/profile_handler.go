package clients

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/utils"
)

func profileHandler(c *gin.Context) {
	clientUUID := c.Param("client_id")
	if !security.ValidUUID(clientUUID) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Client was not updated",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	repositories := ioc.GetRepositories(c)
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
