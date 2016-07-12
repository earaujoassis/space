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
        users.GET("/", func(c *gin.Context) {
            c.String(http.StatusOK, "not implemented")
        })
        users.POST("/", func(c *gin.Context) {
            var buf bytes.Buffer
            var imageData string

            dataStore := datastore.GetDataStoreConnection()
            /*
             * TODO There should be an actionToken prior to any action.
             * An actionToken should be created whenever a client (web or remote client)
             * attempts a creational action (like creating a user or a session)
             */
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
                c.JSON(406, utils.H{
                    "_status":  "error",
                    "_message": "User was not created",
                    "_links": utils.H{
                        "rel": "self",
                        "href": c.Request.RequestURI,
                    },
                    "error": fmt.Sprintf("%v", result.GetErrors()),
                    "user": user,
                })
            } else {
                c.JSON(200, utils.H{
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
            //id := c.Param("id")
            c.String(http.StatusOK, "not implemented")
        })
        users.PUT("/:id", func(c *gin.Context) {
            //id := c.Param("id")
            c.String(http.StatusOK, "not implemented")
        })
        users.DELETE("/:id", func(c *gin.Context) {
            //id := c.Param("id")
            c.String(http.StatusOK, "not implemented")
        })
    }
}
