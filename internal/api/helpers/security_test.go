package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func defaultRoute(c *gin.Context) {
	c.String(http.StatusOK, "All good")
}

func (s *HelpersTestSuite) TestRequiresConformance() {
	router := s.Router
	router.Use(RequiresConformance())
	router.GET("/", defaultRoute)
	w := s.PerformRequest(router, "GET", "/", nil, nil, nil)
	s.Equal(400, w.Code)
}

func (s *HelpersTestSuite) TestActionTokenBearerAuthorization() {
	router := s.Router
	router.Use(ActionTokenBearerAuthorization())
	router.GET("/", defaultRoute)
	w := s.PerformRequest(router, "GET", "/", nil, nil, nil)
	s.Equal(400, w.Code)
}

func (s *HelpersTestSuite) TestRequiresApplicationSession() {
	router := s.Router
	router.Use(RequiresApplicationSession())
	router.GET("/", defaultRoute)
	w := s.PerformRequest(router, "GET", "/", nil, nil, nil)
	s.Equal(401, w.Code)
}
