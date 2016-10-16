package api

import (
    "net/http"
    "fmt"
    "bytes"
    "encoding/base64"
    "image/png"
    "strings"
    "time"

    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/services/logger"
    "github.com/earaujoassis/space/policy"
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
                    "_status": "error",
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
                go logger.LogAction("user.created", utils.H{
                    "Email": user.Email,
                    "FirstName": user.FirstName,
                })
                c.JSON(http.StatusOK, utils.H{
                    "_status": "created",
                    "_message": "User was created",
                    "recover_secret": recoverSecret,
                    "code_secret_image": imageData,
                    "user": user,
                })
            }
        })

        // Authorization type: access session / Bearer (for OAuth sessions)
        users.POST("/introspect", func(c *gin.Context) {
            var publicId string = c.PostForm("user_id")

            authorizationBearer := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", 1)
            session := oauth.AccessAuthentication(authorizationBearer)
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
                "user": user,
            })
        })

        users.POST("/update", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        users.POST("/deactivate", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        // Authorization type: action token / Bearer (for web use)
        users.GET("/:id/clients", func(c *gin.Context) {
            var uuid string = c.Param("id")

            authorizationBearer := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", 1)
            action := services.ActionAuthentication(authorizationBearer)
            if action.UUID == "" || !services.ActionGrantsReadAbility(action) {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            user := services.FindUserByUUID(uuid)
            if user.ID == 0 || user.ID != action.UserID {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            c.JSON(http.StatusOK, utils.H{
                "clients": services.ActiveClientsForUser(user.ID),
            })
        })

        // Authorization type: action token / Bearer (for web use)
        users.DELETE("/:user_id/clients/:client_id/revoke", func(c *gin.Context) {
            var userUUID string = c.Param("user_id")
            var clientUUID string = c.Param("client_id")

            authorizationBearer := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", 1)
            action := services.ActionAuthentication(authorizationBearer)
            if action.UUID == "" || !services.ActionGrantsReadAbility(action) {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            user := services.FindUserByUUID(userUUID)
            if user.ID == 0 || user.ID != action.UserID {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            client := services.FindClientByUUID(clientUUID)
            services.RevokeClientAccess(client.ID, user.ID)

            c.Status(http.StatusNoContent)
        })

        // Authorization type: action token / Bearer (for web use)
        users.GET("/:id/profile", func(c *gin.Context) {
            var uuid string = c.Param("id")

            authorizationBearer := strings.Replace(c.Request.Header["Authorization"][0], "Bearer ", "", 1)
            action := services.ActionAuthentication(authorizationBearer)
            if action.UUID == "" || !services.ActionGrantsReadAbility(action) {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            user := services.FindUserByUUID(uuid)
            if user.ID == 0 || user.ID != action.UserID {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            c.JSON(http.StatusOK, utils.H{
                "username": user.Username,
                "first_name": user.FirstName,
                "last_name": user.LastName,
                "email": user.Email,
                "timezone_identifier": user.TimezoneIdentifier,
            })
        })
    }
    sessions := router.Group("/sessions")
    {
        sessions.POST("/create", func(c *gin.Context) {
            var holder string = c.PostForm("holder")
            var state string = c.PostForm("state")

            var Ip string = c.Request.RemoteAddr
            var userID string = Ip
            var statusSignInAttempts = policy.SignInAttemptStatus(Ip)

            user := services.FindUserByAccountHolder(holder)
            client := services.FindOrCreateClient("Jupiter")
            if user.ID != 0 && statusSignInAttempts != policy.Blocked {
                userID = user.UUID
                statusSignInAttempts = policy.SignInAttemptStatus(userID)
                if user.Authentic(c.PostForm("password"), c.PostForm("passcode")) && statusSignInAttempts != policy.Blocked {
                    session := services.CreateSession(user, client,
                        c.Request.RemoteAddr,
                        c.Request.UserAgent(),
                        models.PublicScope,
                        models.GrantToken)
                    if session.ID != 0 {
                        go logger.LogAction("session.created", utils.H{
                            "Email": user.Email,
                            "FirstName": user.FirstName,
                            "Ip": session.Ip,
                            "CreatedAt": session.CreatedAt.Format(time.RFC850),
                        })
                        policy.RegisterSuccessfulSignIn(user.UUID)
                        policy.RegisterSuccessfulSignIn(Ip)
                        c.JSON(http.StatusOK, utils.H{
                            "_status": "created",
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
            policy.RegisterSignInAttempt(userID)
            c.JSON(http.StatusNotAcceptable, utils.H{
                "error": oauth.AccessDenied,
                "error_description": "Unauthentic user; authorization token was not created",
                "attempts": statusSignInAttempts,
            })
        })

        // Authorization type: Basic (for clients use)
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
                "active": true,
                "scope": session.Scopes,
                "client_id": session.Client.Key,
                "token_type": "Bearer",
            })
        })

        // Authorization type: Basic (for clients use)
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
                "_status": "deleted",
                "_message": "Session was deleted (soft)",
                "active": false,
                "scope": session.Scopes,
                "client_id": session.Client.Key,
                "token_type": "Bearer",
            })
        })
    }
}
