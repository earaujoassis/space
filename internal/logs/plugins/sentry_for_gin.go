package plugins

import (
	"fmt"
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SentryForGin(c *gin.Context, err error) {
	if hub := sentrygin.GetHubFromContext(c); hub != nil {
		hub.CaptureException(fmt.Errorf("%v", err))
		hub.Flush(2 * time.Second)
	}
}
