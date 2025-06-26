package web

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/utils"
)

func profileSecretsHandler(c *gin.Context) {
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
}
