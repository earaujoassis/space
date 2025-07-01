package repository

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/database"
	"github.com/earaujoassis/space/internal/gateways/memory"
	"github.com/earaujoassis/space/test/utils"
)

type RepositoryTestSuite struct {
	suite.Suite
	Config *config.Config
	Memory *memory.MemoryService
	DB     *database.DatabaseService
}

func (s *RepositoryTestSuite) SetupSuite() {
	utils.SetupConfigEnv()
	s.Config, _ = config.Load()
	db, err := database.NewDatabaseService(s.Config)
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not create new database service: %v", err)
	}
	s.DB = db
	err = utils.RunUnitTestMigrator(db.GetDB())
	s.Require().NoError(err)
	ms, err := memory.NewMemoryService(s.Config)
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not create new memory service: %v", err)
	}
	s.Memory = ms
}

func (s *RepositoryTestSuite) SetupTest() {
	s.cleanupDatabase()
	s.cleanupRedis()
}

func (s *RepositoryTestSuite) cleanupDatabase() {
	db, err := s.DB.GetDB().DB()
	s.Require().NoError(err)
	db.Exec("DELETE FROM sessions")
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM services")
	db.Exec("DELETE FROM clients")
	db.Exec("DELETE FROM languages")
}

func (s *RepositoryTestSuite) cleanupRedis() {
	s.Memory.Do("FLUSHDB")
}

func (s *RepositoryTestSuite) TearDownSuite() {
	s.Memory.Close()
	s.DB.Close()
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
