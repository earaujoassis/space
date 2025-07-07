package users

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/test/unit"
)

type UsersTestSuite struct {
	unit.ApiBaseTestSuite
	Router *gin.Engine
}

func (s *UsersTestSuite) SetupSuite() {
	s.ApiBaseTestSuite.SetupSuite()
	gin.SetMode(gin.TestMode)
	s.Router = s.SetupRouter()
}

func (s *UsersTestSuite) SetupTest() {
	s.ApiBaseTestSuite.SetupTest()
	s.Router = s.SetupRouter()
	group := s.Router.Group("/api")
	ExposeRoutes(group)
}

func TestClientsSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}
