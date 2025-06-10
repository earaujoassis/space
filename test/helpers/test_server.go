package helpers

import (
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/tasks"
)

type OAuthTestSuite struct {
	suite.Suite
	Resources *TestResources
	Server    *httptest.Server
	Client    *OAuthTestClient
}

func (s *OAuthTestSuite) SetupSuite() {
	if err := os.Chdir("../.."); err != nil {
		s.T().Fatalf("Failed to change to project root: %v", err)
	}

	gin.SetMode(gin.TestMode)
	s.setupConfigEnv()
	s.Resources = NewTestResources()
	s.Resources.StartResources()
	router := s.setupTestRouter()
	s.runMigrations()
	s.Server = httptest.NewServer(router)
	s.Client = NewOAuthTestClient(s.Server.URL)
}

func (s *OAuthTestSuite) TearDownSuite() {
	if s.Resources != nil {
		s.Resources.PurgeResources()
	}
	if s.Server != nil {
		s.Server.Close()
	}
}

func (s *OAuthTestSuite) setupConfigEnv() {
	os.Setenv("SPACE_ENV", "testing")
	os.Setenv("SPACE_APPLICATION_KEY", "masterapplicationkey")
	os.Setenv("SPACE_MAIL_FROM", "example@example.com")
	os.Setenv("SPACE_MAILER_ACCESS", "AccessKeyId:SecretAccessKey:Region")
	os.Setenv("SPACE_SESSION_SECRET", "E93jykumzKrJOp6xKB4JduxaKLmeiPmf")
	os.Setenv("SPACE_SESSION_SECURE", "false")
	os.Setenv("SPACE_STORAGE_SECRET", "KRgwMcZdLPfo9bck")
}

func (s *OAuthTestSuite) setupTestRouter() *gin.Engine {
	config.LoadConfig()
	router := tasks.Routes()

	return router
}

func (s *OAuthTestSuite) runMigrations() {
	tasks.RunMigrations("./configs/migrations")
}
