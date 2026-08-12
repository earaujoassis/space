package clients

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/utils"
)

func createHandler(c *gin.Context) {
	repositories := ioc.GetRepositories(c)

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
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
