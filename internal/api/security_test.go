package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultRoute(c *gin.Context) {
	c.String(http.StatusOK, "All good")
}

func (s *ApiHandlerTestSuite) TestRequiresConformance() {
	router := gin.New()
	router.Use(requiresConformance())
	router.GET("/", defaultRoute)
	w := s.PerformRequest(router, "GET", "/", nil, nil, nil)
	s.Equal(w.Code, 400)
}

func (s *ApiHandlerTestSuite) TestActionTokenBearerAuthorization() {
	router := gin.New()
	router.Use(actionTokenBearerAuthorization())
	router.GET("/", defaultRoute)
	w := s.PerformRequest(router, "GET", "/", nil, nil, nil)
	s.Equal(w.Code, 400)
}
