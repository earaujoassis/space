package api

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/contrib/sessions"
    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/oauth"
    "github.com/earaujoassis/space/security"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/utils"
)

// exposeClientsRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//      in the REST API escope, for the clients resource
func exposeClientsRoutes(router *gin.RouterGroup) {
    clientsRoutes := router.Group("/clients")
    {
        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        clientsRoutes.GET("/", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            action := c.MustGet("Action").(models.Action)
            session := sessions.Default(c)
            userPublicID := session.Get("userPublicID")
            user := services.FindUserByPublicID(userPublicID.(string))
            if userPublicID == nil || user.ID == 0 || user.ID != action.UserID || user.Admin != true {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "Clients are not available",
                    "error": oauth.AccessDenied,
                })
                return
            }

            c.JSON(http.StatusOK, utils.H{
                "_status":  "success",
                "_message": "Clients are available",
                "clients": services.ActiveClients(),
            })
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        clientsRoutes.POST("/create", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            session := sessions.Default(c)
            action := c.MustGet("Action").(models.Action)
            userPublicID := session.Get("userPublicID")
            user := services.FindUserByPublicID(userPublicID.(string))
            if userPublicID == nil || user.ID == 0 || user.ID != action.UserID || user.Admin != true {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "Client was not created",
                    "error": oauth.AccessDenied,
                })
                return
            }

            clientName := c.PostForm("name")
            clientDescription := c.PostForm("description")
            clientSecret := models.GenerateRandomString(64)
            clientScope := models.PublicScope
            canonicalURI := c.PostForm("canonical_uri")
            redirectURI := c.PostForm("redirect_uri")

            client := services.CreateNewClient(clientName,
                clientDescription,
                clientSecret,
                clientScope,
                canonicalURI,
                redirectURI)

            if client.ID == 0 {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "Client was not created",
                    "error":    "cannot create Client",
                    "client":   client,
                })
            } else {
                c.JSON(http.StatusNoContent, nil)
            }
        })

        // In order to avoid an overhead in this endpoint, it relies only on the cookies session data to guarantee security
        // TODO Improve security for this endpoint avoiding any overhead
        clientsRoutes.PATCH("/:client_id/profile", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var clientUUID = c.Param("client_id")

            session := sessions.Default(c)
            action := c.MustGet("Action").(models.Action)
            userPublicID := session.Get("userPublicID")
            user := services.FindUserByPublicID(userPublicID.(string))
            if userPublicID == nil || user.ID == 0 || user.ID != action.UserID || user.Admin != true {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "Client was not updated",
                    "error": oauth.AccessDenied,
                })
                return
            }

            if !security.ValidUUID(clientUUID) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "Client was not updated",
                    "error": "must use valid UUID for identification",
                })
                return
            }

            var newScopes = c.PostForm("scopes")
            // Clients can only have read or public scopes
            if (newScopes != models.PublicScope && newScopes != models.ReadScope) {
                newScopes = ""
            }

            client := services.FindClientByUUID(clientUUID)
            client.CanonicalURI = c.PostForm("canonical_uri")
            client.RedirectURI = c.PostForm("redirect_uri")
            if (newScopes != "") {
                client.Scopes = newScopes
            }
            dataStore := datastore.GetDataStoreConnection()
            dataStore.Save(&client)
            c.JSON(http.StatusNoContent, nil)
        })

        // In order to avoid an overhead in this endpoint, it relies only on the cookies session data to guarantee security
        // TODO Improve security for this endpoint avoiding any overhead
        clientsRoutes.GET("/:client_id/credentials", func(c *gin.Context) {
            var clientUUID = c.Param("client_id")

            session := sessions.Default(c)
            userPublicID := session.Get("userPublicID")
            user := services.FindUserByPublicID(userPublicID.(string))
            if userPublicID == nil || user.ID == 0 || user.Admin != true {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "Client credentials are not available",
                    "error": oauth.AccessDenied,
                })
                return
            }

            if !security.ValidUUID(clientUUID) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "Client credentials are not available",
                    "error": "must use valid UUID for identification",
                })
                return
            }

            client := services.FindClientByUUID(clientUUID)
            // For security reasons, the client's secret is regenerated
            clientSecret := models.GenerateRandomString(64)
            client.UpdateSecret(clientSecret)
            dataStore := datastore.GetDataStoreConnection()
            dataStore.Save(&client)

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
