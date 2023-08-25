package web

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/feature"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/oauth"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

func createCustomRender() multitemplate.Render {
	render := multitemplate.New()
	render.AddFromFiles("satellite", "internal/web/templates/default.html", "internal/web/templates/satellite.html")
	render.AddFromFiles("user.update.secrets", "internal/web/templates/default.html", "internal/web/templates/user.update.secrets.html")
	render.AddFromFiles("error.generic", "internal/web/templates/default.html", "internal/web/templates/error.generic.html")
	render.AddFromFiles("error.password_update", "internal/web/templates/default.html", "internal/web/templates/error.password_update.html")
	render.AddFromFiles("error.secrets_update", "internal/web/templates/default.html", "internal/web/templates/error.secrets_update.html")
	render.AddFromFiles("error.authorization", "internal/web/templates/default.html", "internal/web/templates/error.authorization.html")
	render.AddFromFiles("error.not_found", "internal/web/templates/default.html", "internal/web/templates/error.not_found.html")
	render.AddFromFiles("error.internal", "internal/web/templates/default.html", "internal/web/templates/error.internal.html")
	return render
}

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the WEB escope
func ExposeRoutes(router *gin.Engine) {
	var cfg config.Config = config.GetGlobalConfig()
	router.LoadHTMLGlob("internal/web/templates/*.html")
	router.HTMLRender = createCustomRender()
	router.Static("/public", "web/public")
	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
	store.Options(sessions.Options{
		Secure:   (config.IsEnvironment("production") && cfg.SessionSecure),
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
					"Title":    " - Update Resource Owner Credential",
					"Internal": true,
				})
				return
			}

			c.HTML(http.StatusOK, "satellite", utils.H{
				"Title":     " - Update Resource Owner Credential",
				"Satellite": "amalthea",
				"Internal":  true,
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
					"Title":    " - Update Resource Owner Credential",
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

			services.SaveUser(&user)
			action.Delete()
			c.HTML(http.StatusOK, "user.update.secrets", utils.H{
				"Title":           " - Update Resource Owner Credential",
				"Satellite":       "amalthea",
				"Internal":        true,
				"CodeSecretImage": imageData,
				"RecoveryCode":    strings.Split(recoverSecret, "-"),
			})
		})

		views.GET("/signup", func(c *gin.Context) {
			c.HTML(http.StatusOK, "satellite", utils.H{
				"Title":             " - Sign Up",
				"Satellite":         "io",
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
				"Title":             " - Sign In",
				"Satellite":         "ganymede",
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

			client := services.FindOrCreateClient(services.DefaultClient)
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

		views.GET("/error", func(c *gin.Context) {
			errorReason := c.Query("error")

			c.HTML(http.StatusOK, "error.generic", utils.H{
				"Title":       " - Unexpected Error",
				"Internal":    true,
				"ErrorReason": errorReason,
			})
		})

		views.GET("/authorize", authorizeHandler)
		views.GET("/oauth/authorize", authorizeHandler)
		views.POST("/authorize", authorizeHandler)
		views.POST("/oauth/authorize", authorizeHandler)
		views.POST("/token", tokenHandler)
		views.POST("/oauth/token", tokenHandler)
	}
}
