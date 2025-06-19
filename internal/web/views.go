package web

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

// ExposeRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the Web scope
func ExposeRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("web/templates/*.html")
	router.HTMLRender = createCustomRender()
	router.Static("/public", "web/public")
	views := router.Group("/")
	{
		views.GET("/", satelliteHandler)
		views.GET("/applications", satelliteHandler)
		views.GET("/clients", satelliteHandler)
		views.GET("/clients/edit", satelliteHandler)
		views.GET("/clients/edit/scopes", satelliteHandler)
		views.GET("/clients/new", satelliteHandler)
		views.GET("/services", satelliteHandler)
		views.GET("/services/new", satelliteHandler)
		views.GET("/notifications", satelliteHandler)
		views.GET("/profile", satelliteHandler)
		views.GET("/security", satelliteHandler)

		views.GET("/profile/password", func(c *gin.Context) {
			authorizationBearer := c.Query("_")
			repositories := ioc.GetRepositories(c)
			action := repositories.Actions().Authentication(authorizationBearer)
			if action.UUID == "" || !action.GrantsWriteAbility() || !action.CanUpdateUser() {
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
			var buf bytes.Buffer
			var imageData string

			authorizationBearer := c.Query("_")
			repositories := ioc.GetRepositories(c)
			action := repositories.Actions().Authentication(authorizationBearer)
			user := repositories.Users().FindByID(action.UserID)
			if action.UUID == "" || !action.GrantsWriteAbility() || !action.CanUpdateUser() || user.IsNewRecord() {
				c.HTML(http.StatusUnauthorized, "error.password_update", utils.H{
					"Title":    " - Update Resource Owner Credential",
					"Internal": true,
				})
				return
			}

			codeSecretKey := repositories.Users().SetCodeSecret(&user)
			recoverSecret, _ := repositories.Users().SetRecoverSecret(&user)
			img, err := codeSecretKey.Image(200, 200)
			if err != nil {
				imageData = ""
			} else {
				png.Encode(&buf, img)
				imageData = base64.StdEncoding.EncodeToString(buf.Bytes())
			}

			repositories.Users().Save(&user)
			repositories.Actions().Delete(action)
			c.HTML(http.StatusOK, "user.update.secrets", utils.H{
				"Title":           " - Update Resource Owner Credential",
				"Satellite":       "amalthea",
				"Internal":        true,
				"CodeSecretImage": imageData,
				"RecoveryCode":    strings.Split(recoverSecret, "-"),
			})
		})

		views.GET("/profile/email_verification", func(c *gin.Context) {
			authorizationBearer := c.Query("_")
			repositories := ioc.GetRepositories(c)
			action := repositories.Actions().Authentication(authorizationBearer)
			user := repositories.Users().FindByID(action.UserID)
			if action.UUID == "" || !action.GrantsWriteAbility() || !action.CanUpdateUser() {
				c.HTML(http.StatusUnauthorized, "error.email_confirmation", utils.H{
					"Title":    " - Email Confirmation",
					"Internal": true,
				})
				return
			}

			user.EmailVerified = true

			repositories.Users().Save(&user)
			repositories.Actions().Delete(action)
			c.Redirect(http.StatusFound, "/")
		})

		views.GET("/signup", func(c *gin.Context) {
			fg := ioc.GetFeatureGate(c)
			c.HTML(http.StatusOK, "satellite", utils.H{
				"Title":             " - Sign Up",
				"Satellite":         "io",
				"UserCreateEnabled": fg.IsActive("user.create"),
				"Data": utils.H{
					"feature.gates": utils.H{
						"user.create": fg.IsActive("user.create"),
					},
				},
			})
		})

		views.GET("/signin", func(c *gin.Context) {
			fg := ioc.GetFeatureGate(c)
			c.HTML(http.StatusOK, "satellite", utils.H{
				"Title":             " - Sign In",
				"Satellite":         "ganymede",
				"UserCreateEnabled": fg.IsActive("user.create"),
			})
		})

		views.GET("/signout", func(c *gin.Context) {
			session := sessions.Default(c)

			userPublicID := session.Get("user_public_id")
			if userPublicID != nil {
				session.Delete("user_public_id")
				session.Save()
			}

			c.Redirect(http.StatusFound, "/signin")
		})

		views.GET("/session", func(c *gin.Context) {
			session := sessions.Default(c)

			userPublicID := session.Get("user_public_id")
			if userPublicID != nil {
				c.Redirect(http.StatusFound, "/")
				return
			}

			var nextPath = "/"
			var scope = c.Query("scope")
			var grantType = c.Query("grant_type")
			var code = c.Query("code")
			var clientKey = c.Query("client_id")
			var _nextPath = c.Query("_")
			//var state string = c.Query("state")

			if scope == "" || grantType == "" || code == "" || clientKey == "" {
				// Original response:
				// c.String(http.StatusBadRequest, "Missing required parameters")
				c.Redirect(http.StatusFound, "/signin")
				return
			}
			if _nextPath != "" {
				if _nextPath, err := url.QueryUnescape(_nextPath); err == nil {
					nextPath = _nextPath
				}
			}

			repositories := ioc.GetRepositories(c)
			client := repositories.Clients().FindOrCreate(models.DefaultClient)
			if client.Key == clientKey && grantType == shared.AuthorizationCode && scope == models.PublicScope {
				grantToken := repositories.Sessions().FindByToken(code, models.GrantToken)
				if grantToken.IsSavedRecord() {
					session.Set("user_public_id", grantToken.User.PublicID)
					session.Save()
					repositories.Sessions().Invalidate(&grantToken)
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
	}
}
