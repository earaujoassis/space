package self

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/test/unit"
)

type SelfTestSuite struct {
	unit.ApiBaseTestSuite
	Router *gin.Engine
}

func (s *SelfTestSuite) SetupSuite() {
	s.ApiBaseTestSuite.SetupSuite()
	gin.SetMode(gin.TestMode)
	s.Router = s.SetupRouter()
}

func (s *SelfTestSuite) SetupTest() {
	s.ApiBaseTestSuite.SetupTest()
	s.Router = s.SetupRouter()
	group := s.Router.Group("/api")
	ExposeRoutes(group)
}

func TestClientsSuite(t *testing.T) {
	suite.Run(t, new(SelfTestSuite))
}
