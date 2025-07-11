package security

import (
	"github.com/gin-gonic/gin"
)

func SetTrustedProxies(router *gin.Engine) {
	trustedProxies := []string{
		"127.0.0.1",
		"::1",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	router.SetTrustedProxies(trustedProxies)
}
