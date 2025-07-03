package shared

import (
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
)

func GetUserDefaultEmailForNotifications(c *gin.Context, user models.User) string {
	const (
		realm    = "notifications"
		category = "system-email-notifications"
		property = "email-address"
	)
	repositories := ioc.GetRepositories(c)
	setting := repositories.Settings().FindOrFallback(user, realm, category, property, user.Email)
	value, _ := setting.DeserializeValue()
	return value.(string)
}
