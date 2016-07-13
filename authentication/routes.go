package authentication

import (
    "net/http"
    "fmt"
    "bytes"
    "encoding/base64"
    "image/png"
    "time"

    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/utils"
)

func ExposeRoutes(router *gin.RouterGroup) {
    users := router.Group("/users")
    {
        users.GET("/", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "not implemented")
        })
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
        sessions.GET("/", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })
        sessions.POST("/", func(c *gin.Context) {
            var holder string = c.PostForm("holder")
            dataStore := datastore.GetDataStoreConnection()
            user := services.FindUserByAccountHolder(holder)
            client := services.FindOrCreateClient("Jupiter")
            if !dataStore.NewRecord(user) {
                if user.Authentic(c.PostForm("password"), c.PostForm("passcode")) {
                    session := models.Session{
                        User: user,
                        Client: client,
                        Moment: time.Now().UTC().Unix(),
                        Ip: c.Request.RemoteAddr,
                        UserAgent: c.Request.UserAgent(),
                    }
                    result := dataStore.Create(&session)
                    if count := result.RowsAffected; count > 0 {
                        c.JSON(http.StatusOK, utils.H{
                            "_id": session.UUID,
                            "_status":  "created",
                            "_message": "Session was created",
                            "access_token": session.AccessToken,
                            "token_type": "bearer",
                            "expires_in": 0,
                            "refresh_token": session.RefreshToken,
                            "scope": "public",
                        })
                        return
                    }
                }
            }
            c.JSON(http.StatusNotAcceptable, utils.H{
                "_status":  "error",
                "_message": "Unauthentic user",
                "error": "access_denied",
                "error_description": "Unauthentic user; session was not created",
            })
        })
        sessions.GET("/:id", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })
        sessions.PUT("/:id", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })
        sessions.DELETE("/:id", func(c *gin.Context) {
            id := c.Param("id")
            session := services.FindSessionByUUID(id)
            dataStore := datastore.GetDataStoreConnection()
            if dataStore.NewRecord(session) {
                c.JSON(http.StatusNotAcceptable, utils.H{
                    "_status":  "error",
                    "_message": "Invalid session (not found)",
                    "error": "invalid_grant",
                    "error_description": "Session is not available nor found",
                })
                return
            }
            session.Invalidated = true
            dataStore.Save(&session)
            c.JSON(http.StatusOK, utils.H{
                "_status":  "deleted",
                "_message": "Session was deleted (soft)",
                "access_token": session.AccessToken,
                "token_type": "bearer",
                "expires_in": 0,
                "refresh_token": session.RefreshToken,
                "scope": "public",
            })
        })
    }
}
