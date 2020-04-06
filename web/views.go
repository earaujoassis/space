package web

import (
    "bytes"
    "encoding/base64"
    "fmt"
    "image/png"
    "net/http"
    "net/url"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/renders/multitemplate"
    "github.com/gin-gonic/contrib/sessions"

    "github.com/earaujoassis/space/config"
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/oauth"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/feature"
    "github.com/earaujoassis/space/utils"
)

const (
    errorURI string = "%s?error=%s&state=%s"
)

func createCustomRender() multitemplate.Render {
    render := multitemplate.New()
    render.AddFromFiles("satellite", "web/templates/default.html", "web/templates/satellite.html")
    render.AddFromFiles("user.update.secrets", "web/templates/default.html", "web/templates/user.update.secrets.html")
    render.AddFromFiles("error.generic", "web/templates/default.html", "web/templates/error.generic.html")
    render.AddFromFiles("error.password_update", "web/templates/default.html", "web/templates/error.password_update.html")
    render.AddFromFiles("error.secrets_update", "web/templates/default.html", "web/templates/error.secrets_update.html")
    render.AddFromFiles("error.authorization", "web/templates/default.html", "web/templates/error.authorization.html")
    render.AddFromFiles("error.not_found", "web/templates/default.html", "web/templates/error.not_found.html")
    render.AddFromFiles("error.internal", "web/templates/default.html", "web/templates/error.internal.html")
    return render
}

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//      in the WEB escope
func ExposeRoutes(router *gin.Engine) {
    var cfg config.Config = config.GetGlobalConfig()
    router.LoadHTMLGlob("web/templates/*.html")
    router.HTMLRender = createCustomRender()
    router.Static("/public", "web/public")
    store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
    store.Options(sessions.Options{
        Secure: (config.IsEnvironment("production") && cfg.SessionSecure),
        HttpOnly: true,
    })
    router.Use(sessions.Sessions("jupiter.session", store))
    views := router.Group("/")
    {
        views.GET("/", func(c *gin.Context) {
            c.Redirect(http.StatusFound, "/applications")
        })
        views.GET("/applications", jupiterHandler)
        views.GET("/profile", jupiterHandler)

        views.GET("/profile/password", func(c *gin.Context) {
            var authorizationBearer = c.Query("_")
            action := services.ActionAuthentication(authorizationBearer)

            if action.UUID == "" || !services.ActionGrantsWriteAbility(action) || !action.CanUpdateUser() {
                c.HTML(http.StatusUnauthorized, "error.password_update", utils.H{
                    "Title": " - Update Resource Owner Credential",
                    "Internal": true,
                })
                return
            }

            c.HTML(http.StatusOK, "satellite", utils.H{
                "Title": " - Update Resource Owner Credential",
                "Satellite": "amalthea",
                "Internal": true,
                "Data": utils.H{
                    "action_token": action.Token,
                },
            })
        })

        views.GET("/profile/secrets", func(c *gin.Context) {
            var authorizationBearer = c.Query("_")
            var buf bytes.Buffer
            var imageData string

            action := services.ActionAuthentication(authorizationBearer)
            user := services.FindUserByID(action.UserID)
            if action.UUID == "" || !services.ActionGrantsWriteAbility(action) || !action.CanUpdateUser() || user.ID == 0 {
                c.HTML(http.StatusUnauthorized, "error.password_update", utils.H{
                    "Title": " - Update Resource Owner Credential",
                    "Internal": true,
                })
                return
            }

            codeSecretKey := user.GenerateCodeSecret()
            recoverSecret, _ := user.GenerateRecoverSecret()
            img, err := codeSecretKey.Image(200, 200)
            if err != nil {
                imageData = ""
            } else {
                png.Encode(&buf, img)
                imageData = base64.StdEncoding.EncodeToString(buf.Bytes())
            }

            dataStore := datastore.GetDataStoreConnection()
            dataStore.Save(&user)
            action.Delete()
            c.HTML(http.StatusOK, "user.update.secrets", utils.H{
                "Title": " - Update Resource Owner Credential",
                "Satellite": "amalthea",
                "Internal": true,
                "CodeSecretImage": imageData,
                "RecoveryCode": strings.Split(recoverSecret, "-"),
            })
        })

        views.GET("/signup", func(c *gin.Context) {
            c.HTML(http.StatusOK, "satellite", utils.H{
                "Title": " - Sign Up",
                "Satellite": "io",
                "UserCreateEnabled": feature.IsActive("user.create"),
                "Data": utils.H{
                    "feature.gates": utils.H{
                        "user.create": feature.IsActive("user.create"),
                    },
                },
            })
        })

        views.GET("/signin", func(c *gin.Context) {
            c.HTML(http.StatusOK, "satellite", utils.H{
                "Title": " - Sign In",
                "Satellite": "ganymede",
                "UserCreateEnabled": feature.IsActive("user.create"),
            })
        })

        views.GET("/signout", func(c *gin.Context) {
            session := sessions.Default(c)

            userPublicID := session.Get("userPublicID")
            if userPublicID != nil {
                session.Delete("userPublicID")
                session.Save()
            }

            c.Redirect(http.StatusFound, "/signin")
        })

        views.GET("/session", func(c *gin.Context) {
            session := sessions.Default(c)

            userPublicID := session.Get("userPublicID")
            if userPublicID != nil {
                c.Redirect(http.StatusFound, "/")
                return
            }

            var nextPath = "/"
            var scope = c.Query("scope")
            var grantType = c.Query("grant_type")
            var code = c.Query("code")
            var clientID = c.Query("client_id")
            var _nextPath = c.Query("_")
            //var state string = c.Query("state")

            if scope == "" || grantType == "" || code == "" || clientID == "" {
                // Original response:
                // c.String(http.StatusMethodNotAllowed, "Missing required parameters")
                c.Redirect(http.StatusFound, "/signin")
                return
            }
            if _nextPath != "" {
                if _nextPath, err := url.QueryUnescape(_nextPath); err == nil {
                    nextPath = _nextPath
                }
            }

            client := services.FindOrCreateClient("Jupiter")
            if client.Key == clientID && grantType == oauth.AuthorizationCode && scope == models.PublicScope {
                grantToken := services.FindSessionByToken(code, models.GrantToken)
                if grantToken.ID != 0 {
                    session.Set("userPublicID", grantToken.User.PublicID)
                    session.Save()
                    services.InvalidateSession(grantToken)
                    c.Redirect(http.StatusFound, nextPath)
                    return
                }
            }

            c.Redirect(http.StatusFound, "/signin")
        })

        views.GET("/authorize", authorizeHandler)
        views.POST("/authorize", authorizeHandler)

        views.GET("/error", func(c *gin.Context) {
            errorReason := c.Query("error")

            c.HTML(http.StatusOK, "error.generic", utils.H{
                "Title": " - Unexpected Error",
                "Internal": true,
                "ErrorReason": errorReason,
            })
        })

        views.POST("/token", func(c *gin.Context) {
            var grantType = c.PostForm("grant_type")

            authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)
            client := oauth.ClientAuthentication(authorizationBasic)
            if client.ID == 0 {
                c.Header("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "_status": "error",
                    "_message": "Cannot fulfill token request",
                    "error": oauth.AccessDenied,
                })
                return
            }

            switch grantType {
            // Authorization Code Grant
            case oauth.AuthorizationCode:
                result, err := oauth.AccessTokenRequest(utils.H{
                    "grant_type": grantType,
                    "code": c.PostForm("code"),
                    "redirect_uri": c.PostForm("redirect_uri"),
                    "client": client,
                })
                if err != nil {
                    c.JSON(http.StatusMethodNotAllowed, utils.H{
                        "_status": "error",
                        "_message": "Cannot fulfill token request",
                        "error": result["error"],
                    })
                    return
                }
                c.JSON(http.StatusOK, utils.H{
                    "_status":  "success",
                    "_message": "Token granted",
                    "user_id": result["user_id"],
                    "access_token": result["access_token"],
                    "token_type": result["token_type"],
                    "expires_in": result["expires_in"],
                    "refresh_token": result["refresh_token"],
                    "scope": result["scope"],
                })
                return
            // Refreshing an Access Token
            case oauth.RefreshToken:
                result, err := oauth.RefreshTokenRequest(utils.H{
                    "grant_type": grantType,
                    "refresh_token": c.PostForm("refresh_token"),
                    "scope": c.PostForm("scope"),
                    "client": client,
                })
                if err != nil {
                    c.JSON(http.StatusMethodNotAllowed, utils.H{
                        "_status": "error",
                        "_message": "Cannot fulfill token request",
                        "error": result["error"],
                    })
                    return
                }
                c.JSON(http.StatusOK, utils.H{
                    "_status":  "success",
                    "_message": "Token granted",
                    "user_id": result["user_id"],
                    "access_token": result["access_token"],
                    "token_type": result["token_type"],
                    "expires_in": result["expires_in"],
                    "refresh_token": result["refresh_token"],
                    "scope": result["scope"],
                })
                return
            // Resource Owner Password Credentials Grant
            // Client Credentials Grant
            case oauth.Password, oauth.ClientCredentials:
                c.JSON(http.StatusMethodNotAllowed, utils.H{
                    "_status": "error",
                    "_message": "Cannot fulfill token request",
                    "error": oauth.UnsupportedGrantType,
                })
                return
            default:
                c.JSON(http.StatusBadRequest, utils.H{
                    "_status": "error",
                    "_message": "Cannot fulfill token request",
                    "error": oauth.InvalidRequest,
                })
                return
            }
        })
    }
}

