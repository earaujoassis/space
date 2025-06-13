package api

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

func TestRequiresConformance(t *testing.T) {
	router := gin.New()
	router.Use(requiresConformance)
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
