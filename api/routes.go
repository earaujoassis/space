package api

import (
    "net/http"
    "fmt"
    "bytes"
    "encoding/base64"
    "image/png"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/sessions"

    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/services/logger"
    "github.com/earaujoassis/space/security"
    "github.com/earaujoassis/space/policy"
    "github.com/earaujoassis/space/oauth"
    "github.com/earaujoassis/space/feature"
    "github.com/earaujoassis/space/utils"
    "github.com/earaujoassis/space/config"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//      in the REST API escope
func ExposeRoutes(router *gin.RouterGroup) {
    usersRoutes := router.Group("/users")
    {
        // Requires X-Requested-By and Origin (same-origin policy)
        usersRoutes.POST("/create", requiresConformance, func(c *gin.Context) {
            var buf bytes.Buffer
            var imageData string

            if !feature.IsActive("user.create") {
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
        usersRoutes.POST("/introspect", oAuthTokenBearerAuthorization, func(c *gin.Context) {
            var publicID = c.PostForm("user_id")

            if !security.ValidRandomString(publicID) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "error": "must use valid identification string",
                })
                return
            }
            session := c.MustGet("Session").(models.Session)
            user := services.FindUserByPublicID(publicID)
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
        usersRoutes.PATCH("/:user_id/profile", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        usersRoutes.GET("/:user_id/profile", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var uuid = c.Param("user_id")

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
                "is_admin": user.Admin,
                "username": user.Username,
                "first_name": user.FirstName,
                "last_name": user.LastName,
                "email": user.Email,
                "timezone_identifier": user.TimezoneIdentifier,
            })
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        usersRoutes.DELETE("/:user_id/deactivate", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        usersRoutes.GET("/:user_id/clients", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var uuid = c.Param("user_id")

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
        usersRoutes.DELETE("/:user_id/clients/:client_id/revoke", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var userUUID = c.Param("user_id")
            var clientUUID = c.Param("client_id")

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
        usersRoutes.PATCH("/:user_id/adminify", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var cfg config.Config = config.GetGlobalConfig()
            var uuid = c.Param("user_id")
            var providedApplicationKey = c.PostForm("application_key")

            if !feature.IsActive("user.adminify") {
                c.JSON(http.StatusForbidden, utils.H{
                    "_status": "error",
                    "_message": "User was not updated",
                    "error": "Feature is not available at this time",
                })
                return
            }

            if providedApplicationKey != cfg.ApplicationKey {
                c.JSON(http.StatusForbidden, utils.H{
                    "_status": "error",
                    "_message": "User was not updated",
                    "error": "Application key is incorrect",
                })
                return
            }

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

            dataStore := datastore.GetDataStoreConnection()
            user.Admin = true
            dataStore.Save(&user)

            c.JSON(http.StatusNoContent, nil)
        })
    }
    sessionsRoutes := router.Group("/sessions")
    {
        // Requires X-Requested-By and Origin (same-origin policy)
        sessionsRoutes.POST("/create", requiresConformance, func(c *gin.Context) {
            var holder = c.PostForm("holder")
            var state = c.PostForm("state")

            var IP = c.Request.RemoteAddr
            var userID = IP
            var statusSignInAttempts = policy.SignInAttemptStatus(IP)

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
                            "IP": session.IP,
                            "CreatedAt": session.CreatedAt.Format(time.RFC850),
                        })
                        policy.RegisterSuccessfulSignIn(user.UUID)
                        policy.RegisterSuccessfulSignIn(IP)
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
        sessionsRoutes.POST("/introspect", clientBasicAuthorization, func(c *gin.Context) {
            var token = c.PostForm("access_token")

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
        sessionsRoutes.POST("/invalidate", clientBasicAuthorization, func(c *gin.Context) {
            var token = c.PostForm("access_token")

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
    clientsRoutes := router.Group("/clients")
    {
        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        clientsRoutes.GET("/", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            session := sessions.Default(c)

            action := c.MustGet("Action").(models.Action)
            userPublicID := session.Get("userPublicID")
            user := services.FindUserByPublicID(userPublicID.(string))
            if userPublicID == nil || user.ID == 0 || user.ID != action.UserID || user.Admin != true {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
                return
            }

            c.JSON(http.StatusOK, utils.H{
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
                    "error": oauth.AccessDenied,
                })
                return
            }

            clientName := c.PostForm("name")
            clientDescription := c.PostForm("description")
            clientSecret := models.GenerateRandomString(64)
            clientScope := "public"
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
                    "error": "The client was not created",
                    "client": client,
                })
            } else {
                c.JSON(http.StatusNoContent, nil)
            }
        })
    }
}
