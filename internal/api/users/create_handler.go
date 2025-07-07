package users

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func createHandler(c *gin.Context) {
	var buf bytes.Buffer
	var imageData string

	fg := ioc.GetFeatureGate(c)
	if !fg.IsActive("user.create") {
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
	repositories := ioc.GetRepositories(c)
	codeSecretKey := repositories.Users().SetCodeSecret(&user)
	recoverSecret, _ := repositories.Users().SetRecoverSecret(&user)
	img, err := codeSecretKey.Image(200, 200)
	if err != nil {
		imageData = ""
	} else {
		png.Encode(&buf, img)
		imageData = base64.StdEncoding.EncodeToString(buf.Bytes())
	}

	user.Client = repositories.Clients().FindOrCreate(models.DefaultClient)
	user.Language = repositories.Languages().FindOrCreate("English", "en-US")
	err = repositories.Users().Create(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.H{
			"_status":  "error",
			"_message": "User was not created",
			"error":    fmt.Sprintf("%v", err),
			"user":     user,
		})
	} else {
		notifier := ioc.GetNotifier(c)
		go notifier.Announce(user, "user.created", shared.NotificationTemplateData(c, utils.H{
			"Email":     shared.GetUserDefaultEmailForNotifications(c, user),
			"FirstName": user.FirstName,
		}))
		c.JSON(http.StatusOK, utils.H{
			"_status":           "created",
			"_message":          "User was created",
			"recover_secret":    recoverSecret,
			"code_secret_image": imageData,
			"user":              user,
		})
	}
}
