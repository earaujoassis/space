package helpers

import (
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/gateways/redis"
	"github.com/earaujoassis/space/internal/tasks"
	"github.com/earaujoassis/space/test/factory"
	"github.com/earaujoassis/space/test/utils"
)

type OIDCTestSuite struct {
	suite.Suite
	Resources *TestResources
	Server    *httptest.Server
	Client    *OIDCCTestlient
	Factory   *factory.TestRepositoryFactory
	cfg       *config.Config
}

func (s *OIDCTestSuite) SetupSuite() {
	if err := utils.EnsureProjectRoot(); err != nil {
		s.T().Fatalf("Failed to change to project root: %v", err)
	}

	s.setupConfigEnv()
	s.Resources = NewTestResources()
	s.Resources.StartResources()
	s.cfg, _ = config.Load()
	db, err := database.NewDatabaseService(s.cfg)
	if err != nil {
		s.T().Fatalf("Could not create new database service: %v", err)
	}
	ms, err := redis.NewMemoryService(s.cfg)
	if err != nil {
		s.T().Fatalf("Could not create new memory service: %v", err)
	}
	s.Factory = factory.NewTestRepositoryFactory(db, ms)
	gin.SetMode(gin.TestMode)
	router := s.setupTestRouter()
	s.runMigrations()
	s.Server = httptest.NewServer(router)
	s.Client = NewOIDCTestClient(s.Server.URL)
}

func (s *OIDCTestSuite) TearDownSuite() {
	if s.Resources != nil {
		s.Resources.PurgeResources()
	}
	if s.Server != nil {
		s.Server.Close()
	}
}

func (s *OIDCTestSuite) setupConfigEnv() {
	os.Setenv("SPACE_ENV", "integration")
	os.Setenv("SPACE_APPLICATION_KEY", "masterapplicationkey")
	os.Setenv("SPACE_MAIL_FROM", "example@example.com")
	os.Setenv("SPACE_MAILER_ACCESS", "AccessKeyId:SecretAccessKey:Region")
	os.Setenv("SPACE_SESSION_SECRET", "E93jykumzKrJOp6xKB4JduxaKLmeiPmf")
	os.Setenv("SPACE_SESSION_SECURE", "false")
	os.Setenv("SPACE_STORAGE_SECRET", "KRgwMcZdLPfo9bck")
}

func (s *OIDCTestSuite) setupTestRouter() *gin.Engine {
	router := tasks.SetupRouter(s.cfg)

	return router
}

func (s *OIDCTestSuite) runMigrations() {
	tasks.RunMigrations(s.cfg, "./configs/migrations")
}
