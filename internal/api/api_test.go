package api

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/test/unit"
)

type ApiTestSuite struct {
	unit.ApiBaseTestSuite
	Router *gin.Engine
}

func (s *ApiTestSuite) SetupSuite() {
	s.ApiBaseTestSuite.SetupSuite()
	gin.SetMode(gin.TestMode)
	s.Router = s.SetupRouter()
}

func (s *ApiTestSuite) SetupTest() {
	s.ApiBaseTestSuite.SetupTest()
	s.Router = s.SetupRouter()
	ExposeRoutes(s.Router)
}

func TestClientsSuite(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}
