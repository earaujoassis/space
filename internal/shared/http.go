package shared

import (
	"fmt"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
)

// BasicAuthEncode encodes a key-secret pair to be used in a HTTP Basic Authentication
//
//	strategy (HTTP Request Header)
func BasicAuthEncode(key, secret string) string {
	token := key + ":" + secret
	return base64.StdEncoding.EncodeToString([]byte(token))
}

// BasicAuthDecode decodes a key-secret pair from a token string, originally used in a
//
//	HTTP Basic Authentication strategy (HTTP Request Header)
func BasicAuthDecode(token string) (string, string) {
	bytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", ""
	}
	values := strings.Split(string(bytes), ":")
	if len(values) < 2 {
		return "", ""
	}
	return values[0], values[1]
}

// MustServeJSON determines if the HTTP response must be a JSON content or not
//
//	using the HTTP request path and the request header attribute `Accept`
func MustServeJSON(path string, accept string) bool {
	return strings.HasPrefix(path, "/api") ||
		strings.HasPrefix(path, "/token") ||
		strings.HasPrefix(path, "/oauth/token") ||
		strings.HasPrefix(path, "/revoke") ||
		strings.HasPrefix(path, "/oauth/revoke") ||
		strings.HasPrefix(path, "/introspect") ||
		strings.HasPrefix(path, "/oauth/introspect")
}

func GetBaseUrl(c *gin.Context) string {
	scheme := "http"
    if c.Request.TLS != nil {
        scheme = "https"
    }
	return fmt.Sprintf("%s://%s", scheme, c.Request.Host)
}

func SetHeadersCacheControl(c *gin.Context) {
	c.Header("Cache-Control", "no-store")
	c.Header("Pragma", "no-cache")
}
