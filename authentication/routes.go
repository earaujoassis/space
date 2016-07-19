package authentication

import (
    "net/http"
    "fmt"
    "bytes"
    "encoding/base64"
    "image/png"

    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/utils"
)

func ExposeRoutes(router *gin.RouterGroup) {
    users := router.Group("/users")
    {
        users.POST("/", func(c *gin.Context) {
            var buf bytes.Buffer
            var imageData string

            dataStore := datastore.GetDataStoreConnection()
            user := models.User{
                FirstName: c.PostForm("first_name"),
                LastName: c.PostForm("last_name"),
                Username: c.PostForm("username"),
                Email: c.PostForm("email"),
                Passphrase: c.PostForm("password"),
            }
            if !models.IsValid("essential", user) {
                c.JSON(http.StatusNotAcceptable, utils.H{
                    "_status":  "error",
                    "_message": "User was not created",
                    "error": "Missing essential fields",
                    "user": user,
                })
                return
            }
            if dataStore.NewRecord(user.Client) {
                user.Client = services.FindOrCreateClient("Jupiter")
            }
            if dataStore.NewRecord(user.Language) {
                user.Language = services.FindOrCreateLanguage("English", "en-US")
            }
            codeSecretKey := user.GenerateCodeSecret()
            recoverSecret := user.GenerateRecoverSecret()
            img, err := codeSecretKey.Image(200, 200)
            if err != nil {
                imageData = ""
            } else {
                png.Encode(&buf, img)
                imageData = base64.StdEncoding.EncodeToString(buf.Bytes())
            }

            result := dataStore.Create(&user)
            if count := result.RowsAffected; count < 1 {
                c.JSON(http.StatusNotAcceptable, utils.H{
                    "_status":  "error",
                    "_message": "User was not created",
                    "error": fmt.Sprintf("%v", result.GetErrors()),
                    "user": user,
                })
            } else {
                c.JSON(http.StatusOK, utils.H{
                    "_status":  "created",
                    "_message": "User was created",
                    "_links": utils.H{
                        "rel": "self",
                        "href": fmt.Sprintf("%s/%s", c.Request.RequestURI, user.PublicId),
                    },
                    "recover_secret": recoverSecret,
                    "code_secret_image": imageData,
                    "user": user,
                })
            }
        })

        users.GET("/:id", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        users.PUT("/:id", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        users.DELETE("/:id", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })
    }
    sessions := router.Group("/sessions")
    {
        sessions.POST("/", func(c *gin.Context) {
            var holder string = c.PostForm("holder")
            var state string = c.PostForm("state")

            dataStore := datastore.GetDataStoreConnection()
            user := services.FindUserByAccountHolder(holder)
            client := services.FindOrCreateClient("Jupiter")
            if !dataStore.NewRecord(user) {
                if user.Authentic(c.PostForm("password"), c.PostForm("passcode")) {
                    session := models.Session{
                        User: user,
                        Client: client,
                        Ip: c.Request.RemoteAddr,
                        UserAgent: c.Request.UserAgent(),
                        Scopes: models.PublicScope,
                        TokenType: models.GrantToken,
                    }
                    result := dataStore.Create(&session)
                    if count := result.RowsAffected; count > 0 {
                        c.JSON(http.StatusOK, utils.H{
                            "_status":  "created",
                            "_message": "Session was created",
                            "scope": session.Scopes,
                            "grant_type": "authorization_code",
                            "code": session.Token,
                            "redirect_uri": "/session",
                            "client_id": client.Key,
                            "state": state,
                        })
                        return
                    }
                }
            }
            c.JSON(http.StatusNotAcceptable, utils.H{
                "_status":  "error",
                "_message": "Unauthentic user",
                "error": "access_denied",
                "error_description": "Unauthentic user; authorization token was not created",
            })
        })

        sessions.GET("/:token", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        sessions.PUT("/:token", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        sessions.DELETE("/:token", func(c *gin.Context) {
            var token string = c.Param("token")

            session := services.FindSessionByToken(token, models.GrantToken)
            if session.ID == 0 {
                c.JSON(http.StatusNotAcceptable, utils.H{
                    "_status":  "error",
                    "_message": "Invalid session (not found)",
                    "error": "invalid_session",
                    "error_description": "Session is not available nor found",
                })
                return
            }
            services.InvalidateSession(session)
            c.JSON(http.StatusOK, utils.H{
                "_status":  "deleted",
                "_message": "Session was deleted (soft)",
            })
        })
    }
}
