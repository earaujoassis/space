package shared

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func defaultRoute(c *gin.Context) {
	c.String(http.StatusOK, "All good")
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestScheme(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	assert.Equal(t, "http", Scheme(req), "default to HTTP scheme/protocol")
}

func TestClientBasicAuthorization(t *testing.T) {
	router := gin.New()
	router.Use(ClientBasicAuthorization)
	router.GET("/", defaultRoute)
	w := performRequest(router, "GET", "/")
	assert.Equal(t, w.Code, 400)
}

func TestOAuthTokenBearerAuthorization(t *testing.T) {
	router := gin.New()
	router.Use(OAuthTokenBearerAuthorization)
	router.GET("/", defaultRoute)
	w := performRequest(router, "GET", "/")
	assert.Equal(t, w.Code, 400)
}
