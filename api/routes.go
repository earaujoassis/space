package api

import (
    "bytes"
    "encoding/base64"
    "fmt"
    "image/png"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/contrib/sessions"
    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/space/config"
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/feature"
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/oauth"
    "github.com/earaujoassis/space/policy"
    "github.com/earaujoassis/space/security"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/services/logger"
    "github.com/earaujoassis/space/utils"
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
                    "_status":  "error",
                    "_message": "User was not created",
                    "error":    "feature is not available at this time",
                })
                return
            }

            dataStore := datastore.GetDataStoreConnection()
            user := models.User{
                FirstName:  c.PostForm("first_name"),
                LastName:   c.PostForm("last_name"),
                Username:   c.PostForm("username"),
                Email:      c.PostForm("email"),
                Passphrase: c.PostForm("password"),
            }
            if !models.IsValid("essential", user) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "User was not created",
                    "error":    "missing essential fields",
                    "user":     user,
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
                    "_status":  "error",
                    "_message": "User was not created",
                    "error": fmt.Sprintf("%v", result.GetErrors()),
                    "user":  user,
                })
            } else {
                go logger.LogAction("user.created", utils.H{
                    "Email":     user.Email,
                    "FirstName": user.FirstName,
                })
                c.JSON(http.StatusOK, utils.H{
                    "_status":           "created",
                    "_message":          "User was created",
                    "recover_secret":    recoverSecret,
                    "code_secret_image": imageData,
                    "user":              user,
                })
            }
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        usersRoutes.PATCH("/update/adminify", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var cfg config.Config = config.GetGlobalConfig()
            var uuid = c.PostForm("user_id")
            var providedApplicationKey = c.PostForm("application_key")

            if !feature.IsActive("user.adminify") {
                c.JSON(http.StatusForbidden, utils.H{
                    "_status":  "error",
                    "_message": "User was not updated",
                    "error":    "feature is not available at this time",
                })
                return
            }

            if providedApplicationKey != cfg.ApplicationKey {
                c.JSON(http.StatusForbidden, utils.H{
                    "_status":  "error",
                    "_message": "User was not updated",
                    "error":    "application key is incorrect",
                })
                return
            }

            if !security.ValidUUID(uuid) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "User was not updated",
                    "error": "must use valid UUID for identification",
                })
                return
            }

            action := c.MustGet("Action").(models.Action)
            user := services.FindUserByUUID(uuid)
            if user.ID == 0 || user.ID != action.UserID {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "User was not updated",
                    "error": oauth.AccessDenied,
                })
                return
            }

            dataStore := datastore.GetDataStoreConnection()
            user.Admin = true
            dataStore.Save(&user)
            c.JSON(http.StatusNoContent, nil)
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        usersRoutes.PATCH("/update/password", requiresConformance, func(c *gin.Context) {
            var bearer = c.PostForm("_")
            var newPassword = c.PostForm("new_password")
            var passwordConfirmation = c.PostForm("password_confirmation")

            if !security.ValidRandomString(bearer) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "User password was not updated",
                    "error": "must use valid token string",
                })
                return
            }

            action := services.ActionAuthentication(bearer)
            user := services.FindUserByID(action.UserID)
            if user.ID == 0 {
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "User password was not updated",
                    "error": "token string not valid",
                })
                return
            }

            if newPassword != passwordConfirmation {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "User password was not updated",
                    "error": "new password and password confirmation must match each other",
                })
                return
            }

            user.UpdatePassword(newPassword)
            if !models.IsValid("essential", user) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "User password was not updated",
                    "error":    "invalid password update attempt",
                    "user":     user,
                })
                return
            }

            dataStore := datastore.GetDataStoreConnection()
            dataStore.Save(&user)
            action.Delete()
            c.JSON(http.StatusNoContent, nil)
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        usersRoutes.POST("/update/request", requiresConformance, func(c *gin.Context) {
            var holder = c.PostForm("holder")
            var requestType = c.PostForm("request_type")
            var host = fmt.Sprintf("%s://%s", scheme(c.Request), c.Request.Host)

            const (
                passwordType = "password"
            )

            if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "update request was not created",
                    "error": "must use valid holder string",
                })
                return
            }

            // TODO Use case/match for new request types
            if requestType != passwordType {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "update request was not created",
                    "error": "request type not available",
                })
                return
            }

            user := services.FindUserByAccountHolder(holder)
            client := services.FindOrCreateClient("Jupiter")
            if user.ID != 0 {
                actionToken := services.CreateAction(user, client,
                    c.Request.RemoteAddr,
                    c.Request.UserAgent(),
                    models.ReadWriteScope,
                    models.UpdateUserAction,
                )
                go logger.LogAction("session.magic", utils.H{
                    "Email":     user.Email,
                    "FirstName": user.FirstName,
                    "Callback": fmt.Sprintf("%s/profile/password?_=%s", host, actionToken.Token),
                })
            }
            // No Content is the default response
            c.JSON(http.StatusNoContent, nil)
        })

        // Authorization type: access session / Bearer (for OAuth sessions)
        usersRoutes.POST("/introspect", oAuthTokenBearerAuthorization, func(c *gin.Context) {
            var publicID = c.PostForm("user_id")

            if !security.ValidRandomString(publicID) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "User instropection failed",
                    "error": "must use valid identification string",
                })
                return
            }
            session := c.MustGet("Session").(models.Session)
            user := services.FindUserByPublicID(publicID)
            if user.ID == 0 || user.ID != session.UserID {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "User instropection failed",
                    "error": oauth.AccessDenied,
                })
                return
            }

            c.JSON(http.StatusOK, utils.H{
                "_status":  "success",
                "_message": "User instropection fulfilled",
                "user": user,
            })
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        // Authorization type: action token / Bearer (for web use)
        usersRoutes.GET("/:user_id/profile", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
            var uuid = c.Param("user_id")

            if !security.ValidUUID(uuid) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "User instropection failed",
                    "error": "must use valid UUID for identification",
                })
                return
            }

            action := c.MustGet("Action").(models.Action)
            user := services.FindUserByUUID(uuid)
            if user.ID == 0 || user.ID != action.UserID {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "User instropection failed",
                    "error": oauth.AccessDenied,
                })
                return
            }

            c.JSON(http.StatusOK, utils.H{
                "_status":  "success",
                "_message": "User instropection fulfilled",
                "user": utils.H{
                    "is_admin":            user.Admin,
                    "username":            user.Username,
                    "first_name":          user.FirstName,
                    "last_name":           user.LastName,
                    "email":               user.Email,
                    "timezone_identifier": user.TimezoneIdentifier,
                },
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
                    "_status":  "error",
                    "_message": "User's clients unavailable",
                    "error": "must use valid UUID for identification",
                })
                return
            }

            action := c.MustGet("Action").(models.Action)
            user := services.FindUserByUUID(uuid)
            if user.ID == 0 || user.ID != action.UserID {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "User's clients unavailable",
                    "error": oauth.AccessDenied,
                })
                return
            }

            c.JSON(http.StatusOK, utils.H{
                "_status":  "success",
                "_message": "User's clients available",
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
                    "_status":  "error",
                    "_message": "Client application irrevocable",
                    "error": "must use valid UUID for identification",
                })
                return
            }

            action := c.MustGet("Action").(models.Action)
            user := services.FindUserByUUID(userUUID)
            if user.ID == 0 || user.ID != action.UserID {
                c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "Client application irrevocable",
                    "error": oauth.AccessDenied,
                })
                return
            }

            client := services.FindClientByUUID(clientUUID)
            services.RevokeClientAccess(client.ID, user.ID)
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
                    "_status":  "error",
                    "_message": "Session was not created",
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
                            "Email":     user.Email,
                            "FirstName": user.FirstName,
                            "IP":        session.IP,
                            "CreatedAt": session.CreatedAt.Format(time.RFC850),
                        })
                        policy.RegisterSuccessfulSignIn(user.UUID)
                        policy.RegisterSuccessfulSignIn(IP)
                        c.JSON(http.StatusOK, utils.H{
                            "_status":      "created",
                            "_message":     "Session was created",
                            "scope":        session.Scopes,
                            "grant_type":   "authorization_code",
                            "code":         session.Token,
                            "redirect_uri": "/session",
                            "client_id":    client.Key,
                            "state":        state,
                        })
                        return
                    }
                }
            }
            policy.RegisterSignInAttempt(userID)
            c.JSON(http.StatusBadRequest, utils.H{
                "_status":  "error",
                "_message": "Session was not created",
                "error":    oauth.AccessDenied,
                "attempts": statusSignInAttempts,
            })
        })

        // Requires X-Requested-By and Origin (same-origin policy)
        sessionsRoutes.POST("/magic", requiresConformance, func(c *gin.Context) {
            var holder = c.PostForm("holder")
            var next = c.PostForm("next")
            var state = c.PostForm("state")

            var host = fmt.Sprintf("%s://%s", scheme(c.Request), c.Request.Host)

            var IP = c.Request.RemoteAddr
            var userID = IP
            var statusSignInAttempts = policy.SignInAttemptStatus(IP)

            if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "Magic Session was not created",
                    "error": "must use valid holder string",
                })
                return
            }

            user := services.FindUserByAccountHolder(holder)
            client := services.FindOrCreateClient("Jupiter")
            if user.ID != 0 && statusSignInAttempts != policy.Blocked {
                userID = user.UUID
                statusSignInAttempts = policy.SignInAttemptStatus(userID)
                if statusSignInAttempts != policy.Blocked {
                    session := services.CreateSession(user, client,
                        c.Request.RemoteAddr,
                        c.Request.UserAgent(),
                        models.PublicScope,
                        models.GrantToken)
                    if session.ID != 0 {
                        go logger.LogAction("session.magic", utils.H{
                            "Email":     user.Email,
                            "FirstName": user.FirstName,
                            "CreatedAt": session.CreatedAt.Format(time.RFC850),
                            "Callback": fmt.Sprintf("%s/session?client_id=%s&code=%s&grant_type=authorization_code&scope=%s&state=%s&_=%s",
                                host, client.Key, session.Token, session.Scopes, state, next),
                        })
                        policy.RegisterSuccessfulSignIn(user.UUID)
                        policy.RegisterSuccessfulSignIn(IP)
                        c.JSON(http.StatusNoContent, nil)
                        return
                    }
                }
            }
            policy.RegisterSignInAttempt(userID)
            c.JSON(http.StatusNoContent, nil)
        })

        // Authorization type: Basic (for OAuth clients use)
        sessionsRoutes.POST("/introspect", clientBasicAuthorization, func(c *gin.Context) {
            var token = c.PostForm("access_token")

            if !security.ValidToken(token) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "Session instropection failed",
                    "error": "must use valid token string",
                })
                return
            }

            session := services.FindSessionByToken(token, models.AccessToken)
            if session.ID == 0 {
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "Session instropection failed",
                    "error": oauth.InvalidSession,
                })
                return
            }
            c.JSON(http.StatusOK, utils.H{
                "_status":  "success",
                "_message": "Session instropection fulfilled",
                "active":     true,
                "scope":      session.Scopes,
                "client_id":  session.Client.Key,
                "token_type": "Bearer",
            })
        })

        // Authorization type: Basic (for OAuth clients use)
        sessionsRoutes.POST("/invalidate", clientBasicAuthorization, func(c *gin.Context) {
            var token = c.PostForm("access_token")

            if !security.ValidToken(token) {
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status":  "error",
                    "_message": "Session invalidation failed",
                    "error": "must use valid token string",
                })
                return
            }

            session := services.FindSessionByToken(token, models.AccessToken)
            if session.ID == 0 {
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status":  "error",
                    "_message": "Session invalidation failed",
                    "error": oauth.InvalidSession,
                })
                return
            }
            services.InvalidateSession(session)
            c.JSON(http.StatusOK, utils.H{
                "_status":    "deleted",
                "_message":   "Session was deleted (soft)",
                "active":     false,
                "scope":      session.Scopes,
                "client_id":  session.Client.Key,
                "token_type": "Bearer",
            })
        })
    }
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
