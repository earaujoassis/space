package unit

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/test/factory"
	"github.com/earaujoassis/space/test/utils"
)

type BaseTestSuite struct {
	suite.Suite
	Config  *config.Config
	AppCtx  *ioc.AppContext
	Factory *factory.TestRepositoryFactory
}

func (s *BaseTestSuite) SetupSuite() {
	if err := utils.EnsureProjectRoot(); err != nil {
		s.T().Fatalf("Failed to change to project root: %v", err)
	}
	utils.SetupConfigEnv()
	s.Config, _ = config.Load()
	appCtx, err := ioc.NewAppContext(s.Config)
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not create new database service: %v", err)
	}
	s.AppCtx = appCtx
	dbService := appCtx.DB
	err = utils.RunUnitTestMigrator(dbService.GetDB())
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not auto-migrate models: %v", err)
	}
	ms := appCtx.Memory
	s.Factory = factory.NewTestRepositoryFactory(dbService, ms)
}

func (s *BaseTestSuite) TearDownSuite() {
	s.AppCtx.Close()
}

func (s *BaseTestSuite) SetupTest() {
	s.cleanupDatabase()
	s.cleanupRedis()
}

func (s *BaseTestSuite) cleanupDatabase() {
	db, err := s.AppCtx.DB.GetDB().DB()
	s.Require().NoError(err)
	db.Exec("DELETE FROM sessions")
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM services")
	db.Exec("DELETE FROM clients")
	db.Exec("DELETE FROM languages")
}

func (s *BaseTestSuite) cleanupRedis() {
	s.AppCtx.Memory.Do("FLUSHDB")
}

func (s *BaseTestSuite) PerformRequest(r http.Handler, method, path string, header *http.Header, cookie *http.Cookie, form *strings.Reader) *httptest.ResponseRecorder {
	var req *http.Request

	if form == nil {
		req, _ = http.NewRequest(method, path, nil)
	} else {
		req, _ = http.NewRequest(method, path, form)
	}
	if header != nil {
		req.Header = *header
	}
	if cookie != nil {
		req.AddCookie(cookie)
	}
	req.Header.Set("User-Agent", gofakeit.UserAgent())
	if form != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req.RemoteAddr = gofakeit.IPv4Address()
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
