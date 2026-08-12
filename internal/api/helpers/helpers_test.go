package helpers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/test/unit"
)

type HelpersTestSuite struct {
	unit.ApiBaseTestSuite
	Router *gin.Engine
}

func (s *HelpersTestSuite) SetupSuite() {
	s.ApiBaseTestSuite.SetupSuite()
	gin.SetMode(gin.TestMode)
	s.Router = s.SetupRouter()
}

func (s *HelpersTestSuite) SetupTest() {
	s.ApiBaseTestSuite.SetupTest()
	s.Router = s.SetupRouter()
}

func TestClientsSuite(t *testing.T) {
	suite.Run(t, new(HelpersTestSuite))
}
