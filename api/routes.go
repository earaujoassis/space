package api

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
    "github.com/earaujoassis/space/services/logger"
    "github.com/earaujoassis/space/security"
    "github.com/earaujoassis/space/policy"
    "github.com/earaujoassis/space/oauth"
    "github.com/earaujoassis/space/feature"
    "github.com/earaujoassis/space/utils"
)

func ExposeRoutes(router *gin.RouterGroup) {
    users := router.Group("/users")
    {
        // Requires X-Requested-By and Origin (same-origin policy)
        users.POST("/create", requiresConformance, func(c *gin.Context) {
            var buf bytes.Buffer
            var imageData string

            if !feature.Active("user.create") {
                c.JSON(http.StatusForbidden, utils.H{
                    "_status": "error",
                    "_message": "User was not created",
                    "error": "Feature is not available at this time",
                })
                return
            }

            dataStore := datastore.GetDataStoreConnection()
            user := models.User{
                FirstName: c.PostForm("first_name"),
                LastName: c.PostForm("last_name"),
                Username: c.PostForm("username"),
                Email: c.PostForm("email"),
                Passphrase: c.PostForm("password"),
            }
            if !models.IsValid("essential", user) {
                c.JSON(http.StatusBadRequest, utils.H{
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
                c.JSON(http.StatusBadRequest, utils.H{
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
        users.POST("/introspect", oAuthTokenBearerAuthorization, func(c *gin.Context) {
            var publicId string = c.PostForm("user_id")

            if !security.ValidRandomString(publicId) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "error": "must use valid identification string",
                })
                return
            }

            session := c.MustGet("Session").(models.Session)
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

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        users.POST("/update", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        users.POST("/deactivate", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        users.GET("/:id/clients", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var uuid string = c.Param("id")

            if !security.ValidUUID(uuid) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "error": "must use valid UUID for identification",
                })
                return
            }

            action := c.MustGet("Action").(models.Action)
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

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        users.DELETE("/:user_id/clients/:client_id/revoke", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var userUUID string = c.Param("user_id")
            var clientUUID string = c.Param("client_id")

            if !security.ValidUUID(userUUID) || !security.ValidUUID(clientUUID) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "error": "must use valid UUID for identification",
                })
                return
            }

            action := c.MustGet("Action").(models.Action)
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

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        users.GET("/:id/profile", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var uuid string = c.Param("id")

            if !security.ValidUUID(uuid) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "error": "must use valid UUID for identification",
                })
                return
            }

            action := c.MustGet("Action").(models.Action)
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
        // Requires X-Requested-By and Origin (same-origin policy)
        sessions.POST("/create", requiresConformance, func(c *gin.Context) {
            var holder string = c.PostForm("holder")
            var state string = c.PostForm("state")

            var Ip string = c.Request.RemoteAddr
            var userID string = Ip
            var statusSignInAttempts = policy.SignInAttemptStatus(Ip)

            if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "error": "must use valid holder string",
                })
                return
            }

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
            c.JSON(http.StatusBadRequest, utils.H{
                "error": oauth.AccessDenied,
                "error_description": "Unauthentic user; authorization token was not created",
                "attempts": statusSignInAttempts,
            })
        })

        // Authorization type: Basic (for OAuth clients use)
        sessions.POST("/introspect", clientBasicAuthorization, func(c *gin.Context) {
            var token string = c.PostForm("access_token")

            if !security.ValidToken(token) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "error": "must use valid token string",
                })
                return
            }

            session := services.FindSessionByToken(token, models.AccessToken)
            if session.ID == 0 {
                c.JSON(http.StatusUnauthorized, utils.H{
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

        // Authorization type: Basic (for OAuth clients use)
        sessions.POST("/invalidate", clientBasicAuthorization, func(c *gin.Context) {
            var token string = c.PostForm("access_token")

            if !security.ValidToken(token) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "error": "must use valid token string",
                })
                return
            }

            session := services.FindSessionByToken(token, models.AccessToken)
            if session.ID == 0 {
                c.JSON(http.StatusUnauthorized, utils.H{
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
