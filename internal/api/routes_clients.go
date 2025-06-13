package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

// exposeClientsRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the clients resource
func exposeClientsRoutes(router *gin.RouterGroup) {
	// Requires X-Requested-By and Origin (same-origin policy)
	// Authorization type: action token / Bearer (for web use)
	router.GET("/clients", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
		action := c.MustGet("Action").(models.Action)
		session := sessions.Default(c)
		userPublicID := session.Get("user_public_id")
		user := services.FindUserByPublicID(userPublicID.(string))
		if userPublicID == nil || user.ID == 0 || user.ID != action.UserID || !user.Admin {
			c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
			c.JSON(http.StatusUnauthorized, utils.H{
				"_status":  "error",
				"_message": "Clients are not available",
				"error":    shared.AccessDenied,
			})
			return
		}

		c.JSON(http.StatusOK, utils.H{
			"_status":  "success",
			"_message": "Clients are available",
			"clients":  services.ActiveClients(),
		})
	})

	clientsRoutes := router.Group("/clients")
	{
		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		clientsRoutes.POST("/create", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
			session := sessions.Default(c)
			action := c.MustGet("Action").(models.Action)
			userPublicID := session.Get("user_public_id")
			user := services.FindUserByPublicID(userPublicID.(string))
			if userPublicID == nil || user.ID == 0 || user.ID != action.UserID || !user.Admin {
				c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "Client was not created",
					"error":    shared.AccessDenied,
				})
				return
			}

			client:= models.Client{
				Name:         c.PostForm("name"),
				Description:  c.PostForm("description"),
				Scopes:       models.PublicScope,
				CanonicalURI: strings.Split(c.PostForm("canonical_uri"), "\n"),
				RedirectURI:  strings.Split(c.PostForm("redirect_uri"), "\n"),
				Type:         models.ConfidentialClient,
			}

			ok, err := services.CreateNewClient(&client)
			if !ok {
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "Client was not created",
					"error":    fmt.Sprintf("%v", err),
					"client":   client,
				})
			} else {
				c.JSON(http.StatusNoContent, nil)
			}
		})

		// In order to avoid an overhead in this endpoint, it relies only on the cookies session data to guarantee security
		// Authorization type: action token / Bearer (for web use)
		// TODO Improve security for this endpoint avoiding any overhead
		clientsRoutes.PATCH("/:client_id/profile", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
			var clientUUID = c.Param("client_id")

			session := sessions.Default(c)
			action := c.MustGet("Action").(models.Action)
			userPublicID := session.Get("user_public_id")
			user := services.FindUserByPublicID(userPublicID.(string))
			if userPublicID == nil || user.ID == 0 || user.ID != action.UserID || !user.Admin {
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

			var newScopes = c.PostForm("scopes")
			// Clients can only have read or public scopes
			if newScopes != models.PublicScope && newScopes != models.ReadScope {
				newScopes = ""
			}

			client := services.FindClientByUUID(clientUUID)
			client.CanonicalURI = utils.TrimStrings(strings.Split(c.PostForm("canonical_uri"), "\n"))
			client.RedirectURI = utils.TrimStrings(strings.Split(c.PostForm("redirect_uri"), "\n"))
			if newScopes != "" {
				client.Scopes = newScopes
			}
			services.SaveClient(&client)
			c.JSON(http.StatusNoContent, nil)
		})

		// In order to avoid an overhead in this endpoint, it relies only on the cookies session data to guarantee security
		// TODO Improve security for this endpoint avoiding any overhead
		clientsRoutes.GET("/:client_id/credentials", func(c *gin.Context) {
			var clientUUID = c.Param("client_id")

			session := sessions.Default(c)
			userPublicID := session.Get("user_public_id")
			user := services.FindUserByPublicID(userPublicID.(string))
			if userPublicID == nil || user.ID == 0 || !user.Admin {
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

			client := services.FindClientByUUID(clientUUID)
			// For security reasons, the client's secret is regenerated
			clientSecret := models.GenerateRandomString(64)
			client.UpdateSecret(clientSecret)
			services.SaveClient(&client)

			contentString := fmt.Sprintf("name,client_key,client_secret\n%s,%s,%s\n", client.Name, client.Key, clientSecret)
			content := strings.NewReader(contentString)
			contentLength := int64(len(contentString))
			contentType := "text/csv"

			extraHeaders := map[string]string{
				"Content-Disposition": `attachment; filename="credentials.csv"`,
			}

			c.DataFromReader(http.StatusOK, contentLength, contentType, content, extraHeaders)
		})
	}
}
