package shared

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/earaujoassis/space/internal/utils"
)

func TestNotificationTemplateData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("GET", "/", nil)
	req.Host = "example.com"
	c.Request = req
	require.Equal(t, "example.com", c.Request.Host)

	data := NotificationTemplateData(c, utils.H{})
	assert.Equal(t, "http://example.com/notifications", data["NotificationsLink"])
	assert.Equal(t, time.Now().Year(), data["Year"])

	extend := make(utils.H)
	extend["Additional"] = "Content"
	data = NotificationTemplateData(c, extend)
	assert.Equal(t, "http://example.com/notifications", data["NotificationsLink"])
	assert.Equal(t, time.Now().Year(), data["Year"])
	assert.Equal(t, "Content", data["Additional"])
}
