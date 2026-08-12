package services

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/test/unit"
)

type ServicesTestSuite struct {
	unit.ApiBaseTestSuite
	Router *gin.Engine
}

func (s *ServicesTestSuite) SetupSuite() {
	s.ApiBaseTestSuite.SetupSuite()
	gin.SetMode(gin.TestMode)
	s.Router = s.SetupRouter()
}

func (s *ServicesTestSuite) SetupTest() {
	s.ApiBaseTestSuite.SetupTest()
	s.Router = s.SetupRouter()
	group := s.Router.Group("/api")
	ExposeRoutes(group)
}

func TestClientsSuite(t *testing.T) {
	suite.Run(t, new(ServicesTestSuite))
}
