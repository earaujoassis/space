package web

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/renders/multitemplate"
    "github.com/gin-gonic/contrib/sessions"

    "github.com/earaujoassis/space/config"
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/oauth"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/utils"
)

const (
    errorURI string = "%s?error=%s&state=%s"
)

func createCustomRender() multitemplate.Render {
    render := multitemplate.New()
    render.AddFromFiles("satellite", "web/templates/default.html", "web/templates/satellite.html")
    return render
}

func ExposeRoutes(router *gin.Engine) {
    router.LoadHTMLGlob(config.GetConfig("template_folder").(string))
    router.HTMLRender = createCustomRender()
    router.Static("/public", config.GetConfig("public_folder").(string))
    store := sessions.NewCookieStore([]byte(config.GetConfig("http.session_secret").(string)))
    router.Use(sessions.Sessions("jupiter", store))
    views := router.Group("/")
    {
        views.GET("/", func(c *gin.Context) {
            session := sessions.Default(c)
            userPublicId := session.Get("userPublicId")
            if userPublicId == nil {
                c.Redirect(http.StatusFound, "/signin")
                return
            }
            c.HTML(http.StatusOK, "satellite", utils.H{
                "Title": " - Mission control",
                "Satellite": "europa",
            })
        })

        views.GET("/signup", func(c *gin.Context) {
            c.HTML(http.StatusOK, "satellite", utils.H{
                "Title": " - Sign up",
                "Satellite": "io",
            })
        })

        views.GET("/signin", func(c *gin.Context) {
            c.HTML(http.StatusOK, "satellite", utils.H{
                "Title": " - Sign in",
                "Satellite": "ganymede",
            })
        })

        views.GET("/signout", func(c *gin.Context) {
            session := sessions.Default(c)

            userPublicId := session.Get("userPublicId")
            if userPublicId != nil {
                session.Delete("userPublicId")
                session.Save()
            }

            c.Redirect(http.StatusFound, "/signin")
        })

        views.GET("/session", func(c *gin.Context) {
            session := sessions.Default(c)

            userPublicId := session.Get("userPublicId")
            if userPublicId != nil {
                c.Redirect(http.StatusFound, "/")
                return
            }

            var scope string = c.Query("scope")
            var grantType string = c.Query("grant_type")
            var code string = c.Query("code")
            var clientId string = c.Query("client_id")
            //var state string = c.Query("state")

            if scope == "" || grantType == "" || code == "" || clientId == "" {
                c.String(http.StatusMethodNotAllowed, "Missing required parameters")
                return
            }

            client := services.FindOrCreateClient("Jupiter")
            dataStore := datastore.GetDataStoreConnection()
            if client.Key == clientId && grantType == oauth.AuthorizationCode && scope == models.PublicScope {
                grantToken := services.FindSessionByToken(code, models.GrantToken)
                if !dataStore.NewRecord(grantToken) {
                    session.Set("userPublicId", grantToken.User.PublicId)
                    session.Save()
                    // FIXME `InvalidateSession` is not working :(
                    services.InvalidateSession(grantToken)
                    c.Redirect(http.StatusFound, "/")
                    return
                }
            }

            c.Redirect(http.StatusFound, "/signin")
        })

        views.GET("/authorize", authorizeHandler)

        views.POST("/authorize", authorizeHandler)

        views.GET("/error", func(c *gin.Context) {
            c.String(http.StatusMethodNotAllowed, "Not implemented")
        })

        views.POST("/token", func(c *gin.Context) {
            var grantType string = c.PostForm("grant_type")

            authorizationBasic := strings.Replace(c.Request.Header["Authorization"][0], "Basic ", "", 1)
            client := oauth.ClientAuthentication(authorizationBasic)
            if client.ID == 0 {
                c.Header("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", c.Request.RequestURI))
                c.JSON(http.StatusUnauthorized, utils.H{
                    "error": oauth.AccessDenied,
                })
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
                        "error": result["error"],
                    })
                    return
                } else {
                    c.JSON(http.StatusOK, utils.H{
                        "user_id": result["user_id"],
                        "access_token": result["access_token"],
                        "token_type": result["token_type"],
                        "expires_in": result["expires_in"],
                        "refresh_token": result["refresh_token"],
                        "scope": result["scope"],
                    })
                    return
                }
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
                        "error": result["error"],
                    })
                    return
                } else {
                    c.JSON(http.StatusOK, utils.H{
                        "user_id": result["user_id"],
                        "access_token": result["access_token"],
                        "token_type": result["token_type"],
                        "expires_in": result["expires_in"],
                        "refresh_token": result["refresh_token"],
                        "scope": result["scope"],
                    })
                    return
                }
                return
            // Resource Owner Password Credentials Grant
            // Client Credentials Grant
            case oauth.Password, oauth.ClientCredentials:
                c.JSON(http.StatusMethodNotAllowed, utils.H{
                    "error": oauth.UnsupportedGrantType,
                })
                return
            default:
                c.JSON(http.StatusBadRequest, utils.H{
                    "error": oauth.InvalidRequest,
                })
                return
            }
        })
    }
}

func authorizeHandler(c *gin.Context) {
    var location string
    var responseType string
    var clientId string
    var redirectURI string
    var scope string
    var state string

    session := sessions.Default(c)
    userPublicId := session.Get("userPublicId")
    if userPublicId == nil {
        location = fmt.Sprintf("/signin?%s", c.Request.URL.RawQuery)
        c.Redirect(http.StatusFound, location)
        return
    }
    user := services.FindUserByPublicId(userPublicId.(string))
    if user.ID == 0 {
        session.Delete("userPublicId")
        session.Save()
        location = fmt.Sprintf("/signin?%s", c.Request.URL.RawQuery)
        c.Redirect(http.StatusFound, location)
        return
    }

    responseType = c.Query("response_type")
    clientId = c.Query("client_id")
    redirectURI = c.Query("redirect_uri")
    scope = c.Query("scope")
    state = c.Query("state")

    if redirectURI == "" {
        redirectURI = "/error"
    }

    client := services.FindClientByKey(clientId)
    if client.ID == 0 {
        redirectURI = "/error"
        location = fmt.Sprintf("%s?error=%s&state=%s",
            redirectURI, oauth.UnauthorizedClient, state)
        c.Redirect(http.StatusFound, location)
        return
    }

    if scope != models.PublicScope && scope != models.ReadScope && scope != models.ReadWriteScope {
        scope = "public"
    }

    switch responseType {
    // Authorization Code Grant
    case oauth.Code:
        if c.Request.Method == "GET" {
            c.HTML(http.StatusOK, "satellite", utils.H{
                "Title": " - Authorize",
                "Satellite": "callisto",
                "Data": utils.H{
                    "first_name": user.FirstName,
                    "last_name": user.LastName,
                    "client_name": client.Name,
                    "client_uri": client.DefaultRedirectURI(),
                    "requested_scope": scope,
                },
            })
            return
        } else if c.Request.Method == "POST" {
            if c.PostForm("access_denied") == "true" {
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
                c.Redirect(http.StatusFound, location)
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
        location = fmt.Sprintf(errorURI,
            redirectURI, oauth.UnsupportedResponseType, state)
        c.Redirect(http.StatusFound, location)
        return
    default:
        location = fmt.Sprintf(errorURI,
            redirectURI, oauth.InvalidRequest, state)
        c.Redirect(http.StatusFound, location)
        return
    }
}
