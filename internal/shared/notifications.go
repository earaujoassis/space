package shared

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/utils"
)

func NotificationTemplateData(c *gin.Context, data utils.H) utils.H {
	baseURL := GetBaseUrl(c)
	defaults := utils.H{
		"Year":              time.Now().Year(),
		"NotificationsLink": fmt.Sprintf("%s/notifications", baseURL),
	}

	for k, v := range data {
		defaults[k] = v
	}

	return defaults
}
