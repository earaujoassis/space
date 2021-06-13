package api

import (
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/stretchr/testify/assert"
    "github.com/gin-gonic/gin"
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
    assert.Equal(t, "http", scheme(req), "default to HTTP scheme/protocol")
}

func TestRequiresConformance(t *testing.T) {
    router := gin.New()
    router.Use(requiresConformance)
    router.GET("/", defaultRoute)
    w := performRequest(router, "GET", "/")
    assert.Equal(t, w.Code, 400)
}

func TestClientBasicAuthorization(t *testing.T) {
    router := gin.New()
    router.Use(clientBasicAuthorization)
    router.GET("/", defaultRoute)
    w := performRequest(router, "GET", "/")
    assert.Equal(t, w.Code, 400)
}

func TestActionTokenBearerAuthorization(t *testing.T) {
    router := gin.New()
    router.Use(actionTokenBearerAuthorization)
    router.GET("/", defaultRoute)
    w := performRequest(router, "GET", "/")
    assert.Equal(t, w.Code, 400)
}

func TestOAuthTokenBearerAuthorization(t *testing.T) {
    router := gin.New()
    router.Use(oAuthTokenBearerAuthorization)
    router.GET("/", defaultRoute)
    w := performRequest(router, "GET", "/")
    assert.Equal(t, w.Code, 400)
}
