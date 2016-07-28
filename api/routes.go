package api

import (
    "net/http"
    "fmt"
    "bytes"
    "encoding/base64"
    "image/png"
    "strings"

    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/oauth"
    "github.com/earaujoassis/space/utils"
)

func ExposeRoutes(router *gin.RouterGroup) {
    users := router.Group("/users")
    {
        users.POST("/create", func(c *gin.Context) {
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

        users.POST("/introspect", func(c *gin.Context) {
            var publicId string = c.PostForm("user_id")

            authorizationBearer := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", 1)
            session := oauth.SessionAuthentication(authorizationBearer)
            if session.ID == 0 || !services.SessionGrantsReadAbility(session) {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            user := services.FindUserByPublicId(publicId)
            if user.ID == 0 || user.ID != session.UserID {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            c.JSON(http.StatusOK, utils.H{
                "user":  user,
            })
        })

        users.POST("/update", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        users.POST("/deactivate", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })
    }
    sessions := router.Group("/sessions")
    {
        sessions.POST("/create", func(c *gin.Context) {
            var holder string = c.PostForm("holder")
            var state string = c.PostForm("state")

            user := services.FindUserByAccountHolder(holder)
            client := services.FindOrCreateClient("Jupiter")
            if user.ID != 0 {
                if user.Authentic(c.PostForm("password"), c.PostForm("passcode")) {
                    session := services.CreateSession(user, client,
                        c.Request.RemoteAddr,
                        c.Request.UserAgent(),
                        models.PublicScope,
                        models.GrantToken)
                    if session.ID != 0 {
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
                "error": oauth.AccessDenied,
                "error_description": "Unauthentic user; authorization token was not created",
            })
        })

        sessions.POST("/introspect", func(c *gin.Context) {
            var token string = c.PostForm("access_token")

            authorizationBasic := strings.Replace(c.Request.Header["Authorization"][0], "Basic ", "", 1)
            client := oauth.ClientAuthentication(authorizationBasic)
            if client.ID == 0 {
                c.Header("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            session := services.FindSessionByToken(token, models.AccessToken)
            if session.ID == 0 {
                c.JSON(http.StatusNotAcceptable, utils.H{
                    "error": oauth.InvalidSession,
                })
                return
            }
            c.JSON(http.StatusOK, utils.H{
                "active":  true,
                "scope": session.Scopes,
                "client_id": session.Client.Key,
                "token_type": "Bearer",
            })
        })

        sessions.POST("/invalidate", func(c *gin.Context) {
            var token string = c.PostForm("access_token")

            authorizationBasic := strings.Replace(c.Request.Header["Authorization"][0], "Basic ", "", 1)
            client := oauth.ClientAuthentication(authorizationBasic)
            if client.ID == 0 {
                c.Header("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            session := services.FindSessionByToken(token, models.AccessToken)
            if session.ID == 0 {
                c.JSON(http.StatusNotAcceptable, utils.H{
                    "error": oauth.InvalidSession,
                })
                return
            }
            services.InvalidateSession(session)
            c.JSON(http.StatusOK, utils.H{
                "_status":  "deleted",
                "_message": "Session was deleted (soft)",
                "active":  false,
                "scope": session.Scopes,
                "client_id": session.Client.Key,
                "token_type": "Bearer",
            })
        })
    }
}
