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
	ms *memory.MemoryService
	db *database.DatabaseService
}

func (s *RepositoryTestSuite) SetupSuite() {
	if err := utils.EnsureProjectRoot(); err != nil {
		s.T().Fatalf("Failed to change to project root: %v", err)
	}

	utils.SetupConfigEnv()
	cfg, _ := config.Load()
	db, err := database.NewDatabaseService(cfg)
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not create new database service: %v", err)
	}
	s.db = db
	err = utils.RunUnitTestMigrator(db.GetDB())
	s.Require().NoError(err)
	ms, err := memory.NewMemoryService(cfg)
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not create new memory service: %v", err)
	}
	s.ms = ms
}

func (s *RepositoryTestSuite) SetupTest() {
	s.cleanupDatabase()
	s.cleanupRedis()
}

func (s *RepositoryTestSuite) cleanupDatabase() {
	db, err := s.db.GetDB().DB()
	s.Require().NoError(err)
	db.Exec("DELETE FROM sessions")
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM services")
	db.Exec("DELETE FROM clients")
	db.Exec("DELETE FROM languages")
}

func (s *RepositoryTestSuite) cleanupRedis() {
	s.ms.Do("FLUSHDB")
}

func (s *RepositoryTestSuite) TearDownSuite() {
	s.ms.Close()
	s.db.Close()
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
