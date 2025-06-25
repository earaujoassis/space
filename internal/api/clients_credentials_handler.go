package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func clientsCredentialsHandler(c *gin.Context) {
	clientUUID := c.Param("client_id")
	repositories := ioc.GetRepositories(c)
	session := sessions.Default(c)
	userPublicID := session.Get("user_public_id")
	user := repositories.Users().FindByPublicID(userPublicID.(string))
	if userPublicID == nil || user.IsNewRecord() || !user.Admin {
		c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
		c.JSON(http.StatusUnauthorized, utils.H{
			"_status":  "error",
			"_message": "Client credentials are not available",
			"error":    shared.AccessDenied,
		})
		return
	}

	if !security.ValidUUID(clientUUID) {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "Client credentials are not available",
			"error":    "must use valid UUID for identification",
		})
		return
	}

	client := repositories.Clients().FindByUUID(clientUUID)
	// For security reasons, the client's secret is regenerated
	clientSecret := models.GenerateRandomString(64)
	client.SetSecret(clientSecret)
	repositories.Clients().Save(&client)

	contentString := fmt.Sprintf("name,client_key,client_secret\n%s,%s,%s\n", client.Name, client.Key, clientSecret)
	content := strings.NewReader(contentString)
	contentLength := int64(len(contentString))
	contentType := "text/csv"

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="credentials.csv"`,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, content, extraHeaders)
}
