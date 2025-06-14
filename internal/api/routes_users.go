package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/feature"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/services/communications"
	"github.com/earaujoassis/space/internal/utils"
)

// exposeUsersRoutes defines and exposes HTTP routes for a given gin.RouterGroup
//
//	in the REST API scope, for the users resource
func exposeUsersRoutes(router *gin.RouterGroup) {
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
			codeSecretKey := user.GenerateCodeSecret()
			recoverSecret, _ := user.GenerateRecoverSecret()
			img, err := codeSecretKey.Image(200, 200)
			if err != nil {
				imageData = ""
			} else {
				png.Encode(&buf, img)
				imageData = base64.StdEncoding.EncodeToString(buf.Bytes())
			}

			ok, err := services.CreateNewUser(&user)
			if !ok {
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "User was not created",
					"error":    fmt.Sprintf("%v", err),
					"user":     user,
				})
			} else {
				go communications.Announce("user.created", utils.H{
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
					"error":    "must use valid UUID for identification",
				})
				return
			}

			action := c.MustGet("Action").(models.Action)
			user := services.FindUserByUUID(uuid)
			if user.IsNewRecord() || user.ID != action.UserID {
				c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "User was not updated",
					"error":    shared.AccessDenied,
				})
				return
			}

			user.Admin = true
			services.SaveUser(&user)
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
					"error":    "must use valid token string",
				})
				return
			}

			action := services.ActionAuthentication(bearer)
			if action.UUID == "" || !services.ActionGrantsWriteAbility(action) || !action.CanUpdateUser() {
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "User password was not updated",
					"error":    "token string not valid",
				})
				return
			}

			user := services.FindUserByID(action.UserID)
			if user.IsNewRecord() {
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "User password was not updated",
					"error":    "token string not valid",
				})
				return
			}

			if newPassword != passwordConfirmation {
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "User password was not updated",
					"error":    "new password and password confirmation must match each other",
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

			services.SaveUser(&user)
			action.Delete()
			c.JSON(http.StatusNoContent, nil)
		})

		// Requires X-Requested-By and Origin (same-origin policy)
		usersRoutes.POST("/update/request", requiresConformance, func(c *gin.Context) {
			var holder = c.PostForm("holder")
			var requestType = c.PostForm("request_type")
			var host = fmt.Sprintf("%s://%s", shared.Scheme(c.Request), c.Request.Host)

			const (
				passwordType = "password"
				secretsType  = "secrets"
			)

			if !security.ValidEmail(holder) && !security.ValidRandomString(holder) {
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "update request was not created",
					"error":    "must use valid holder string",
				})
				return
			}

			switch requestType {
			case passwordType:
				user := services.FindUserByAccountHolder(holder)
				client := services.FindOrCreateClient(services.DefaultClient)
				if user.ID != 0 {
					actionToken := services.CreateAction(user, client,
						c.Request.RemoteAddr,
						c.Request.UserAgent(),
						models.WriteScope,
						models.UpdateUserAction,
					)
					go communications.Announce("session.magic", utils.H{
						"Email":     user.Email,
						"FirstName": user.FirstName,
						"Callback":  fmt.Sprintf("%s/profile/password?_=%s", host, actionToken.Token),
					})
				}
			case secretsType:
				user := services.FindUserByAccountHolder(holder)
				client := services.FindOrCreateClient(services.DefaultClient)
				if user.ID != 0 {
					actionToken := services.CreateAction(user, client,
						c.Request.RemoteAddr,
						c.Request.UserAgent(),
						models.WriteScope,
						models.UpdateUserAction,
					)
					go communications.Announce("session.magic", utils.H{
						"Email":     user.Email,
						"FirstName": user.FirstName,
						"Callback":  fmt.Sprintf("%s/profile/secrets?_=%s", host, actionToken.Token),
					})
				}
			default:
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "update request was not created",
					"error":    "request type not available",
				})
				return
			}

			// No Content is the default response
			c.JSON(http.StatusNoContent, nil)
		})

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/profile", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
			var uuid = c.Param("user_id")

			if !security.ValidUUID(uuid) {
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "User instropection failed",
					"error":    "must use valid UUID for identification",
				})
				return
			}

			action := c.MustGet("Action").(models.Action)
			user := services.FindUserByUUID(uuid)
			if user.IsNewRecord() || user.ID != action.UserID {
				c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "User instropection failed",
					"error":    shared.AccessDenied,
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
			c.String(http.StatusNotImplemented, "Not implemented")
		})

		// Requires X-Requested-By and Origin (same-origin policy)
		// Authorization type: action token / Bearer (for web use)
		usersRoutes.GET("/:user_id/clients", requiresConformance, actionTokenBearerAuthorization, func(c *gin.Context) {
			var uuid = c.Param("user_id")

			if !security.ValidUUID(uuid) {
				c.JSON(http.StatusBadRequest, utils.H{
					"_status":  "error",
					"_message": "User's clients unavailable",
					"error":    "must use valid UUID for identification",
				})
				return
			}

			action := c.MustGet("Action").(models.Action)
			user := services.FindUserByUUID(uuid)
			if user.IsNewRecord() || user.ID != action.UserID {
				c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "User's clients unavailable",
					"error":    shared.AccessDenied,
				})
				return
			}

			c.JSON(http.StatusOK, utils.H{
				"_status":  "success",
				"_message": "User's clients available",
				"clients":  services.ActiveClientsForUser(user.ID),
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
					"error":    "must use valid UUID for identification",
				})
				return
			}

			action := c.MustGet("Action").(models.Action)
			user := services.FindUserByUUID(userUUID)
			if user.IsNewRecord() || user.ID != action.UserID {
				c.Header("WWW-Authenticate", fmt.Sprintf("Bearer realm=\"%s\"", c.Request.RequestURI))
				c.JSON(http.StatusUnauthorized, utils.H{
					"_status":  "error",
					"_message": "Client application irrevocable",
					"error":    shared.AccessDenied,
				})
				return
			}

			client := services.FindClientByUUID(clientUUID)
			services.RevokeClientAccess(client.ID, user.ID)
			c.JSON(http.StatusNoContent, nil)
		})
	}
}
