package plugins

import (
	"fmt"
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SetupSentryForRouter(router *gin.Engine) {
	router.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: false,
	}))
}

func SentryForGin(c *gin.Context, err error) {
	if hub := sentrygin.GetHubFromContext(c); hub != nil {
		hub.CaptureException(fmt.Errorf("%v", err))
		hub.Flush(2 * time.Second)
	}
}