func jupiterHandler(c *gin.Context) {
    session := sessions.Default(c)
    userPublicID := session.Get("userPublicID")
    if userPublicID == nil {
        c.Redirect(http.StatusFound, "/signin")
        return
    }
    client := services.FindOrCreateClient("Jupiter")
    user := services.FindUserByPublicID(userPublicID.(string))
    actionToken := services.CreateAction(user, client,
        c.Request.RemoteAddr,
        c.Request.UserAgent(),
        models.ReadWriteScope,
        models.NotSpecialAction,
    )
    c.HTML(http.StatusOK, "satellite", utils.H{
        "Title": " - Mission Control",
        "Satellite": "europa",
        "Internal": true,
        "Data": utils.H {
            "action_token": actionToken.Token,
            "user_id": user.UUID,
            "user_is_admin": user.Admin,
            "feature.gates": utils.H{
                "user.adminify": feature.IsActive("user.adminify"),
            },
        },
    })
}

func authorizeHandler(c *gin.Context) {
    var location string
    var responseType string
    var clientID string
    var redirectURI string
    var scope string
    var state string

    session := sessions.Default(c)
    userPublicID := session.Get("userPublicID")
    nextPath := url.QueryEscape(fmt.Sprintf("%s?%s", c.Request.URL.Path, c.Request.URL.RawQuery))
    if userPublicID == nil {
        location = fmt.Sprintf("/signin?_=%s", nextPath)
        c.Redirect(http.StatusFound, location)
        return
    }
    user := services.FindUserByPublicID(userPublicID.(string))
    if user.ID == 0 {
        session.Delete("userPublicID")
        session.Save()
        location = fmt.Sprintf("/signin?_=%s", nextPath)
        c.Redirect(http.StatusFound, location)
        return
    }

    responseType = c.Query("response_type")
    clientID = c.Query("client_id")
    redirectURI = c.Query("redirect_uri")
    scope = c.Query("scope")
    state = c.Query("state")

    if redirectURI == "" {
        redirectURI = "/error"
    }

    client := services.FindClientByKey(clientID)
    if client.ID == 0 {
        // REFACTOR This scenario is the trickiest one
        // redirectURI = "/error"
        // location = fmt.Sprintf("%s?error=%s&state=%s", redirectURI, oauth.UnauthorizedClient, state)
        // Previous return: c.HTML(http.StatusFound, location)
        c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
            "Title": " - Authorization Error",
            "Internal": true,
            "ProceedTo": nil,
            "ErrorCode": oauth.UnauthorizedClient,
        })
        return
    }

    if scope != models.PublicScope && scope != models.ReadScope && scope != models.ReadWriteScope {
        scope = models.PublicScope
    }

    switch responseType {
    // Authorization Code Grant
    case oauth.Code:
        activeSessions := services.ActiveSessionsForClient(client.ID, user.ID)
        if c.Request.Method == "GET" && activeSessions == 0 {
            c.HTML(http.StatusOK, "satellite", utils.H{
                "Title": " - Authorize",
                "Satellite": "callisto",
                "Internal": true,
                "Data": utils.H{
                    "first_name": user.FirstName,
                    "last_name": user.LastName,
                    "client_name": client.Name,
                    "client_uri": client.CanonicalURI,
                    "requested_scope": scope,
                },
            })
            return
        } else if c.Request.Method == "POST" || (activeSessions > 0 && c.Request.Method == "GET") {
            if c.PostForm("access_denied") == "true" {
                // In this scenario, the user requested to deny access; it's not the client application's fault
                // The client application is safe, so the user may proceed (client application must handle this)
                location = fmt.Sprintf(errorURI, redirectURI, oauth.AccessDenied, state)
                c.Redirect(http.StatusFound, location)
                return
            }
            result, err := oauth.AuthorizationCodeGrant(utils.H{
                "response_type": responseType,
                "client": client,
                "user": user,
                "ip": c.Request.RemoteAddr,
                "userAgent": c.Request.UserAgent(),
                "redirect_uri": redirectURI,
                "scope": scope,
                "state": state,
            })
            if err != nil {
                location = fmt.Sprintf(errorURI, redirectURI, result["error"], result["state"])
                // Previous return: c.HTML(http.StatusFound, location)
                c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
                    "Title": " - Authorization Error",
                    "Internal": true,
                    "ProceedTo": location,
                    "ErrorCode": result["error"],
                })
            } else {
                location = fmt.Sprintf("%s?code=%s&scope=%s&state=%s",
                    redirectURI, result["code"], result["scope"], result["state"])
                c.Redirect(http.StatusFound, location)
            }
        } else {
            c.String(http.StatusNotFound, "404 Not Found")
        }
    // Implicit Grant
    case oauth.Token:
        location = fmt.Sprintf(errorURI, redirectURI, oauth.UnsupportedResponseType, state)
        // Previous return: c.HTML(http.StatusFound, location)
        c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
            "Title": " - Authorization Error",
            "Internal": true,
            "ProceedTo": location,
            "ErrorCode": oauth.UnsupportedResponseType,
        })
    default:
        location = fmt.Sprintf(errorURI, redirectURI, oauth.InvalidRequest, state)
        // Previous return: c.HTML(http.StatusFound, location)
        c.HTML(http.StatusBadRequest, "error.authorization", utils.H{
            "Title": " - Authorization Error",
            "Internal": true,
            "ProceedTo": location,
            "ErrorCode": oauth.InvalidRequest,
        })
    }
}
